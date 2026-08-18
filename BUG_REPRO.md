# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

创建班次时，结束时间跨过上海业务日边界仍返回成功，导致日结统计漏掉班次。请修复班次创建的业务日窗口校验。

## 含 Bug 版本

- 仓库：zhanglei10281852-gif/t63-qa-27
- 仓库地址：https://github.com/zhanglei10281852-gif/t63-qa-27.git
- parent SHA：685df8e80586c3feef6a96146ec9d7a56b6aa0e6

## 复现步骤

```bash
git clone -- https://github.com/zhanglei10281852-gif/t63-qa-27.git bug-repro
cd bug-repro
git checkout --detach 685df8e80586c3feef6a96146ec9d7a56b6aa0e6
go test ./internal/httpapi -run TestShiftMustRemainInsideSelectedServiceDay -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestShiftMustRemainInsideSelectedServiceDay -count=1
--- FAIL: TestShiftMustRemainInsideSelectedServiceDay (0.61s)
    planning_window_test.go:12: POST /api/v1/shifts status=409 want=400 body={"code":"conflict","message":"constraint failed: UNIQUE constraint failed: shifts.route_id, shifts.service_date (2067)","request_id":"request-test"}
FAIL
FAIL	sanitation-operations/internal/httpapi	0.611s
FAIL

```

stderr：

```text
warning: internal/httpapi/planning_window_test.go has type 100755, expected 100644
warning: internal/httpapi/server_test.go has type 100755, expected 100644
warning: internal/httpapi/planning_window_test.go has type 100755, expected 100644
warning: internal/httpapi/server_test.go has type 100755, expected 100644

```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestShiftMustRemainInsideSelectedServiceDay -count=1
--- FAIL: TestShiftMustRemainInsideSelectedServiceDay (1.06s)
    planning_window_test.go:12: POST /api/v1/shifts status=409 want=400 body={"code":"conflict","message":"constraint failed: UNIQUE constraint failed: shifts.route_id, shifts.service_date (2067)","request_id":"request-test"}
FAIL
FAIL	sanitation-operations/internal/httpapi	1.288s
FAIL

```

stderr：

```text
warning: internal/httpapi/planning_window_test.go has type 100755, expected 100644
warning: internal/httpapi/server_test.go has type 100755, expected 100644
warning: internal/httpapi/planning_window_test.go has type 100755, expected 100644
warning: internal/httpapi/server_test.go has type 100755, expected 100644

```

## 通过条件

在触发条件下，定向测试 TestShiftMustRemainInsideSelectedServiceDay 应通过，相关包、全量测试、竞态测试和构建检查均通过；回退 gold 唯一修复后定向测试重新失败。
