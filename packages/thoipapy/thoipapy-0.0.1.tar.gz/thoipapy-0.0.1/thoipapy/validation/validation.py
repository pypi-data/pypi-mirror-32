import warnings
warnings.simplefilter(action='ignore', category=FutureWarning)
import pandas as pd
import numpy as np
import thoipapy
import sys
from scipy import interp
import matplotlib.pyplot as plt
from sklearn.ensemble import ExtraTreesClassifier, ExtraTreesRegressor
from sklearn.metrics import roc_curve, auc
from sklearn.model_selection import StratifiedKFold
import os
from sklearn.externals import joblib
import pickle
import time
#from korbinian.utils import convert_truelike_to_bool, convert_falselike_to_bool, Log_Only_To_Console
from thoipapy.utils import convert_truelike_to_bool, convert_falselike_to_bool, Log_Only_To_Console
from multiprocessing import Pool

# intersect function
def intersect(a, b):
     return list(set(a) & set(b))

def drop_cols_not_used_in_ML(logging, df_data, excel_file_with_settings, i=0):
    """Remove columns not used in machine learning training or testing.

    This includes
        1) text and name columns (e.g. residue_name)
        2) bind columns (interface, interface_score).
        3) columns of non-normalised features, or features that we don't want to include anymore.

    Parameters
    ----------
    logging : logging.Logger
        Python object with settings for logging to console and file.
    df_data : pd.DataFrame
        Dataframe with either the training or test dataset
        columns = 'acc_db', "residue_num", "residue_name", etc
        rows = range of number of residues
    excel_file_with_settings : str
        Path to excel file with all settings.
        Necessary for getting the list of features included in the THOIPA algorithm.


    Returns
    -------
    df_data : pd.DataFrame
        Dataframe with either the training or test dataset
        columns :
    """
    # read the features tab of the excel settings file
    features_df = pd.read_excel(excel_file_with_settings, sheetname="features", index_col=0)
    features_df.drop("Notes", axis=0, inplace=True)
    # convert "WAHR" etc to true and false
    features_df["include"] = features_df["include"].apply(convert_truelike_to_bool, convert_nontrue=False)
    features_df["include"] = features_df["include"].apply(convert_falselike_to_bool)

    # print any features that are not shared between the two lists
    unused_features_in_excel = set(features_df.index) - set(df_data.columns)
    features_missing_from_excel = set(df_data.columns) - set(features_df.index)

    # only print this stuff for the first TMD analysed, if in a list
    if i == 0:
        if len(features_missing_from_excel) > 1:
            logging.info("\nfeatures_missing_from_excel, {}".format(features_missing_from_excel))
            if "Unnamed: 0" in features_missing_from_excel:
                raise IndexError("Unnamed column is in dataframe. Try opening csv with index_col=0.")
        if len(unused_features_in_excel) > 1:
            logging.info("\nunused_features_in_excel, {}".format(unused_features_in_excel))
    # drop any features that are not labeled TRUE for inclusion
    features_df = features_df.loc[features_df.include == True]
    # filter df_data to only keep the desired feature columns
    feature_list = features_df.index.tolist()
    df_data = df_data.loc[:, feature_list]
    return df_data

def THOIPA_classifier_with_settings(s, n_features, totally_randomized_trees=False):
    """ For tuning the RF parameters, they are always in one place, and determined by the settings file.

    Parameters
    ----------
    s

    Returns
    -------

    """
    # convert max_features to python None if "None"
    max_features = None if s["max_features"] == "None" else s["max_features"]
    # default for min_samples_leaf is 1 (no filter)
    min_samples_leaf = 1 if s["min_samples_leaf"] == "None" else s["min_samples_leaf"]
    # default for min_samples_leaf is None (no restriction on tree size)
    max_depth = None if s["max_depth"] == "None" else s["max_depth"]

    # forest = RandomForestClassifier(n_estimators=s["RF_number_of_estimators"], n_jobs=s["n_CPU_cores"], criterion=s["criterion"],
    #                                 min_samples_leaf=min_samples_leaf,
    #                                 max_depth=max_depth,
    #                                 oob_score=True, max_features=max_features, bootstrap=bool(s["bootstrap"]),
    #                                 #random_state=s["random_state"]
    #                                 )

    if totally_randomized_trees:
        max_features = 1
        min_samples_leaf = 1
        max_depth = None

    if s["bind_column"] == "interface":
        cls = ExtraTreesClassifier(n_estimators=s["RF_number_of_estimators"], n_jobs=s["n_CPU_cores"], criterion=s["criterion"],
                                        min_samples_leaf=min_samples_leaf,
                                        max_depth=max_depth,
                                        oob_score=bool(s["oob_estimation"]),
                                        bootstrap=bool(s["bootstrap"]),
                                        max_features=max_features,
                                        random_state=s["random_state"]
                                        )

    # in case you want to try to run the ML as a regressor
    elif s["bind_column"] == "interface_score_norm":
        cls = ExtraTreesRegressor(n_estimators=s["RF_number_of_estimators"], n_jobs=s["n_CPU_cores"],
                                        #criterion=s["criterion"],
                                        min_samples_leaf=min_samples_leaf,
                                        max_depth=max_depth,
                                        oob_score=bool(s["oob_estimation"]),
                                        bootstrap=bool(s["bootstrap"]),
                                        max_features=max_features,
                                        random_state=s["random_state"]
                                        )
    else:
        raise ValueError("bind column in excel settings file is not recognised ({})".format(s["bind_column"]))

    return cls

