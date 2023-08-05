import argparse
import hashlib
import os
import platform
import re
import sys
import unicodedata
from sklearn.externals import joblib
from io import StringIO

import pandas as pd
from thoipapy.homologues.NCBI_download import download_homologues_from_ncbi
from thoipapy.validation.validation import drop_cols_not_used_in_ML

import thoipapy
from thoipapy.features.feature_calculate import create_PSSM_from_MSA, lipo_from_pssm, entropy_calculation, \
    coevolution_calculation_with_freecontact, parse_freecontact_coevolution, calc_relative_position, \
    LIPS_score_calculation, \
    parse_LIPS_score, motifs_from_seq, combine_all_features, add_physical_parameters_to_features
from thoipapy.homologues.NCBI_parser import parse_NCBI_xml_to_csv, extract_filtered_csv_homologues_to_alignments


def get_start_end_pl_surr(TMD_start, TMD_end, seqlen, surr):
    TMD_start_pl_surr = TMD_start - surr
    if TMD_start_pl_surr < 1:
        TMD_start_pl_surr = 1
    TMD_end_pl_surr = TMD_end + surr
    if TMD_end_pl_surr > seqlen:
        TMD_end_pl_surr = seqlen
    return TMD_start_pl_surr, TMD_end_pl_surr

def run_THOIPA_prediction(s,protein_name, TMD_seq, full_seq, predictions_folder):
    """Function to run standalone THOIPA prediction for a protein transmembrane domain of interest.

    Parameters
    ----------
    s : dict
        Settings dictionary
    protein_name : str
        Protein name
    TMD_seq : str
        Transmembrane domain sequence
    full_seq : str
        Full protein sequence
    predictions_folder : str
        Path to folder where the output will be saved

    Saved files and figures
    -----------------------


    Usage
    -----
    import thoipapy
    protein_name = "1c17_A"
    TMD_seq = "AAVMMGLAAIGAAIGIGILG"
    full_seq = "MENLNMDLLYMAAAVMMGLAAIGAAIGIGILGGKFLEGAARQPDLIPLLRTQFFIVMGLVDAIPMIAVGLGLYVMFAVA"
    predictions_folder = "/path/to/your/output/folder"
    thoipapy.run_THOIPA_prediction(protein_name, TMD_seq, full_seq, predictions_folder)
    """
    # default model is the combined ETRA, NMR, crystal non-redundant dataset (set05)
    set_number = 5

    TMD_plus_full_seq = TMD_seq + "_" + full_seq
    # adjust encoding for md5 creation
    TMD_plus_full_seq = unicodedata.normalize('NFKD', TMD_plus_full_seq).encode('ascii','ignore')
    hash_object = hashlib.md5(TMD_plus_full_seq)
    md5 = hash_object.hexdigest()
    folder_with_md5_exists = False
    if folder_with_md5_exists:
        create_indiv_name_and_email = 99
        sys.stdout.write("You shuld just email the results here.")
    acc = md5[0:6]
    out_folder = os.path.join(predictions_folder, md5)
    if os.path.isdir(out_folder):
        sys.stdout.write("This sequence has already been analysed![folder exists]")
    else:
        thoipapy.utils.make_sure_path_exists(out_folder)
    logfile = os.path.join(out_folder, "logfile.txt")
    logging = thoipapy.common.setup_error_logging(logfile, "DEBUG", "DEBUG", print_system_info=False)
    seqlen = len(full_seq)
    hitlist = re.findall(TMD_seq, full_seq)
    if len(hitlist) > 1:
        logging.warning("TMD sequence is found multiple times in the full sequence.")
        return
    elif len(hitlist) == 0:
        logging.warning("TMD sequence was not found in full protein sequence. Please check input sequences.")
        return
        # the following code will only continue if the TMD seq is found ONCE in the full protein seq
    s["Protein_folder"] = predictions_folder
    blast_xml_file = os.path.join(out_folder, "BLAST_results.xml")
    xml_txt = blast_xml_file[:-4] + "_details.txt"
    xml_tar_gz = os.path.join(out_folder, "BLAST_results.xml.tar.gz")
    BLAST_csv_tar = os.path.join(out_folder, "BLAST_results.csv.tar.gz")
    fasta_all_TMD_seqs = os.path.join(out_folder,"homologues.redundant.fas")
    path_uniq_TMD_seqs_for_PSSM_FREECONTACT = os.path.join(out_folder,"homologues.uniq.for_PSSM_FREECONTACT.txt")
    path_uniq_TMD_seqs_no_gaps_for_LIPS = os.path.join(out_folder,"homologues.uniq.for_LIPS.txt")
    path_uniq_TMD_seqs_surr5_for_LIPO = os.path.join(out_folder,"homologues.uniq.for_LIPO.txt")
    pssm_csv = os.path.join(out_folder,"pssm.csv")
    pssm_surr5_csv = os.path.join(out_folder,"pssm_surr5.csv")
    lipo_csv = os.path.join(out_folder,"lipo.csv")
    entropy_file = os.path.join(out_folder,"entropy.csv")
    freecontact_file = os.path.join(out_folder,"freecontact_out.csv")
    freecontact_parsed_csv = os.path.join(out_folder,"freecontact_parsed.csv")
    relative_position_file = os.path.join(out_folder,"relative_position.csv")
    LIPS_output_file = os.path.join(out_folder,"LIPS_output.csv")
    LIPS_parsed_csv = os.path.join(out_folder,"LIPS_output_parsed.csv")
    motifs_file = os.path.join(out_folder,  "motifs.csv")
    full_seq_fasta_file = os.path.join(out_folder,  "protein.fasta")
    full_seq_phobius_output_file = os.path.join(out_folder,  "protein.phobius")
    feature_combined_file = os.path.join(out_folder, "features_combined.csv")
    alignment_summary_csv = os.path.join(out_folder, "homologues.alignment_summary.csv")
    THOIPA_full_out_csv = os.path.join(out_folder, "THOIPA_full_out.csv")
    THOIPA_pretty_out_xlsx = os.path.join(out_folder, "THOIPA_out.xlsx")
    THOIPA_pretty_out_csv = os.path.join(out_folder, "THOIPA_out.csv")

    logging.info("md5 checksum of TMD and full sequence = {}".format(md5))
    logging.info("acc = md5[0:6] = {}".format(acc))
    #logging.shutdown()
    #logging = korbinian.utils.Log_Only_To_Console()

    m = re.search(TMD_seq, full_seq)
    TMD_start = m.start()
    TMD_end = m.end()
    logging.info("TMD found in full sequence. start = {}, end = {}".format(TMD_start, TMD_end))

    surr = s["num_of_sur_residues"]
    TMD_start_pl_surr, TMD_end_pl_surr = get_start_end_pl_surr(TMD_start, TMD_end, seqlen, surr=surr)
    TMD_start_pl_5, TMD_end_pl_5 = get_start_end_pl_surr(TMD_start, TMD_end, seqlen, surr=5)

    logging.info("TMD_start_pl_surr = {}, TMD_end_pl_surr = {}".format(TMD_start_pl_surr, TMD_end_pl_surr))

    # due to uniprot indexing, the TMD seq starts 1 residue earlier in the python index
    TMD_seq_pl_surr = full_seq[TMD_start_pl_surr: TMD_end_pl_surr]
    # currently I don't know why this should not be TMD_start_pl_5 - 1
    query_TMD_seq_surr5 = full_seq[TMD_start_pl_5: TMD_end_pl_5]
    logging.info("TMD_seq : {}".format(TMD_seq))
    logging.info("query_TMD_seq_surr5 : {}".format(query_TMD_seq_surr5))
    logging.info("TMD_seq_pl_surr : {}".format(TMD_seq_pl_surr))

    expect_value = s["expect_value"]
    hit_list_size = s["hit_list_size"]
    if not os.path.isfile(xml_tar_gz):
        download_homologues_from_ncbi(acc, TMD_seq_pl_surr, blast_xml_file, xml_txt, xml_tar_gz, expect_value, hit_list_size, logging)

    if not os.path.isfile(BLAST_csv_tar):
        parse_NCBI_xml_to_csv(s, acc, xml_tar_gz, BLAST_csv_tar, TMD_start, TMD_end, logging)

    #if not os.path.isfile(path_uniq_TMD_seqs_for_PSSM_FREECONTACT):
    extract_filtered_csv_homologues_to_alignments(s, acc, len(TMD_seq), fasta_all_TMD_seqs, path_uniq_TMD_seqs_for_PSSM_FREECONTACT,
                                                      path_uniq_TMD_seqs_no_gaps_for_LIPS, path_uniq_TMD_seqs_surr5_for_LIPO, BLAST_csv_tar,
                                                      TMD_seq, query_TMD_seq_surr5, logging)

    create_PSSM_from_MSA(path_uniq_TMD_seqs_for_PSSM_FREECONTACT, pssm_csv, acc, TMD_seq, logging)
    create_PSSM_from_MSA(path_uniq_TMD_seqs_surr5_for_LIPO, pssm_surr5_csv, acc, query_TMD_seq_surr5, logging)

    tm_surr_left = TMD_start - TMD_start_pl_surr
    tm_surr_right = TMD_end_pl_surr - TMD_end
    # reset to 5. AM NOT SURE WHY.
    if tm_surr_left >= 5:
        tm_surr_left_lipo = 5
    else:
        tm_surr_left_lipo = tm_surr_left
    if tm_surr_right >= 5:
        tm_surr_right_lipo = 5
    else:
        tm_surr_right_lipo = tm_surr_right

    thoipapy_module_path = os.path.dirname(os.path.abspath(thoipapy.__file__))
    hydrophob_scale_path = os.path.join(thoipapy_module_path, "setting", "hydrophobicity_scales.xlsx")

    lipo_from_pssm(acc, pssm_surr5_csv, lipo_csv, tm_surr_left_lipo, tm_surr_right_lipo, s["lipophilicity_scale"], logging, plot_linechart=True)
                 # (acc, pssm_csv_surr5, lipo_csv, tm_surr_left, tm_surr_right, scalename, logging, plot_linechart = False)

    entropy_calculation(acc, path_uniq_TMD_seqs_for_PSSM_FREECONTACT, TMD_seq, entropy_file, logging)

    if "Windows" in platform.system():
        logging.warning("\n Freecontact cannot be run in Windows! Skipping coevolution_calculation_with_freecontact."
                         "For testing, copy file from Linux and rename to {}".format(freecontact_file))
    else:
        coevolution_calculation_with_freecontact(path_uniq_TMD_seqs_for_PSSM_FREECONTACT, freecontact_file, s["freecontact_dir"], logging)

    parse_freecontact_coevolution(acc, freecontact_file, freecontact_parsed_csv, TMD_start, TMD_end, logging)

    calc_relative_position(acc, path_uniq_TMD_seqs_for_PSSM_FREECONTACT, relative_position_file, TMD_start, seqlen, logging)

    LIPS_score_calculation(path_uniq_TMD_seqs_no_gaps_for_LIPS, LIPS_output_file)

    parse_LIPS_score(acc, LIPS_output_file, LIPS_parsed_csv, logging)
    motifs_from_seq(TMD_seq, TMD_seq_pl_surr, tm_surr_left, tm_surr_right, motifs_file, logging)

    database = "standalone_prediction"
    #combine_all_features(acc, database, TMD_seq, feature_combined_file, entropy_file, pssm_csv, lipo_csv, freecontact_parsed_csv, relative_position_file, LIPS_parsed_csv,motifs_file, alignment_summary_csv, logging)
    combine_all_features(s, full_seq, acc, database, TMD_seq, TMD_start, feature_combined_file, entropy_file, pssm_csv,
                         lipo_csv, freecontact_parsed_csv, relative_position_file, LIPS_parsed_csv, motifs_file,
                         alignment_summary_csv, full_seq_fasta_file,full_seq_phobius_output_file,logging)

    add_physical_parameters_to_features(acc, feature_combined_file, logging)

    df_data = pd.read_csv(feature_combined_file, index_col=0)
    df_ML_input = drop_cols_not_used_in_ML(logging, df_data, s["excel_file_with_settings"])

    cols_to_keep = ["residue_num", "residue_name", "res_num_full_seq"]

    #df_out = df_data[cols_to_keep]
    #df_out.merge(df_ML_input, how="left")

    df_out = pd.concat([df_data[cols_to_keep], df_ML_input], axis=1, join="outer")

    # load the trained machine learning model, saved in the thoipapy directory
    thoipapy_module_path = os.path.dirname(os.path.abspath(thoipapy.__file__))
    machine_learning_model = os.path.join(thoipapy_module_path, "ML_model", "set{:02d}_ML_model.lpkl".format(set_number))
    #machine_learning_model = "/media/mark/sindy/m_data/THOIPA_data/Results/set05/set05_ML_model.lpkl"
    fit = joblib.load(machine_learning_model)

    df_out["THOIPA"] = fit.predict_proba(df_ML_input)[:, 1]

    df_pretty_out = df_out[["residue_num", "residue_name", "THOIPA"]]
    df_pretty_out.columns = ["residue number", "residue name", "THOIPA"]

    df_pretty_out["THOIPA"] = df_pretty_out["THOIPA"].round(3)
    df_pretty_out.to_excel(THOIPA_pretty_out_xlsx, index=False)

    # pad all content with spaces so it lines up with the column name
    df_pretty_out["residue number"] = df_pretty_out["residue number"].apply(lambda x: "{: >14}".format(x))
    df_pretty_out["residue name"] = df_pretty_out["residue name"].apply(lambda x: "{: >12}".format(x))
    df_pretty_out["THOIPA"] = df_pretty_out["THOIPA"].apply(lambda x : "{:>6.03f}".format(x))

    df_pretty_out.set_index("residue number", inplace=True)

    df_out.to_csv(THOIPA_full_out_csv)
    df_pretty_out.to_csv(THOIPA_pretty_out_csv, sep="\t")

    # print exactly what the CSV looks like
    out = StringIO()
    df_pretty_out.to_csv(out, sep="\t")
    print(out.getvalue())


