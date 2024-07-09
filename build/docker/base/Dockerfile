# Copyright 2024 Innkeeper Base <https://www.base.cn>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://www.base.cn


FROM BASE_IMAGE
LABEL maintainer="<my@dengmengmian.com>"

WORKDIR /opt/base

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone && \
      mkdir -p /var/log/base

COPY base /opt/base/bin/

ENTRYPOINT ["/opt/base/bin/base"]
CMD ["-c", "/etc/base/base.yaml"]