def train_machine_learning_model(s, logging):
    """Train the machine learning model for a particular set.

    Parameters
    ----------
    s : dict
        Settings dictionary
    logging : logging.Logger
        Python object with settings for logging to console and file.

    Sawed Files
    -----------
    model_pkl : pickle
        Pickle containing the trained machine learning model.

    """
    logging.info('starting to predict etra data with THOIPA prediction model')

    train_data_csv = os.path.join(s["set_results_folder"], "{}_train_data.csv".format(s["setname"]))
    train_data_used_for_model_csv = os.path.join(s["set_results_folder"], "{}_train_data_used_for_model.csv".format(s["setname"]))
    model_pkl = os.path.join(s["set_results_folder"], "{}_ML_model.lpkl".format(s["setname"]))

    df_data = pd.read_csv(train_data_csv, index_col=0)

    df_data = df_data.loc[df_data.n_homologues >= s["min_n_homol_training"]]
    df_data = df_data.dropna()

    y = df_data["interface"]
    X = drop_cols_not_used_in_ML(logging, df_data, s["excel_file_with_settings"])

    X.to_csv(train_data_used_for_model_csv)

    if 1 not in y.tolist():
        raise ValueError("None of the residues are marked 1 for an interface residue!")

    n_features = X.shape[1]
    forest = THOIPA_classifier_with_settings(s, n_features)
    # save machine learning model into local driver
    # pkl_file = r'D:\thoipapy\RandomForest\ML_model.lpkl'
    fit = forest.fit(X, y)
    joblib.dump(fit, model_pkl)

    tree_depths = np.array([estimator.tree_.max_depth for estimator in forest.estimators_])
    logging.info("tree depth mean = {} ({})".format(tree_depths.mean(), tree_depths))

    logging.info('finished training machine learning algorithm ({})'.format(model_pkl))


def run_10fold_cross_validation(s, logging):
    """Run 10-fold cross-validation for a particular set of TMDs (e.g. set04).

    The SAME SET is used for both training and cross-validation.

    The number of folds is determined by "cross_validation_number_of_splits" in the settings file.

    The number of estimators in the machine learning algorithm is determined by "RF_number_of_estimators" in the settings file,
    therefore it should match the full training result of the whole dataset.

    IMPORTANT. CURRENTLY THERE IS NO AUTOMATIC REDUNDANCY CHECK.
     - homologues of the tested protein could be in the training dataset

    Parameters
    ----------
    s : dict
        Settings dictionary
    logging : logging.Logger
        Python object with settings for logging to console and file.

    Saved Files
    -----------
    crossvalidation_pkl : pickle
        Pickled dictionary (xv_dict) containing the results for each fold of validation.
        Also contains the mean ROC curve, and the mean AUC.
    """
    logging.info('10-fold cross validation is running')
    train_data_csv = os.path.join(s["set_results_folder"], "{}_train_data.csv".format(s["setname"]))
    #crossvalidation_csv = os.path.join(s["set_results_folder"], "crossvalidation", "data", "{}_10F_data.csv".format(s["setname"]))
    crossvalidation_pkl = os.path.join(s["set_results_folder"], "crossvalidation", "data", "{}_10F_data.pkl".format(s["setname"]))
    features_csv = os.path.join(s["set_results_folder"], "{}_test_features.csv".format(s["setname"]))

    thoipapy.utils.make_sure_path_exists(crossvalidation_pkl, isfile=True)

    df_data = pd.read_csv(train_data_csv, index_col=0)
    df_data = df_data.dropna()

    # drop training data (full protein) that don't have enough homologues
    df_data = df_data.loc[df_data.n_homologues >= s["min_n_homol_training"]]

    X = drop_cols_not_used_in_ML(logging, df_data, s["excel_file_with_settings"])
    y = df_data["interface"]

    skf = StratifiedKFold(n_splits=s["cross_validation_number_of_splits"])
    cv = list(skf.split(X, y))

    n_features = X.shape[1]
    forest = THOIPA_classifier_with_settings(s, n_features)

    mean_tpr = 0.0
    mean_fpr = np.linspace(0, 1, 100)
    # save all outputs to a cross-validation dictionary, to be saved as a pickle file
    xv_dict = {}

    start = time.clock()

    for i, (train, test) in enumerate(cv):
        sys.stdout.write("f{}.".format(i+1)), sys.stdout.flush()
        probas_ = forest.fit(X.iloc[train], y.iloc[train]).predict_proba(X.iloc[test])
        # Compute ROC curve and area the curve
        fpr, tpr, thresholds = roc_curve(y.iloc[test], probas_[:, 1], drop_intermediate=False)
        xv_dict["fpr{}".format(i)] = fpr
        xv_dict["tpr{}".format(i)] = tpr
        mean_tpr += interp(mean_fpr, fpr, tpr)
        mean_tpr[0] = 0.0
    sys.stdout.write("\n"), sys.stdout.flush()

    logging.info("tree depths : {}".format([estimator.tree_.max_depth for estimator in forest.estimators_]))

    duration = time.clock() - start

    mean_tpr /= len(cv)
    mean_tpr[-1] = 1.0

    mean_auc = auc(mean_fpr, mean_tpr)

    xv_dict["true_positive_rate_mean"] = mean_tpr
    xv_dict["false_positive_rate_mean"] = mean_fpr
    xv_dict["mean_auc"] = mean_auc

    # save dict as pickle
    with open(crossvalidation_pkl, "wb") as f:
        pickle.dump(xv_dict, f, protocol=pickle.HIGHEST_PROTOCOL)
    
    features_ser = pd.Series(X.columns)
    features_ser.to_csv(features_csv)
    logging.info('{} 10-fold validation. AUC({:.3f}). Time taken = {:.2f}.\nFeatures: {}'.format(s["setname"], mean_auc, duration, X.columns.tolist()))

