
from urllib.parse import quote
import urllib;
import base64
import hmac
from hashlib import sha256

"""
使用getSignature 获取签名
"""


#构造URL请求参数部分
def getSignatureArgs(str_map):
    print(str_map +'\n')
    # 字典排序
    str_url = sort_dict(str_map);
    print(str_url +'\n')

    #构造url
    str_url = urllib.parse.urlencode(str_url);
    print(str_url +'\n')

    #对参数名称和参数值进行URL编码
    # str_url = percentEncode(str_url)
    return  str_url;


def getSignatureStr(method,uri,url):
    """
    #构造被签名的串，被签名串 = HTTP请求方式 + “\n” + uri + “\n” + url 请求串
    :param method:
    :param uri:
    :param url:
    :return:
    """
    str = method + "\n" + uri + "\n"+ url;
    return  str;

def percentEncode(str):
    res = urllib.parse.quote(str, '')
    res = res.replace('+', '%20')
    res = res.replace('*', '%2A')
    res = res.replace('%7E', '~')
    return res


def sort_dict(dict_words):
    """
    字典排序
    :param dict_words:
    :return:
    """
    return dict(sorted(dict_words.items(), key=lambda d: d[0]))




def getSignature(method,uri,url_map,secret_access_key):
    """
    获取签名
    :param method:请求方法
    :param uri:请求的uri路径
    :param url_map:请求参数字典
    :param secret_access_key:私钥
    :return:签名
    """

    url =  getSignatureArgs(url_map);
    string_to_sign = getSignatureStr(method,uri,url);

    sb = bytes(secret_access_key, encoding="utf8")

    h = hmac.new(sb, digestmod=sha256)
    h.update(string_to_sign.encode("utf8"))
    sign = base64.b64encode(h.digest()).strip()
    signature = urllib.parse.quote_plus(sign)
    return  signature;



def test():
    """
    测试的
    :return:
    """
    url_map = {
      "count":1,
      "vxnets.1":"vxnet-0",
      "zone":"pek1",
      "instance_type":"small_b",
      "signature_version":1,
      "signature_method":"HmacSHA256",
      "instance_name":"demo",
      "image_id":"centos64x86a",
      "login_mode":"passwd",
      "login_passwd":"QingCloud20130712",
      "version":1,
      "access_key_id":"QYACCESSKEYIDEXAMPLE",
      "action":"RunInstances",
      "time_stamp":"2013-08-27T14:30:10Z"
    };
    s_url = 'GET\n/iaas/\naccess_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1'

    s_s = '32bseYy39DOlatuewpeuW5vpmW51sD1A%2FJdGynqSpP8%3D';
    method = 'GET';
    uri = "/iaas/";
    secret_access_key = 'SECRETACCESSKEY'
    print(getSignature(method,uri,url_map,secret_access_key) == s_s)