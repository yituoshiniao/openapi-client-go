FROM registry.intsig.net/cs_healthy/base:v1




RUN mkdir -p /apps/webroot/production/goods-center-logic \
&& mkdir  -p /apps/conf/goods-center-logic   \
&& mkdir  -p /apps/log/app/goods-center-logic

#拷贝文件<源文件  目标>
ADD app /apps/webroot/production/goods-center-logic/app
RUN chmod +x /apps/webroot/production/goods-center-logic/app \
&& touch /apps/log/app/goods-center-logic/test.txt \
&& chmod -R +rwx /apps/log/app/goods-center-logic


#工作目录
WORKDIR /apps/webroot/production/goods-center-logic

COPY config/* /apps/conf/goods-center-logic/

ENV PATH=/apps/webroot/production/goods-center-logic:$PATH

#ARG GOTMP=test111

ENV envPath = ""
#ENV input1=""

#测试环境

#不支持动态命令替换
#ENTRYPOINT ["/apps/webroot/production/goods-center-logic/app" , "-envPath=$envPath"]

#shell 模式 支持动态命令替换
#CMD /apps/webroot/production/goods-center-logic/app -envPath=${envPath}
ENTRYPOINT /apps/webroot/production/goods-center-logic/app -envPath=${envPath}

#ENTRYPOINT ["tail", "-f", "/apps/log/app/goods-center-logic/test.txt"]


EXPOSE 3013 6013