def create_10fold_cross_validation_fig(s, logging):
    """Create figure showing ROC curve for each fold in a 10-fold validation.

    The underlying data is created by run_10fold_cross_validation. If this has not been run,
    it will return a file-not-found error.

    Parameters
    ----------
    s : dict
        Settings dictionary
    logging : logging.Logger
        Python object with settings for logging to console and file.
    """
    #plt.rcParams.update({'font.size': 7})
    crossvalidation_png = os.path.join(s["set_results_folder"], "crossvalidation", "{}_10F_ROC.png".format(s["setname"]))
    crossvalidation_pkl = os.path.join(s["set_results_folder"], "crossvalidation", "data", "{}_10F_data.pkl".format(s["setname"]))

    # open pickle file
    with open(crossvalidation_pkl, "rb") as f:
        xv_dict = pickle.load(f)

    figsize = np.array([3.42, 3.42]) * 2 # DOUBLE the real size, due to problems on Bo computer with fontsizes
    fig, ax = plt.subplots(figsize=figsize)

    for i in range(s["cross_validation_number_of_splits"]):
        roc_auc = auc(xv_dict["fpr{}".format(i)], xv_dict["tpr{}".format(i)])
        ax.plot(xv_dict["fpr{}".format(i)], xv_dict["tpr{}".format(i)], lw=1, label='fold %d (area = %0.2f)' % (i, roc_auc), alpha=0.8)

    mean_auc = xv_dict["mean_auc"]

    ax.plot(xv_dict["false_positive_rate_mean"], xv_dict["true_positive_rate_mean"], color="k", label='mean (area = %0.2f)' % mean_auc, lw=1.5)
    ax.plot([0, 1], [0, 1], '--', color=(0.6, 0.6, 0.6), label='random')
    ax.set_xlim([-0.05, 1.05])
    ax.set_ylim([-0.05, 1.05])
    ax.set_xlabel("False positive rate")
    ax.set_ylabel("True positive rate")
    ax.legend(loc="lower right")
    fig.tight_layout()
    fig.savefig(crossvalidation_png, dpi=240)
    #fig.savefig(crossvalidation_png[:-4] + ".pdf")
    fig.savefig(thoipapy.utils.pdf_subpath(crossvalidation_png))


