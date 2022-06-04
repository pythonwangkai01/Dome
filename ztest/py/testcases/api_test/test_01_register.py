import pytest
import allure
from operation.user import register_user
from testcases.conftest import api_data
from common.logger import logger

@allure.step("步骤1 ==>> 注册用户")
def step_1(username, password, phone, sex, address):
    logger.info("步骤1 ==>> 注册用户 ==>> {}, {}, {}, {}, {}".format(username, password, phone, sex, address))


"""
@allure.severity:提供了划分等级的功能，且可以展示到测试报告内
@allure.epic:敏捷里面的概念，定义史诗，相当于module级的标签
@allure.feature:功能点的描述，可以理解成模块，相当于class级的标签
"""
@allure.severity(allure.severity_level.NORMAL)
@allure.epic("针对单个接口的测试")
@allure.feature("用户注册模块")
class TestUserRegister():
    """
    用户注册
    @allure.story:可以理解为测试场景
    @allure.description:可以添加足够详细的测试用例描述
    @allure.link()表示访问网址的链接
    @allure.issue()表示bug的链接
    @allure.testcase()表示测试用例的链接
    @allure.title:
    1.使得测试用例的标题更具有可读性，毕竟我们可以写成中文
    2.支持占位符传递关键字参数哦
    @pytest.mark.parametrize(args_name,args_value):
    其实就是把我们测试用例的数据放到excel，yaml，csv，mysql，
    然后通过去改变数据达到改变测试用例的执行结果 。
    @pytest.mark.usefixtures:的使用场景是被测试函数需要多个fixture做前后置工作时使用
    """
    @allure.story("用例--注册用户信息")
    @allure.description("该用例是针对获取用户注册接口的测试")
    # @allure.issue("https://www.cnblogs.com/wintest", name="点击，跳转到对应BUG的链接地址")
    # @allure.testcase("https://www.cnblogs.com/wintest", name="点击，跳转到对应用例的链接地址")
    @allure.title(
        "测试数据：【 {username}，{password}，{telephone}，{sex}，{address}，{except_result}，{except_code}，{except_msg}】")
    @pytest.mark.smoke
    @pytest.mark.parametrize("username, password, telephone, sex, address, except_result, except_code, except_msg",
                             api_data["test_register_user"])
    @pytest.mark.usefixtures("delete_register_user")
    def test_register_user(self,username, password, telephone, sex, address, except_result, except_code, except_msg):
        logger.info("*************** 开始执行用例 ***************")
        result = register_user(username, password, telephone, sex, address)
        step_1(username, password, telephone, sex, address)
        assert result.success == except_result
        assert result.response.status_code == except_code
        assert result.msg == except_msg
        logger.info("code ==>> 期望结果：{}， 实际结果：【 {} 】".format(except_code, result.response.json().get("code")))
        assert result.response.json().get("code") == except_code
        assert except_msg in result.msg
        logger.info("*************** 结束执行用例 ***************")
if __name__ == '__main__':
    pytest.main(["-q", "-s", "test_01_register.py"])