'''
DataFrame接口
'''
# -*- coding: utf-8 -*-
from lixinger_openapi.baseapi import(
    STOCK_FUNDAMENTAL_INFO_URL,
    STOCK_FS_INFO_URL,
    INDICE_FUNDAMENTAL_INFO_URL,
    basic_query_info_df,
)

def stock_fundamental_info(date=None, startDate=None, endDate=None, stockCodes=None, metrics=None):
    '''
    获取公司基本面数据，返回DataFrame结构
    '''
    return basic_query_info_df(STOCK_FUNDAMENTAL_INFO_URL, date, startDate, endDate, stockCodes, metrics)

def stock_fs_info(date=None, startDate=None, endDate=None, stockCodes=None, metrics=None):
    '''
    获取公司财务数据，返回DataFrame结构
    '''
    return basic_query_info_df(STOCK_FS_INFO_URL, date, startDate, endDate, stockCodes, metrics)

def indice_fundamental_info(date=None, startDate=None, endDate=None, stockCodes=None, metrics=None):
    '''
    获取指数基本面数据，返回DataFrame结构
    '''
    return basic_query_info_df(INDICE_FUNDAMENTAL_INFO_URL, date, startDate, endDate, stockCodes, metrics)