def run_LOO_validation(s, df_set, logging):
    """Run Leave-One-Out cross-validation for a particular set of TMDs (e.g. set05).

    The SAME SET is used for both training and cross-validation.
    ONLY USE IF YOUR DATASET IS NON-REDUNDANT! OR USE AUTO CD-HIT REDUNDANCY CHECKS!
    Each protein is uniquely identified by the acc and the database (acc_db).

    The training set (df_train) consists of all proteins except the one tested.
    The test dataset (df_test) contains only the interface and features of the one test protein

    The model is trained on the train data csv (e.g. set38_train_data.csv)
         - for crystal subset, hetero contacts (folding residues) are removed from this training set!

    The model is validated against each combined CSV with features (e.g. "D:\data_thoipapy\Features\combined\ETRA\Q12983.surr20.gaps5.combined_features.csv")
     - for crystal subset, folding residues (hetero contacts) are INCLUDED here.
     - the model created without the folding contacts is therefore validated against the full seq, including folding residues

    Parameters
    ----------
    s : dict
        Settings dictionary
    df_set : pd.DataFrame
        Dataframe containing the list of proteins to process, including their TMD sequences and full-length sequences
        index : range(0, ..)
        columns : ['acc', 'seqlen', 'TMD_start', 'TMD_end', 'tm_surr_left', 'tm_surr_right', 'database',  ....]
    logging : logging.Logger
        Python object with settings for logging to console and file.

    Saved Files
    -----------
    LOO_crossvalidation_pkl, pickle
        Pickled dictionary (xv_dict) containing the results for each fold of validation.
        Also contains the mean ROC curve, and the mean AUC.
    BO_all_data_csv, csv
        CSV with the BO curve underlying data
    BO_data_excel, csv
        excel file with the processed BO-curve data
    """
    logging.info('Leave-One-Out cross validation is running')
    names_excel_path = os.path.join(os.path.dirname(s["sets_folder"]), "ETRA_NMR_names.xlsx")

    # drop redundant proteins according to CD-HIT
    df_set = thoipapy.utils.drop_redundant_proteins_from_list(df_set, logging)

    train_data_csv = os.path.join(s["set_results_folder"], "{}_train_data.csv".format(s["setname"]))
    LOO_crossvalidation_pkl = os.path.join(s["set_results_folder"], "crossvalidation", "data", "{}_LOO_crossvalidation.pkl".format(s["setname"]))
    BO_all_data_csv = os.path.join(s["set_results_folder"], "crossvalidation", "data", "{}_LOO_BO_data.csv".format(s["setname"]))
    BO_curve_folder = os.path.join(s["set_results_folder"], "crossvalidation")
    BO_data_excel = os.path.join(BO_curve_folder, "data", "{}_BO_curve_data.xlsx".format(s["setname"]))

    thoipapy.utils.make_sure_path_exists(BO_data_excel, isfile=True)

    df_data = pd.read_csv(train_data_csv, index_col=0)
    df_data = df_data.dropna()

    # drop training data (full protein) that don't have enough homologues
    df_data = df_data.loc[df_data.n_homologues >= s["min_n_homol_training"]]

    acc_db_list = df_data.acc_db.unique()
    start = time.clock()
    pred_colname = "THOIPA_{}_LOO".format(s["set_number"])

    n_features = thoipapy.validation.validation.drop_cols_not_used_in_ML(logging, df_data, s["excel_file_with_settings"]).shape[1]
    forest = thoipapy.validation.validation.THOIPA_classifier_with_settings(s, n_features)

    #list of all dictionaries for each protein, for multiprocessing
    d_list = []

    if s["use_multiprocessing"]:
        # TURN LOGGING OFF BEFORE MULTIPROCESSING
        logger = Log_Only_To_Console()
    else:
        logger = logging

    val_list = []
    for n, i in enumerate(df_set.index):
        acc, acc_db, database  = df_set.loc[i, "acc"], df_set.loc[i, "acc_db"], df_set.loc[i, "database"]
        THOIPA_prediction_csv = os.path.join(s["thoipapy_data_folder"], "Predictions", "leave_one_out", database, "{}.{}.{}.LOO.prediction.csv".format(acc, database, s["setname"]))
        testdata_combined_file = os.path.join(s["features_folder"], "combined", database, "{}.surr20.gaps5.combined_features.csv".format(acc))
        thoipapy.utils.make_sure_path_exists(THOIPA_prediction_csv, isfile=True)
        #######################################################################################################
        #                                                                                                     #
        #      Train data is based on large training csv, after dropping the protein of interest              #
        #                   (positions without closedist/disruption data will be excluded)                    #
        #                                                                                                     #
        #######################################################################################################

        if not acc_db in acc_db_list:
            logging.warning("{} is in protein set, but not found in training data".format(acc_db))
            # skip protein
            continue
        # df_train = df_data.loc[df_data.acc_db != acc_db]
        # X_train = thoipapy.features.RF_Train_Test.drop_cols_not_used_in_ML(logging, df_train, s["excel_file_with_settings"])
        # y_train = df_train[s["bind_column"]]

        #######################################################################################################
        #                                                                                                     #
        #         Pack all variables into a dictionary compatible with multiprocessing Pool                   #
        #                                                                                                     #
        #######################################################################################################

        d = {}
        d["testdata_combined_file"], d["THOIPA_prediction_csv"] = testdata_combined_file, THOIPA_prediction_csv
        #d["X_train"], d["y_train"] = X_train, y_train
        d["df_data"] = df_data
        d["excel_file_with_settings"] = s["excel_file_with_settings"]
        d["forest"], d["pred_colname"] = forest, pred_colname
        d["i"], d["acc_db"], d["database"], d["bind_column"] = i, acc_db, database, s["bind_column"]
        d["n"], d["logger"], d["excel_file_with_settings"] = n, logger, s["excel_file_with_settings"]

        if s["use_multiprocessing"]:
            d_list.append(d)
        else:
            auc_dict, BO_df = LOO_single_prot(d)
            val_tuple = (auc_dict, BO_df)
            val_list.append(val_tuple)

    if s["use_multiprocessing"]:
        with Pool(processes=s["multiple_tmp_simultaneous"]) as pool:
            val_list = pool.map(LOO_single_prot, d_list)

    #######################################################################################################
    #                                                                                                     #
    #                            Get mean AUC etc for all proteins in list                                #
    #                                                                                                     #
    #######################################################################################################
    duration = time.clock() - start
    sys.stdout.write("\n")

    # copied from original mean_tpr code
    mean_tpr = 0.0
    mean_fpr = np.linspace(0, 1, 100)

    BO_all_df = pd.DataFrame()
    all_auc = []
    xv_dict = {}
    acc_db_list = df_set.acc_db.tolist()

    #iterate through the output tuple (auc_dict, BO_df)
    for nn, val_tuple in enumerate(val_list):
        acc_db = acc_db_list[nn]
        auc_dict = val_tuple[0]
        all_auc.append(auc_dict["auc"])
        BO_df = val_tuple[1]
        # join the data for all BO curves together
        if nn == 0:
            BO_all_df = BO_df
        else:
            BO_all_df = pd.concat([BO_all_df, BO_df], axis=1, join="outer")
        mean_tpr += interp(mean_fpr, auc_dict["fpr"], auc_dict["tpr"])
        mean_tpr[0] = 0.0

        xv_dict[acc_db] = {"fpr": auc_dict["fpr"], "tpr": auc_dict["tpr"], "auc": auc_dict["auc"]}

    # copied from original mean_tpr code
    mean_tpr /= df_set.shape[0]
    mean_tpr[-1] = 1.0
    mean_auc_from_joined_data = auc(mean_fpr, mean_tpr)

    # calculate mean of each protein AUC separately
    mean_auc_all_prot = np.array(all_auc).mean()
    xv_dict["mean_auc_all_prot"] = mean_auc_all_prot

    # add to dict that can be used for figure creation later
    xv_dict["true_positive_rate_mean"] = mean_tpr
    xv_dict["false_positive_rate_mean"] = mean_fpr
    xv_dict["mean_auc_from_joined_data"] = mean_auc_from_joined_data
    xv_dict["mean_auc_all_prot"] = mean_auc_all_prot

    # save dict as pickle
    thoipapy.utils.make_sure_path_exists(LOO_crossvalidation_pkl, isfile=True)
    with open(LOO_crossvalidation_pkl, "wb") as f:
        pickle.dump(xv_dict, f, protocol=pickle.HIGHEST_PROTOCOL)

    #######################################################################################################
    #                                                                                                     #
    #      Processing BO CURVE data, saving to csv and running the BO curve analysis script               #
    #                                                                                                     #
    #######################################################################################################

    BO_all_df.to_csv(BO_all_data_csv)
    #names_excel_path = os.path.join(os.path.dirname(s["sets_folder"]), "ETRA_NMR_names.xlsx")

    #linechart_mean_obs_and_rand = thoipapy.figs.Create_Bo_Curve_files.analyse_bo_curve_underlying_data(BO_all_data_csv, crossvalidation_folder, names_excel_path)
    thoipapy.figs.create_BOcurve_files.parse_BO_data_csv_to_excel(BO_all_data_csv, BO_data_excel, logging)

    logging.info('{} LOO crossvalidation. Time taken = {:.2f}.'.format(s["setname"], duration))
    logging.info('---AUC({:.2f})({:.2f})---'.format(mean_auc_all_prot, mean_auc_from_joined_data))

