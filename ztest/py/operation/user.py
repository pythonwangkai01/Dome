import json
from core.rest_base import ResultBase
from api.user import user
from common.logger import logger

def register_user(username,password,passwordconfirm,phone,sex="",address=""):
    """
    注册用户信息
    :param username: 用户名
    :param password: 密码
    :param passwordconfirm: 确定密码
    :param phone: 手机号
    :param sex: 性别
    :param address: 联系地址
    :return: 自定义的关键字返回结果 result
    """
    result = ResultBase()
    json_data = {
        "UserName": username,
        "Password": password,
        "PasswordConfirm":passwordconfirm,
        "Sex": sex,
        "Phone": phone,
        "Address": address
    }
    hander = {
        "content-type": "application/json"
    }
    res = user.register(json=json_data,handers=hander)
    result.success = False
    if res.json()["code"] == 200:
        result.success = True
    else:
        result.error = "接口返回码是 【 {} 】, 返回信息：{} ".format(res.json()["code"], res.json()["msg"])
    result.msg = res.json()["msg"]
    result.response = res
    logger.info("注册用户 ==>> 返回结果 ==>> {}".format(result.response.text))
    return result

def login_user(username, password):
    """
    登录用户
    :param username: 用户名
    :param password: 密码
    :return: 自定义的关键字返回结果 result
    """
    result = ResultBase()
    payload = {
        "username": username,
        "password": password
    }
    header = {
        "content-type": "application/json"
    }
    res = user.login(data=payload, headers=header)
    result.success = False
    if res.json()["code"] == 200:
        result.success = True
        result.token = res.json()["data"]["token"]
    else:
        result.error = "接口返回码是 【 {} 】, 返回信息：{} ".format(res.json()["code"], res.json()["msg"])
    result.msg = res.json()["msg"]
    result.response = res
    logger.info("登录用户 ==>> 返回结果 ==>> {}".format(result.response.text))
    return result

def delete_user(username, admin_user, token):
    """
    根据用户名，删除用户信息
    :param username: 用户名
    :param admin_user: 当前操作的管理员用户
    :param token: 当前管理员用户的token
    :return: 自定义的关键字返回结果 result
    """
    result = ResultBase()
    json_data = {
        "admin_user": admin_user,
        "token": token,
    }
    header = {
        "Content-Type": "application/json"
    }
    res = user.delete(username, json=json_data, headers=header)
    result.success = False
    if res.json()["code"] == 0:
        result.success = True
    else:
        result.error = "接口返回码是 【 {} 】, 返回信息：{} ".format(res.json()["code"], res.json()["msg"])
    result.msg = res.json()["msg"]
    result.response = res
    logger.info("删除用户 ==>> 返回结果 ==>> {}".format(result.response.text))
    return result