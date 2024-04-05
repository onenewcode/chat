#!/bin/bash
set -e # 这一行设置了Shell脚本的退出状态行为。-e选项使得脚本在遇到任何命令返回非零退出状态时立即终止执行。
if [ "$ENV" = 'DEV' ]; then
 echo "Running Development Server" # 打印字符串
# exec python "identidock.py" # 它将整个Shell脚本进程替换为运行Python解释器，并指定要执行的脚本文件为 "identidock.py"
else
 ./main
# exec uwsgi --http 0.0.0.0:9090 --wsgi-file /app/identidock.py \
# --callable app --stats 0.0.0.0:9191
fi