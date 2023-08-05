
import os
import json
import copy

import numpy as np
import datetime as dt
import pandas as pd

from ._util import Util

from IPython.display import Markdown
from sgmarkets_api_auth.util import save_result


class ResponseRotbComputeStrategyComponents:
    """
    """

    def __init__(self,
                 li_raw_data=None,
                 obj_req=None):
        """
        """
        assert isinstance(li_raw_data, list), \
            'Error: li_raw_data must be a list - Run call_api() again with debug=True'
        for dic_res in li_raw_data:
            assert isinstance(dic_res, dict), \
                'Error: Each dic_res must be a list - Run call_api() again with debug=True'
            assert 'componentSeries' in dic_res, \
                'Error: componentSeries must be a key of each dic_res - Run call_api() again with debug=True'

        raw_data = []
        for dic_res in li_raw_data:
            raw_data += dic_res['componentSeries']

        self.raw_data = copy.deepcopy(raw_data)
        self.obj_req = obj_req

        self.df_req, self.df_res, self.df_set = self._build_df_res_req()
        self.dic_req_param, self.dic_res_param = self._build_dic_param()

    def _get_dates(self, df_res):
        """
        """
        dic = self.obj_req.df_top.to_dict()
        dic = dic['Value']

        if 'dates' in dic:
            res = dic['dates'].replace("'", '"')
            return json.loads(res)

        return Util.get_unique_list(df_res['date'])

    def _build_df_res_req(self):
        """
        """
        # df_res (response)
        li_data = [f for e in self.raw_data for f in e]
        for e in li_data:
            if 'greeks' in e:
                for greek in ['delta', 'gamma', 'vega', 'theta']:
                    e[greek] = e['greeks'][greek]
                e.pop('greeks')
        df_res = pd.DataFrame(li_data)

        # build list of dates
        li_date = self._get_dates(df_res)
        N = len(li_date)

        # df_req (request)
        # the order of results is by order of input
        # for each input the order of dates - but this is changed below
        # duplicate df_leg by number of dates
        df_leg = self.obj_req.df_leg
        df_req = pd.concat([df_leg] * N,
                           axis=0).reset_index(drop=True)

        # reorder results by date then initial order (tag)
        df_res['tag'] = range(len(df_res))
        df_res = df_res.sort_values(['date', 'tag']).reset_index(drop=True)

        # move date from df_res to df_req (more natural)
        df_req['date'] = pd.to_datetime(df_res['date'].copy())
        df_res = df_res.drop('date', axis=1)

        df_res = df_res.rename(columns={
            'strike': 'strike_res',
            'nominal': 'nominal_res',
        })

        if 'error' not in df_res:
            df_res['error'] = 'No error'
        else:
            df_res['error'] = df_res['error'].fillna('No error')

        # move col error to last position
        cols = [c for c in df_res.columns if c != 'error']+['error']
        df_res = df_res[cols]

        # replace NaN returned by API
        df_res = df_res.replace('NaN', np.nan)

        # join df_req and df_res to make df_set
        df_set = pd.concat([df_req, df_res], axis=1)

        return df_req, df_res, df_set

    def _build_dic_param(self):
        """
        """
        dic_req = self.df_req.to_dict()
        dic_req_param = {k: Util.get_unique_list(v.values())
                         for k, v in dic_req.items()}

        dic_data = self.df_res.to_dict()
        dic_res_param = {k: Util.get_unique_list(v.values())
                         for k, v in dic_data.items()}

        return dic_req_param, dic_res_param

    def save(self,
             folder_save='dump',
             name=None,
             tagged=True,
             excel=False):
        """
        """
        if name is None:
            name = 'SG_Research_ROTB'

        save_result(self.df_set,
                    folder_save, name=name + '_Components_response',
                    tagged=tagged,
                    excel=excel)

    def _repr_html_(self):
        """
        """
        return self.df_res.to_html()

    def info(self):
        """
        """
        md = """
A PostprocessROTB object from ComputeStrategyComponents endpoint has the properties:
+ `df_req`: request data (dataframe)
+ `df_res`: response data (dataframe)
+ `df_set`: request and response data combined (dataframe)

+ `dic_req_param`: params in request, each param contains a list of all values taken (dictionary)
+ `dic_res_param`: params in response, each param contains a list of all values taken (dictionary)

+ `raw_data`: raw data in response under key 'componentSeries' (dictionary)

and the methods:
+ `save()` to save the data as `.csv` and `.xlsx` files
        """
        return Markdown(md)