def LOO_single_prot(d):
    """Create Leave-One-Out cross-validation for a single protein in a dataset

    see docstring of run_LOO_validation

    Parameters
    ----------
    d : dict
        dictionary with all necessary values for the function
        having a single variable input allows python multiprocessing with Pool
    """
    testdata_combined_file, THOIPA_prediction_csv = d["testdata_combined_file"], d["THOIPA_prediction_csv"]
    #X_train, y_train = d["X_train"], d["y_train"]
    df_data, logger = d["df_data"], d["logger"]
    excel_file_with_settings = d["excel_file_with_settings"]
    forest, pred_colname = d["forest"], d["pred_colname"]
    i, acc_db, database, bind_column = d["i"], d["acc_db"], d["database"], d["bind_column"]

    #######################################################################################################
    #                                                                                                     #
    #                 Train data is based on the large file of all residues in the dataset                #
    #                       - positions without closedist/disruption data will be EXCLUDED)               #
    #                       - crystal positions with "folding/hetero contacts" will be EXCLUDED           #
    #                                                                                                     #
    #######################################################################################################

    df_train = df_data.loc[df_data.acc_db != acc_db]
    X_train = thoipapy.validation.validation.drop_cols_not_used_in_ML(logger, df_train, excel_file_with_settings, i)
    y_train = df_train[bind_column]

    n, logger, excel_file_with_settings = d["n"], d["logger"], d["excel_file_with_settings"]

    #######################################################################################################
    #                                                                                                     #
    #                  Test data is based on the combined features file for that protein and TMD          #
    #                              - positions without closedist/disruption data will be INCLUDED)        #
    #                              - positions with "folding/hetero contacts" will be included            #
    #                                                                                                     #
    #######################################################################################################
    # df_test = df_data.loc[df_data.acc_db == acc_db]
    df_test = pd.read_csv(testdata_combined_file)

    X_test = thoipapy.validation.validation.drop_cols_not_used_in_ML(logger, df_test, excel_file_with_settings, i)
    y_test = df_test["interface"].fillna(0).astype(int)

    #######################################################################################################
    #                                                                                                     #
    #                  Run prediction and save output individually for each protein                       #
    #                                                                                                     #
    #######################################################################################################

    fitted = forest.fit(X_train, y_train)
    if bind_column == "interface":
        prediction = fitted.predict_proba(X_test)[:, 1]
    elif bind_column == "interface_score_norm":
        prediction = fitted.predict(X_test)#[:, 1]
    else:
        raise ValueError("bind_column in excel settings file is not recognised ({})".format(bind_column))
    # add the prediction to the combined file
    df_test[pred_colname] = prediction
    # save just the prediction alone to csv
    prediction_df = df_test[["residue_num", "residue_name", pred_colname]]
    prediction_df.to_csv(THOIPA_prediction_csv, index=False)

    fpr, tpr, thresholds = roc_curve(y_test, prediction, drop_intermediate=False)
    roc_auc = auc(fpr, tpr)

    auc_dict = {"fpr": fpr, "tpr": tpr, "auc": roc_auc}

    if database == "crystal" or database == "NMR":
        # low closest distance means high importance at interface
        df_test["interface_score"] = -1 * df_test["interface_score"]

    BO_df = thoipapy.figs.fig_utils.calc_best_overlap(acc_db, df_test, experiment_col="interface_score", pred_col=pred_colname)

    if n == 0:
        tree_depths = np.array([estimator.tree_.max_depth for estimator in forest.estimators_])
        logger.info("tree depth mean = {} ({})".format(tree_depths.mean(), tree_depths))
    logger.info("{} AUC : {:.2f}".format(acc_db, roc_auc))

    return auc_dict, BO_df


