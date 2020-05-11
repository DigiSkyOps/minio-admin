FROM node:10.15.3
RUN npm install -g cnpm --registry=https://registry.npm.taobao.org
WORKDIR /data/webapps
COPY . .
ARG PLATFORM_URL
ENV PLATFORM_URL $PLATFORM_URL
RUN yarn && yarn build
EXPOSE 8000

CMD ["sh","+x","/data/webapps/run.sh"]