# read the command line arguments
parser = argparse.ArgumentParser()

parser.add_argument("-s",  # "-settingsfile",
                    help=r'Full path to your excel settings file.'
                         r'E.g. "\Path\to\your\settingsfile.xlsx"')
parser.add_argument("-i",  # "-input_file",
                    help=r'Full path to your input file with name, TMD_seq, and full_seq.'
                         r'E.g. "\Path\to\your\file\Q12983.txt"')
parser.add_argument("-f",  # "-folder",
                    help=r'Path to your output folder.'
                         r'E.g. "D:\data_thoipapy\Predictions"')

if __name__ == "__main__":
    sys.stdout.write("\nUsage example:\n")
    sys.stdout.write(r"python run_THOIPA_prediction.py -s D:\Dropbox\tm_homodimer_dropbox\thoipapy_run_settings_MT_win10_work.xlsx -i D:\Dropbox\tm_homodimer_dropbox\THOIPA_prediction_inputs\Q12983.txt -f D:\data_thoipapy\Predictions\server-like")
    sys.stdout.write("\n\n")
    sys.stdout.flush()
    # get the command-line arguments
    args = parser.parse_args()
    #settings_path = r"D:\Dropbox\tm_homodimer_dropbox\thoipapy_run_settings_MT_win10_work.xlsx"
    settings_path = args.s
    s = thoipapy.common.create_settingdict(settings_path)
    #test_protein, TMD_seq, full_seq = "1c17_A", "AAVMMGLAAIGAAIGIGILG", "MENLNMDLLYMAAAVMMGLAAIGAAIGIGILGGKFLEGAARQPDLIPLLRTQFFIVMGLVDAIPMIAVGLGLYVMFAVA"
    #test_protein, TMD_seq, full_seq = "Q12983", "FLKVFLPSLLLSHLLAIGLGIYIG", "MGDAAADPPGPALPCEFLRPGCGAPLSPGAQLGRGAPTSAFPPPAAEAHPAARRGLRSPQLPSGAMSQNGAPGMQEESLQGSWVELHFSNNGNGGSVPASVSIYNGDMEKILLDAQHESGRSSSKSSHCDSPPRSQTPQDTNRASETDTHSIGEKNSSQSEEDDIERRKEVESILKKNSDWIWDWSSRPENIPPKEFLFKHPKRTATLSMRNTSVMKKGGIFSAEFLKVFLPSLLLSHLLAIGLGIYIGRRLTTSTSTF"

    input_csv = args.i

    input_ser = pd.Series.from_csv(input_csv)
    protein_name = input_ser["name"]
    TMD_seq = input_ser["TMD_seq"]
    full_seq = input_ser["full_seq"]

    #predictions_folder = r"D:\data_thoipapy\Predictions"
    predictions_folder = os.path.normpath(args.f)

    run_THOIPA_prediction(s,protein_name, TMD_seq, full_seq, predictions_folder)