def create_LOO_validation_fig(s, df_set, logging):
    """Create Leave-One-Out cross-validation for each TMD in a dataset.

    Training dataset = all residues in full dataset, except that being trained.

    ASSUMES YOUR DATASET IS NON-REDUNDANT! AUTO CD-HIT REDUNDANCY CHECKS ARE MOSTLY UNUSED!

    The model is trained on the train data csv (e.g. set38_train_data.csv)
         - for crystal subset, hetero contacts (folding residues) are removed from this training set!

    The model is validated against each combined CSV with features
     - for crystal subset, folding residues (hetero contacts) are INCLUDED here.
     - the model created without the folding contacts is therefore validated against the full seq, including folding residues

    Parameters
    ----------
    s : dict
        Settings dictionary
    df_set : pd.DataFrame
        Dataframe containing the list of proteins to process, including their TMD sequences and full-length sequences
        index : range(0, ..)
        columns : ['acc', 'seqlen', 'TMD_start', 'TMD_end', 'tm_surr_left', 'tm_surr_right', 'database',  ....]
    logging : logging.Logger
        Python object with settings for logging to console and file.
    """
    # drop redundant proteins according to CD-HIT
    df_set = thoipapy.utils.drop_redundant_proteins_from_list(df_set, logging)

    #plt.rcParams.update({'font.size': 7})
    LOO_crossvalidation_pkl = os.path.join(s["set_results_folder"], "crossvalidation", "data", "{}_LOO_crossvalidation.pkl".format(s["setname"]))
    LOO_crossvalidation_ROC_png = os.path.join(s["set_results_folder"], "crossvalidation", "{}_LOO_crossvalidation_ROC.png".format(s["setname"]))
    LOO_crossvalidation_AUC_bar_png = os.path.join(s["set_results_folder"], "crossvalidation", "{}_LOO_crossvalidation_AUC_bar.png".format(s["setname"]))
    AUC_csv = os.path.join(s["set_results_folder"], "crossvalidation", "data", "{}_LOO_AUC.csv".format(s["setname"]))
    BO_all_data_csv = os.path.join(s["set_results_folder"], "crossvalidation", "data", "{}_LOO_BO_data.csv".format(s["setname"]))
    BO_curve_folder = os.path.join(s["set_results_folder"], "crossvalidation")
    BO_data_excel = os.path.join(BO_curve_folder, "data", "{}_BO_curve_data.xlsx".format(s["setname"]))
    BO_linechart_png = os.path.join(BO_curve_folder, "{}_BO_linechart.png".format(s["setname"]))
    BO_barchart_png = os.path.join(BO_curve_folder, "{}_LOO_AUBOC10_barchart.png".format(s["setname"]))

    names_excel_path = os.path.join(os.path.dirname(s["sets_folder"]), "ETRA_NMR_names.xlsx")
    namedict = thoipapy.utils.create_namedict(names_excel_path)

    # open pickle file
    with open(LOO_crossvalidation_pkl, "rb") as f:
        xv_dict = pickle.load(f)

    # due to problems on Bo's computer, set the figsize to double what we should be using for the publication?
    figsize = np.array([3.42, 3.42]) * 2 # DOUBLE the real size, due to problems on Bo computer with fontsizes
    fig, ax = plt.subplots(figsize=figsize)
    auc_dict = {}

    for acc_db in df_set.acc_db:
        if acc_db in xv_dict:
            roc_auc = xv_dict[acc_db]["auc"]
            auc_dict[acc_db] = roc_auc
            ax.plot(xv_dict[acc_db]["fpr"], xv_dict[acc_db]["tpr"], lw=1, label='{} ({:.2f})'.format(acc_db, roc_auc), alpha=0.8)
        else:
            logging.warning("{} not in xv_dict after LOO validation".format(acc_db))

    mean_auc_all_prot = xv_dict["mean_auc_all_prot"]

    ax.plot(xv_dict["false_positive_rate_mean"], xv_dict["true_positive_rate_mean"], color="k", label='mean (area = %0.2f)' % mean_auc_all_prot, lw=1.5)
    ax.plot([0, 1], [0, 1], '--', color=(0.6, 0.6, 0.6), label='random')
    ax.set_xlim([-0.05, 1.05])
    ax.set_ylim([-0.05, 1.05])
    ax.set_xlabel("False positive rate")
    ax.set_ylabel("True positive rate")
    ax.legend(loc="lower right")
    fig.tight_layout()
    fig.savefig(LOO_crossvalidation_ROC_png, dpi=240)
    #fig.savefig(LOO_crossvalidation_ROC_png[:-4] + ".pdf")
    fig.savefig(thoipapy.utils.pdf_subpath(LOO_crossvalidation_ROC_png))

    AUC_ser = pd.Series(auc_dict)
    AUC_ser.sort_values(inplace=True, ascending=False)
    AUC_ser.to_csv(AUC_csv)
    plt.close("all")
    figsize = np.array([3.42, 3.42]) * 2 # DOUBLE the real size, due to problems on Bo computer with fontsizes
    fig, ax = plt.subplots(figsize=figsize)
    AUC_ser.plot(ax=ax, kind="bar")
    ax.set_ylabel("performance (AUBOC10)")
    fig.tight_layout()
    fig.savefig(LOO_crossvalidation_AUC_bar_png, dpi=240)
    fig.savefig(thoipapy.utils.pdf_subpath(LOO_crossvalidation_AUC_bar_png))

    AUBOC10 = thoipapy.figs.create_BOcurve_files.save_BO_linegraph_and_barchart(s, BO_data_excel, BO_linechart_png, BO_barchart_png, namedict, logging, AUC_ser)

    if "TRUE" == "TRUE":
        other_figs_path = os.path.join(BO_curve_folder, "other_figs")
        thoipapy.utils.make_sure_path_exists(other_figs_path)
        thoipapy.figs.create_BOcurve_files.save_extra_BO_figs(BO_data_excel, other_figs_path)

    logging.info('{} LOO crossvalidation. AUBOC10({:.2f}).'.format(s["setname"], AUBOC10))
    logging.info("create_LOO_validation_fig finished ({})".format(LOO_crossvalidation_AUC_bar_png))

def calculate_variable_importance(s, logging):
    """Calculate the variable importance (mean decrease gini) for all variables in THOIPA.

    Parameters
    ----------
    s : dict
        Settings dictionary
    logging : logging.Logger
        Python object with settings for logging to console and file.

    Saved Files
    -----------
    variable_importance_csv : csv
        List of variables, sorted by their importance to the algorithm.
        Also includes the standard deviation supplied by the machine learning algorithm
    """
    #logging.info('RF_variable_importance_calculate is running\n')
    train_data_csv = os.path.join(s["set_results_folder"], "{}_train_data.csv".format(s["setname"]))

    variable_importance_csv = os.path.join(s["set_results_folder"], "crossvalidation", "data", "{}_variable_importance.csv".format(s["setname"]))
    thoipapy.utils.make_sure_path_exists(variable_importance_csv, isfile=True)

    df_data = pd.read_csv(train_data_csv, index_col=0)
    df_data = df_data.dropna()

    X = drop_cols_not_used_in_ML(logging, df_data, s["excel_file_with_settings"])
    y = df_data["interface"]
    n_features = X.shape[1]

    # regular trees, or Totally Randomized Trees
    model_types = ["", "_TRT"]
    output_dfs = []

    for model_type in model_types:
        if model_type == "_TRT":
            forest = THOIPA_classifier_with_settings(s, n_features, totally_randomized_trees=True)
            logging.info("IMPORTANCES FOR TOTALLY RANDOMIZED TREES (max_features=1, max_depth=None, min_samples_leaf=1)")
        elif model_type == "":
            forest = THOIPA_classifier_with_settings(s, n_features)
            logging.info("Feature ranking:")
        else:
            raise ValueError("model type unknown")
        forest.fit(X, y)
        importances_arr = forest.feature_importances_
        std_arr = np.std([tree.feature_importances_ for tree in forest.estimators_], axis=0)
        indices_arr = np.argsort(importances_arr)[::-1]

        importances_text_list = X.columns.tolist()

        #order_list = [importances_arr[indices_arr[f]] for f in range(X.shape[1])]



        nested_dict = {}

        for f in range(X.shape[1]):
            if f < 10 :
                logging.info("%d. feature %d (%f) %s" % (f + 1, indices_arr[f], importances_arr[indices_arr[f]], importances_text_list[indices_arr[f]]))
            single_feature_dict = {"original_order" : indices_arr[f], "mean_decrease_gini{}".format(model_type) : importances_arr[indices_arr[f]], "feature{}".format(model_type) : importances_text_list[indices_arr[f]],  "std{}".format(model_type) : std_arr[f]}
            nested_dict[f + 1] = single_feature_dict

        sys.stdout.write("\n\n"), sys.stdout.flush()

        df_imp = pd.DataFrame(nested_dict).T
        df_imp["order_importance{}".format(model_type)] = df_imp.index
        #df_imp.set_index("feature", inplace=True)
        df_imp.set_index("original_order", inplace=True)
        output_dfs.append(df_imp)

    df_imp2 = pd.concat(output_dfs, axis=1)
    df_imp2.sort_values("order_importance", inplace=True)
    df_imp2["original_order"] = df_imp2.index
    df_imp2.set_index("feature", inplace=True)

    df_imp2.to_csv(variable_importance_csv)

def fig_variable_importance(s, logging):
    """Create figures showing ML feature importance.

    Fig1 : Barchart all features
    Fig2 : Barchart top features (currently set at 30)

    Parameters
    ----------
    s : dict
        Settings dictionary
    logging : logging.Logger
        Python object with settings for logging to console and file.
    """
    plt.style.use('seaborn-whitegrid')
    plt.rcParams['errorbar.capsize'] = 1
    plt.rcParams.update({'font.size':4})
    #from thoipapy.utils import create_colour_lists
    from thoipapy.utils import create_colour_lists
    colour_dict = create_colour_lists()

    variable_importance_csv = os.path.join(s["set_results_folder"], "crossvalidation", "data", "{}_variable_importance.csv".format(s["setname"]))
    variable_importance_all_png = os.path.join(s["set_results_folder"], "crossvalidation", "{}_all_var_import.png".format(s["setname"]))
    variable_importance_top_png = os.path.join(s["set_results_folder"], "crossvalidation", "{}_top_var_import.png".format(s["setname"]))

    df_imp = pd.read_csv(variable_importance_csv, index_col = 0)

    create_var_imp_plot(df_imp, colour_dict, variable_importance_all_png, df_imp.shape[0])
    create_var_imp_plot(df_imp, colour_dict, variable_importance_top_png, 30)


def create_var_imp_plot(df_imp, colour_dict, variable_importance_png, n_features_in_plot):
    """Plot function for fig_variable_importance, allowing a variable number of features.
    """
    df_sel = df_imp.iloc[:n_features_in_plot, :].copy()

    # regular trees, or Totally Randomized Trees
    model_types = ["", "_TRT"]

    for model_type in model_types:

        # add suffix for the totally randomised trees
        variable_importance_png = variable_importance_png[:-4] + model_type + ".png"

        df_sel.sort_values("mean_decrease_gini{}".format(model_type), ascending=True, inplace=True)
        #min_ = df_sel.mean_decrease_gini.min()

        # determine the plot height by the number of features
        # currently set for 30
        plot_height = 4 * n_features_in_plot / 30
        figsize = np.array([4.42, plot_height])
        fig, ax = plt.subplots(figsize=figsize)

        TUMblue = colour_dict["TUM_colours"]['TUMBlue']
        df_sel["mean_decrease_gini{}".format(model_type)].plot(kind="barh", color=TUMblue, ax=ax)# xerr=df_sel["std"]
        ax.errorbar(df_sel["mean_decrease_gini{}".format(model_type)], range(len(df_sel.index)), xerr=df_sel["std{}".format(model_type)], fmt="none", ecolor="k", ls="none", capthick=0.5, elinewidth=0.5, capsize=1, label=None)

        ax.set_ylabel("")
        ax.set_xlabel("variable importance\n(mean decrease gini)")
        ax.grid(False)
        fig.tight_layout()
        fig.savefig(variable_importance_png, dpi=240)
        #fig.savefig(variable_importance_png[:-4] + ".pdf")
        fig.savefig(thoipapy.utils.pdf_subpath(variable_importance_png))

