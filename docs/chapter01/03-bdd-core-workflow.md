# BDD 核心工作流程：从对话到可执行规格

## 概述

Behavior-Driven Development (BDD) 的核心是通过协作和具体范例来建立共同理解，并将这些理解转化为可执行的规格。本文档详细说明 BDD 的核心工作流程，重点关注团队如何通过对话、范例和自动化测试来交付高质量的软件。

## BDD 工作流程概览

```
1. 业务对话        2. 范例探索        3. 规格转化        4. 测试自动化       5. 活文档
   (Discovery)       (Examples)         (Formalize)        (Automate)         (Living Doc)
       ↓                 ↓                  ↓                  ↓                   ↓
   Feature          Concrete           Gherkin          Automated           Product
   Capability       Scenarios          Specs            Tests               Documentation
```

## 阶段 1: 业务对话 - 理解 Feature 和 Capability

### 目的
通过协作对话，让业务代表、开发人员和测试人员建立对需求的共同理解。

### 参与者：Three Amigos

| 角色 | 职责 | 关注点 |
|------|------|--------|
| **业务代表** (Business) | 定义业务价值和需求 | What needs to be built? Why? |
| **开发人员** (Developer) | 评估技术可行性 | How will it be built? What constraints? |
| **测试人员** (Tester) | 识别边界和风险 | What could go wrong? What edge cases? |

### 对话框架：Feature 定义

使用用户故事格式来定义 feature：

```gherkin
Feature: [Feature Name]
  As a [role/persona]
  I want [capability/feature]
  So that [business value/benefit]
```

### 实践案例：转账功能

#### 初始对话

**业务代表**：「我们需要让客户能在自己的帐户间转账。这是客户最常要求的功能之一。」

**开发人员**：「转账涉及两个帐户，我们需要确保资金安全地从一个帐户移到另一个。这需要交易管理。」

**测试人员**：「如果转账过程中发生错误怎么办？余额不足时该如何处理？我们需要考虑各种失败情况。」

#### Feature 定义

```gherkin
Feature: 帐户转账
  As a 银行客户
  I want to 在我的帐户之间转账
  So that 我可以灵活管理我的资金
```

### 探索 Capability 的关键问题

#### 业务价值问题
- 这个 feature 为谁解决什么问题？
- 成功的标准是什么？
- 有哪些业务规则必须遵守？

#### 技术可行性问题
- 需要哪些系统和服务？
- 有什么技术限制或依赖？
- 性能要求是什么？

#### 质量和风险问题
- 什么情况下功能应该失败？
- 边界条件是什么？
- 安全和合规要求？

### 对话产出

- ✅ 清晰的 Feature 描述
- ✅ 明确的业务价值
- ✅ 识别的风险和约束
- ✅ 准备好进入下一阶段的共识

## 阶段 2: 范例探索 - 具体化需求

### 目的
通过具体的范例和反例，将抽象的需求转化为可验证的场景。

### Example Mapping 技巧

Example Mapping 是一种结构化的协作技术，用于探索和组织范例。

```
┌─────────────────────────────┐
│     Feature/User Story      │  (黄色卡片)
└─────────────────────────────┘
         ↓
┌─────────────────────────────┐
│      Business Rules         │  (蓝色卡片)
└─────────────────────────────┘
         ↓
┌─────────────────────────────┐
│     Concrete Examples       │  (绿色卡片)
└─────────────────────────────┘
         ↓
┌─────────────────────────────┐
│    Questions/Unknowns       │  (红色卡片)
└─────────────────────────────┘
```

### 实践案例：转账功能的范例探索

#### Business Rule 1: 余额验证

**规则**：转账金额不能超过来源帐户的可用余额

**正面范例** (成功场景):
```
范例 1: 余额充足的转账
Given Chris 的支票帐户有 $1,000
And Lisa 的储蓄帐户有 $500
When Chris 转账 $300 到 Lisa 的帐户
Then Chris 的帐户应该有 $700
And Lisa 的帐户应该有 $800
And 转账应该成功完成
```

**反面范例** (失败场景):
```
范例 2: 余额不足
Given Chris 的支票帐户有 $100
When Chris 尝试转账 $150 到 Lisa 的帐户
Then 转账应该被拒绝
And Chris 应该看到错误信息 "余额不足"
And Chris 的帐户仍然有 $100
And Lisa 的帐户余额不变
```

#### Business Rule 2: 金额验证

**规则**：转账金额必须大于零

**反面范例**:
```
范例 3: 零金额转账
Given Chris 的帐户有 $1,000
When Chris 尝试转账 $0
Then 转账应该被拒绝
And 错误信息应该是 "金额必须大于零"

范例 4: 负金额转账
Given Chris 的帐户有 $1,000
When Chris 尝试转账 $-50
Then 转账应该被拒绝
And 错误信息应该是 "金额不能为负数"
```

#### Business Rule 3: 帐户验证

**规则**：不能转账到相同的帐户

**反面范例**:
```
范例 5: 转账到自己的帐户
Given Chris 的帐户编号是 "ACC001"
When Chris 尝试从 "ACC001" 转账到 "ACC001"
Then 转账应该被拒绝
And 错误信息应该是 "不能转账到相同帐户"
```

#### 边界情况和边缘案例

```
范例 6: 最小金额转账
Given Chris 的帐户有 $100
When Chris 转账 $0.01 到 Lisa
Then 转账应该成功
And Chris 应该有 $99.99
And Lisa 应该收到 $0.01

范例 7: 全额转账
Given Chris 的帐户有 $100
When Chris 转账 $100 到 Lisa
Then 转账应该成功
And Chris 的帐户余额应该是 $0
```

### 范例的特征

#### ✅ 好的范例
- **具体**: 使用真实的名字、金额和数据
- **可验证**: 有明确的输入和预期输出
- **独立**: 不依赖其他范例的结果
- **简洁**: 专注于一个业务规则
- **业务语言**: 避免技术术语

#### ❌ 不好的范例
- **抽象**: "用户转账一些金额"
- **模糊**: "系统应该正确处理"
- **依赖**: "基于前一个范例的结果"
- **复杂**: 测试多个不相关的规则
- **技术化**: "调用 TransferService.execute()"

### 探索技巧

#### 1. 使用 "What about...?" 问题
- "如果帐户被冻结怎么办？"
- "如果是不同货币怎么办？"
- "如果在转账过程中网络断开怎么办？"

#### 2. 边界值分析
- 最小值：$0.01
- 最大值：帐户余额
- 零值：$0
- 负值：$-100

#### 3. 状态转换
- 转账前的状态
- 转账过程中的状态
- 转账成功后的状态
- 转账失败后的状态

## 阶段 3: 规格形式化 - 编写 Gherkin

### 目的
将探索出的范例转化为结构化的 Gherkin 规格，作为自动化测试的基础。

### Gherkin 语法结构

```gherkin
Feature: [功能名称]
  [用户故事或功能描述]

  Background:
    [所有场景共享的前置条件]

  Scenario: [场景名称]
    Given [前置条件 - 设置测试上下文]
    And [更多前置条件]
    When [执行动作 - 触发被测行为]
    And [更多动作]
    Then [验证结果 - 断言预期行为]
    And [更多验证]

  Scenario Outline: [参数化场景名称]
    Given [带有 <参数> 的步骤]
    When [带有 <参数> 的步骤]
    Then [带有 <参数> 的步骤]

    Examples:
      | 参数1 | 参数2 | 预期结果 |
      | 值1   | 值2   | 结果1    |
      | 值3   | 值4   | 结果2    |
```

### 实践案例：完整的 Feature 文件

**文件**: `features/money_transfer.feature`

```gherkin
Feature: 帐户转账
  As a 银行客户
  I want to 在我的帐户之间转账
  So that 我可以灵活管理我的资金

  Background:
    Given 存在以下帐户:
      | 帐户编号 | 户名   | 余额    |
      | ACC001   | Chris  | 1000.00 |
      | ACC002   | Lisa   | 500.00  |
      | ACC003   | Jordan | 0.00    |

  # 主要成功场景
  Scenario: 成功转账 - 余额充足
    Given Chris 已登录
    When Chris 从帐户 "ACC001" 转账 $300 到帐户 "ACC002"
    Then 转账应该成功
    And 帐户 "ACC001" 的余额应该是 $700
    And 帐户 "ACC002" 的余额应该是 $800
    And Chris 应该看到确认信息 "转账成功"
    And 应该记录一笔转账交易

  # 余额不足场景
  Scenario: 转账失败 - 余额不足
    Given Chris 已登录
    When Chris 从帐户 "ACC001" 转账 $1500 到帐户 "ACC002"
    Then 转账应该失败
    And Chris 应该看到错误 "余额不足"
    And 帐户 "ACC001" 的余额应该保持 $1000
    And 帐户 "ACC002" 的余额应该保持 $500
    And 不应该记录转账交易

  # 参数化测试 - 金额验证
  Scenario Outline: 金额验证规则
    Given Chris 已登录
    When Chris 从帐户 "ACC001" 转账 <金额> 到帐户 "ACC002"
    Then 转账应该 <结果>
    And Chris 应该看到 <信息类型> "<信息内容>"
    And 帐户 "ACC001" 的余额应该是 <ACC001余额>

    Examples: 无效金额
      | 金额    | 结果 | 信息类型 | 信息内容           | ACC001余额 |
      | $0      | 失败 | 错误     | 金额必须大于零     | $1000      |
      | $-100   | 失败 | 错误     | 金额不能为负数     | $1000      |

    Examples: 有效金额
      | 金额    | 结果 | 信息类型 | 信息内容    | ACC001余额 |
      | $0.01   | 成功 | 确认     | 转账成功    | $999.99    |
      | $1000   | 成功 | 确认     | 转账成功    | $0         |

  # 帐户验证场景
  Scenario: 转账失败 - 相同帐户
    Given Chris 已登录
    When Chris 从帐户 "ACC001" 转账 $100 到帐户 "ACC001"
    Then 转账应该失败
    And Chris 应该看到错误 "不能转账到相同帐户"
    And 帐户 "ACC001" 的余额应该保持 $1000

  Scenario: 转账失败 - 目标帐户不存在
    Given Chris 已登录
    When Chris 从帐户 "ACC001" 转账 $100 到帐户 "INVALID"
    Then 转账应该失败
    And Chris 应该看到错误 "目标帐户不存在"
    And 帐户 "ACC001" 的余额应该保持 $1000

  # 边界情况
  Scenario: 转账到零余额帐户
    Given Chris 已登录
    When Chris 从帐户 "ACC001" 转账 $100 到帐户 "ACC003"
    Then 转账应该成功
    And 帐户 "ACC003" 的余额应该是 $100
    And Jordan 应该收到转账通知

  Scenario: 转账全部余额
    Given Chris 已登录
    When Chris 从帐户 "ACC001" 转账 $1000 到帐户 "ACC002"
    Then 转账应该成功
    And 帐户 "ACC001" 的余额应该是 $0
    And 帐户 "ACC002" 的余额应该是 $1500
```

### Gherkin 最佳实践

#### 1. 使用业务语言
```gherkin
# ✅ 好
When Chris 转账 $100 到 Lisa

# ❌ 不好
When the TransferService.execute() method is called with amount=100
```

#### 2. 描述行为，不是实现
```gherkin
# ✅ 好
Then Chris 应该看到确认信息

# ❌ 不好
Then the success message should be displayed in the div with id="message"
```

#### 3. 保持场景独立
```gherkin
# ✅ 好 - 每个场景有完整的设置
Scenario: 转账成功
  Given Chris 的帐户有 $1000
  When Chris 转账 $100
  Then 余额应该是 $900

# ❌ 不好 - 依赖前一个场景
Scenario: 再次转账
  When Chris 转账 $100  # 假设帐户已存在
  Then 余额应该是 $800  # 依赖前一个场景的结果
```

#### 4. 使用 Background 减少重复
```gherkin
# ✅ 好
Background:
  Given Chris 已登录
  And Chris 的帐户有 $1000

Scenario: 转账
  When Chris 转账 $100
  ...

Scenario: 查询余额
  When Chris 查询余额
  ...
```

#### 5. 参数化相似场景
```gherkin
# ✅ 好
Scenario Outline: 金额验证
  When 转账 <金额>
  Then 应该 <结果>

  Examples:
    | 金额  | 结果 |
    | $0    | 失败 |
    | $-100 | 失败 |
    | $100  | 成功 |
```

## 阶段 4: 自动化测试 - 将规格转化为可执行测试

### 目的
开发人员和测试人员协作，将 Gherkin 规格转化为自动化的验收测试。

### Step Definitions 实现

#### 4.1 测试框架设置

**文件**: `features/steps/transfer_test.go`

```go
package steps

import (
    "context"
    "testing"

    "github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
    suite := godog.TestSuite{
        ScenarioInitializer: InitializeScenario,
        Options: &godog.Options{
            Format:   "pretty",
            Paths:    []string{"../"},
            TestingT: t,
        },
    }

    if suite.Run() != 0 {
        t.Fatal("non-zero status returned, failed to run feature tests")
    }
}
```

#### 4.2 Step Definitions 实现

**文件**: `features/steps/transfer_steps.go`

```go
package steps

import (
    "context"
    "fmt"

    "github.com/cucumber/godog"
    "github.com/shopspring/decimal"
)

// Test context - 在场景间共享状态
type transferContext struct {
    accounts       map[string]*Account
    currentUser    string
    transferResult *TransferResult
    lastError      error
    service        *TransferService
}

func InitializeScenario(ctx *godog.ScenarioContext) {
    tc := &transferContext{
        accounts: make(map[string]*Account),
        service:  NewTransferService(),
    }

    // Background steps
    ctx.Step(`^存在以下帐户:$`, tc.thereAreAccounts)

    // Given steps (设置前置条件)
    ctx.Step(`^(\w+) 已登录$`, tc.userIsLoggedIn)
    ctx.Step(`^(\w+) 的帐户有 \$(\d+(?:\.\d+)?)$`, tc.userHasBalance)

    // When steps (执行动作)
    ctx.Step(`^(\w+) 从帐户 "([^"]*)" 转账 \$(\d+(?:\.\d+)?) 到帐户 "([^"]*)"$`,
        tc.userTransfersMoney)

    // Then steps (验证结果)
    ctx.Step(`^转账应该成功$`, tc.transferShouldSucceed)
    ctx.Step(`^转账应该失败$`, tc.transferShouldFail)
    ctx.Step(`^帐户 "([^"]*)" 的余额应该是 \$(\d+(?:\.\d+)?)$`,
        tc.accountBalanceShouldBe)
    ctx.Step(`^(\w+) 应该看到确认信息 "([^"]*)"$`,
        tc.userShouldSeeConfirmation)
    ctx.Step(`^(\w+) 应该看到错误 "([^"]*)"$`,
        tc.userShouldSeeError)
    ctx.Step(`^应该记录一笔转账交易$`, tc.transferShouldBeRecorded)
    ctx.Step(`^不应该记录转账交易$`, tc.noTransferShouldBeRecorded)

    // Cleanup after each scenario
    ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
        tc.cleanup()
        return ctx, nil
    })
}

// Background step implementation
func (tc *transferContext) thereAreAccounts(table *godog.Table) error {
    for i, row := range table.Rows[1:] { // Skip header row
        account := &Account{
            ID:      row.Cells[0].Value,
            Owner:   row.Cells[1].Value,
            Balance: decimal.RequireFromString(row.Cells[2].Value),
        }
        tc.accounts[account.ID] = account
    }
    return nil
}

// Given step implementations
func (tc *transferContext) userIsLoggedIn(username string) error {
    tc.currentUser = username
    return nil
}

func (tc *transferContext) userHasBalance(username string, amount float64) error {
    // Find or create account for user
    var accountID string
    for id, acc := range tc.accounts {
        if acc.Owner == username {
            accountID = id
            break
        }
    }

    if accountID == "" {
        accountID = fmt.Sprintf("ACC_%s", username)
        tc.accounts[accountID] = &Account{
            ID:      accountID,
            Owner:   username,
            Balance: decimal.NewFromFloat(amount),
        }
    } else {
        tc.accounts[accountID].Balance = decimal.NewFromFloat(amount)
    }

    return nil
}

// When step implementation
func (tc *transferContext) userTransfersMoney(
    username, fromID string,
    amount float64,
    toID string,
) error {
    ctx := context.Background()
    transferAmount := decimal.NewFromFloat(amount)

    result, err := tc.service.ExecuteTransfer(
        ctx,
        fromID,
        toID,
        transferAmount,
    )

    tc.transferResult = result
    tc.lastError = err

    return nil
}

// Then step implementations
func (tc *transferContext) transferShouldSucceed() error {
    if tc.lastError != nil {
        return fmt.Errorf("expected transfer to succeed but got error: %v", tc.lastError)
    }
    if tc.transferResult == nil || tc.transferResult.Status != TransferStatusCompleted {
        return fmt.Errorf("transfer did not complete successfully")
    }
    return nil
}

func (tc *transferContext) transferShouldFail() error {
    if tc.lastError == nil {
        return fmt.Errorf("expected transfer to fail but it succeeded")
    }
    return nil
}

func (tc *transferContext) accountBalanceShouldBe(accountID string, expectedBalance float64) error {
    account, exists := tc.accounts[accountID]
    if !exists {
        return fmt.Errorf("account %s not found", accountID)
    }

    expected := decimal.NewFromFloat(expectedBalance)
    if !account.Balance.Equal(expected) {
        return fmt.Errorf(
            "expected balance $%.2f but got $%.2f",
            expectedBalance,
            account.Balance.InexactFloat64(),
        )
    }

    return nil
}

func (tc *transferContext) userShouldSeeConfirmation(username, message string) error {
    if tc.transferResult == nil {
        return fmt.Errorf("no transfer result available")
    }
    if tc.transferResult.Message != message {
        return fmt.Errorf(
            "expected message %q but got %q",
            message,
            tc.transferResult.Message,
        )
    }
    return nil
}

func (tc *transferContext) userShouldSeeError(username, expectedError string) error {
    if tc.lastError == nil {
        return fmt.Errorf("expected error but got none")
    }
    if tc.lastError.Error() != expectedError {
        return fmt.Errorf(
            "expected error %q but got %q",
            expectedError,
            tc.lastError.Error(),
        )
    }
    return nil
}

func (tc *transferContext) transferShouldBeRecorded() error {
    // Check that transfer was persisted
    if tc.transferResult == nil || tc.transferResult.ID == "" {
        return fmt.Errorf("transfer was not recorded")
    }
    return nil
}

func (tc *transferContext) noTransferShouldBeRecorded() error {
    // Verify no transfer was saved
    if tc.transferResult != nil && tc.transferResult.ID != "" {
        return fmt.Errorf("transfer should not have been recorded")
    }
    return nil
}

func (tc *transferContext) cleanup() {
    tc.accounts = make(map[string]*Account)
    tc.currentUser = ""
    tc.transferResult = nil
    tc.lastError = nil
}
```

### 实现业务逻辑

#### Domain Layer

**文件**: `internal/domain/transfer.go`

```go
package domain

import (
    "errors"
    "time"

    "github.com/shopspring/decimal"
)

type TransferStatus string

const (
    TransferStatusPending   TransferStatus = "pending"
    TransferStatusCompleted TransferStatus = "completed"
    TransferStatusFailed    TransferStatus = "failed"
)

var (
    ErrInsufficientBalance = errors.New("余额不足")
    ErrInvalidAmount       = errors.New("金额必须大于零")
    ErrNegativeAmount      = errors.New("金额不能为负数")
    ErrSameAccount         = errors.New("不能转账到相同帐户")
    ErrAccountNotFound     = errors.New("帐户不存在")
)

type Transfer struct {
    ID            string
    FromAccountID string
    ToAccountID   string
    Amount        decimal.Decimal
    Status        TransferStatus
    Message       string
    CreatedAt     time.Time
    CompletedAt   *time.Time
}

// Validate 验证转账请求
func (t *Transfer) Validate() error {
    // 验证金额
    if t.Amount.IsNegative() {
        return ErrNegativeAmount
    }
    if t.Amount.IsZero() {
        return ErrInvalidAmount
    }

    // 验证帐户
    if t.FromAccountID == t.ToAccountID {
        return ErrSameAccount
    }
    if t.FromAccountID == "" || t.ToAccountID == "" {
        return ErrAccountNotFound
    }

    return nil
}

type Account struct {
    ID      string
    Owner   string
    Balance decimal.Decimal
}

// CanTransfer 检查是否可以转出指定金额
func (a *Account) CanTransfer(amount decimal.Decimal) bool {
    return a.Balance.GreaterThanOrEqual(amount)
}

// Debit 从帐户扣款
func (a *Account) Debit(amount decimal.Decimal) error {
    if !a.CanTransfer(amount) {
        return ErrInsufficientBalance
    }
    a.Balance = a.Balance.Sub(amount)
    return nil
}

// Credit 向帐户存款
func (a *Account) Credit(amount decimal.Decimal) {
    a.Balance = a.Balance.Add(amount)
}
```

#### Service Layer

**文件**: `internal/service/transfer_service.go`

```go
package service

import (
    "context"
    "fmt"
    "time"

    "github.com/google/uuid"
    "github.com/shopspring/decimal"

    "your-project/internal/domain"
)

type TransferService struct {
    accountRepo  AccountRepository
    transferRepo TransferRepository
}

func NewTransferService(
    accountRepo AccountRepository,
    transferRepo TransferRepository,
) *TransferService {
    return &TransferService{
        accountRepo:  accountRepo,
        transferRepo: transferRepo,
    }
}

type TransferResult struct {
    ID      string
    Status  domain.TransferStatus
    Message string
}

// ExecuteTransfer 执行转账操作
func (s *TransferService) ExecuteTransfer(
    ctx context.Context,
    fromAccountID, toAccountID string,
    amount decimal.Decimal,
) (*TransferResult, error) {
    // 创建转账对象
    transfer := &domain.Transfer{
        ID:            uuid.New().String(),
        FromAccountID: fromAccountID,
        ToAccountID:   toAccountID,
        Amount:        amount,
        Status:        domain.TransferStatusPending,
        CreatedAt:     time.Now(),
    }

    // 验证转账
    if err := transfer.Validate(); err != nil {
        return nil, err
    }

    // 加载帐户
    fromAccount, err := s.accountRepo.FindByID(ctx, fromAccountID)
    if err != nil {
        return nil, domain.ErrAccountNotFound
    }

    toAccount, err := s.accountRepo.FindByID(ctx, toAccountID)
    if err != nil {
        return nil, fmt.Errorf("目标%w", domain.ErrAccountNotFound)
    }

    // 执行转账（在事务中）
    if err := s.executeInTransaction(ctx, transfer, fromAccount, toAccount); err != nil {
        transfer.Status = domain.TransferStatusFailed
        transfer.Message = err.Error()
        return &TransferResult{
            ID:      transfer.ID,
            Status:  transfer.Status,
            Message: transfer.Message,
        }, err
    }

    // 转账成功
    now := time.Now()
    transfer.Status = domain.TransferStatusCompleted
    transfer.Message = "转账成功"
    transfer.CompletedAt = &now

    return &TransferResult{
        ID:      transfer.ID,
        Status:  transfer.Status,
        Message: transfer.Message,
    }, nil
}

func (s *TransferService) executeInTransaction(
    ctx context.Context,
    transfer *domain.Transfer,
    from, to *domain.Account,
) error {
    // 开始事务
    tx, err := s.accountRepo.BeginTx(ctx)
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }
    defer tx.Rollback()

    // 扣款
    if err := from.Debit(transfer.Amount); err != nil {
        return err
    }

    // 入款
    to.Credit(transfer.Amount)

    // 更新帐户
    if err := s.accountRepo.UpdateInTx(ctx, tx, from); err != nil {
        return fmt.Errorf("failed to update source account: %w", err)
    }

    if err := s.accountRepo.UpdateInTx(ctx, tx, to); err != nil {
        return fmt.Errorf("failed to update destination account: %w", err)
    }

    // 保存转账记录
    if err := s.transferRepo.SaveInTx(ctx, tx, transfer); err != nil {
        return fmt.Errorf("failed to save transfer: %w", err)
    }

    // 提交事务
    if err := tx.Commit(); err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}
```

### 运行自动化测试

```bash
# 运行所有 BDD 场景
go test -v ./features/...

# 运行特定 feature
go test -v ./features/... -godog.tags="@transfer"

# 生成测试报告
go test -v ./features/... -godog.format=cucumber:report.json
```

### 测试输出示例

```
Feature: 帐户转账

  Scenario: 成功转账 - 余额充足                      # features/money_transfer.feature:12
    Given Chris 已登录                               # transfer_steps.go:45
    When Chris 从帐户 "ACC001" 转账 $300 到帐户 "ACC002"  # transfer_steps.go:52
    Then 转账应该成功                                # transfer_steps.go:67
    And 帐户 "ACC001" 的余额应该是 $700                # transfer_steps.go:78
    And 帐户 "ACC002" 的余额应该是 $800                # transfer_steps.go:78
    And Chris 应该看到确认信息 "转账成功"                # transfer_steps.go:93

  Scenario: 转账失败 - 余额不足                      # features/money_transfer.feature:21
    Given Chris 已登录                               # transfer_steps.go:45
    When Chris 从帐户 "ACC001" 转账 $1500 到帐户 "ACC002" # transfer_steps.go:52
    Then 转账应该失败                                # transfer_steps.go:74
    And Chris 应该看到错误 "余额不足"                   # transfer_steps.go:102
    And 帐户 "ACC001" 的余额应该保持 $1000              # transfer_steps.go:78

6 scenarios (6 passed)
28 steps (28 passed)
145.32ms
```

## 阶段 5: 活文档 - 测试即文档

### 目的
自动化的 BDD 测试不仅是验收测试，也是系统行为的精确且最新的文档。

### 活文档的价值

#### 1. 可执行的规格说明
- ✅ 规格即测试，测试即规格
- ✅ 文档与代码保持同步
- ✅ 不会过时的文档

#### 2. 多重用途

```
BDD 自动化测试的价值

1. 验收测试
   ↓
   验证功能符合业务需求

2. 回归测试
   ↓
   确保新变更不破坏现有功能

3. 产品文档
   ↓
   展示系统如何运作的实际范例

4. 手动测试基础
   ↓
   指导探索性测试的方向
```

#### 3. 团队沟通工具

**对业务人员**:
- 理解系统当前实现的功能
- 验证业务规则是否正确实现
- 提供新需求的讨论基础

**对开发人员**:
- 理解业务意图
- 重构时的安全网
- 新成员的学习资源

**对测试人员**:
- 自动化测试的基础
- 探索性测试的起点
- 测试覆盖的可视化

### 生成活文档报告

#### 使用 Cucumber 报告

```bash
# 生成 HTML 报告
go test ./features/... -godog.format=cucumber:cucumber.json

# 使用 cucumber-html-reporter 生成网页
npx cucumber-html-reporter \
    --input cucumber.json \
    --output report.html \
    --theme bootstrap
```

#### 报告内容示例

```html
<!DOCTYPE html>
<html>
<head>
    <title>BDD Test Report - 帐户转账功能</title>
</head>
<body>
    <h1>Feature: 帐户转账</h1>
    <p class="user-story">
        As a 银行客户<br>
        I want to 在我的帐户之间转账<br>
        So that 我可以灵活管理我的资金
    </p>

    <h2>Scenarios (6 total, 6 passed)</h2>

    <div class="scenario passed">
        <h3>✓ 成功转账 - 余额充足</h3>
        <div class="steps">
            <p>✓ Given Chris 已登录</p>
            <p>✓ When Chris 从帐户 "ACC001" 转账 $300 到帐户 "ACC002"</p>
            <p>✓ Then 转账应该成功</p>
            <p>✓ And 帐户 "ACC001" 的余额应该是 $700</p>
            <p>✓ And 帐户 "ACC002" 的余额应该是 $800</p>
        </div>
        <p class="duration">执行时间: 24ms</p>
    </div>

    <!-- 更多场景... -->

    <h2>Statistics</h2>
    <table>
        <tr><td>Total Scenarios</td><td>6</td></tr>
        <tr><td>Passed</td><td>6 (100%)</td></tr>
        <tr><td>Failed</td><td>0</td></tr>
        <tr><td>Total Steps</td><td>28</td></tr>
        <tr><td>Execution Time</td><td>145ms</td></tr>
    </table>
</body>
</html>
```

### 将文档整合到开发流程

#### 1. CI/CD 集成

```yaml
# .github/workflows/bdd-tests.yml
name: BDD Tests

on: [push, pull_request]

jobs:
  bdd-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.21

      - name: Run BDD Tests
        run: |
          go test -v ./features/... \
            -godog.format=cucumber:cucumber.json \
            -godog.format=pretty

      - name: Generate Report
        run: |
          npm install -g cucumber-html-reporter
          cucumber-html-reporter \
            --input cucumber.json \
            --output report.html

      - name: Upload Report
        uses: actions/upload-artifact@v2
        with:
          name: bdd-test-report
          path: report.html
```

#### 2. 文档网站

将 BDD 测试报告发布到团队可访问的位置：
- 内部文档网站
- Confluence 或类似平台
- GitHub Pages

#### 3. 定期审查

在 Sprint Review 或类似会议中：
- 展示新的 BDD 场景
- 讨论场景覆盖是否完整
- 识别需要补充的场景

### 手动和探索性测试

BDD 自动化测试作为手动测试的基础：

#### 手动测试检查清单

基于 `money_transfer.feature`:

```markdown
# 手动测试清单 - 帐户转账

## 功能测试
- [ ] 验证所有自动化场景在 UI 中正常工作
- [ ] 测试转账确认邮件是否发送
- [ ] 验证转账历史记录显示正确
- [ ] 检查并发转账的处理

## 探索性测试
- [ ] 尝试在转账过程中取消操作
- [ ] 测试网络中断时的行为
- [ ] 验证帐户冻结状态下的转账
- [ ] 测试高并发场景下的性能

## 非功能测试
- [ ] 页面加载时间 < 2秒
- [ ] 转账确认响应时间 < 1秒
- [ ] 支持 100+ 并发转账
- [ ] 验证移动设备上的用户体验

## 安全测试
- [ ] 测试未授权访问
- [ ] 验证敏感信息是否加密
- [ ] 测试 SQL 注入防护
- [ ] 验证审计日志完整性
```

#### 探索性测试焦点

BDD 场景覆盖了已知的业务规则，探索性测试关注：
- **未预期的交互**: 用户可能采取的非标准操作
- **边界和极限**: 系统在极端条件下的表现
- **用户体验**: 实际使用时的感受和便利性
- **集成问题**: 与其他系统交互时的问题

## 完整工作流程示例

### 案例：添加转账限额功能

#### 阶段 1: 业务对话

**场景**: 监管要求单笔转账金额不能超过 $10,000

**Three Amigos 会议**:

- **业务**: 为了合规，我们需要限制单笔转账金额
- **开发**: 这需要在验证层添加新的检查
- **测试**: 如果用户需要转账更多金额怎么办？会有豁免机制吗？

**结论**:
- 标准用户限额 $10,000
- VIP 用户限额 $50,000
- 超限需要额外审批

#### 阶段 2: 范例探索

```
规则: 转账限额检查

范例 1: 标准用户 - 限额内转账
Given Chris 是标准用户
And Chris 的帐户有 $20,000
When Chris 转账 $9,000
Then 转账应该成功

范例 2: 标准用户 - 超过限额
Given Chris 是标准用户
And Chris 的帐户有 $20,000
When Chris 转账 $15,000
Then 转账应该被拒绝
And 错误信息应该是 "超过单笔转账限额 $10,000"

范例 3: VIP 用户 - 更高限额
Given Lisa 是 VIP 用户
And Lisa 的帐户有 $100,000
When Lisa 转账 $45,000
Then 转账应该成功

范例 4: VIP 用户 - 超过 VIP 限额
Given Lisa 是 VIP 用户
And Lisa 的帐户有 $100,000
When Lisa 转账 $60,000
Then 转账应该被拒绝
And 错误信息应该是 "超过单笔转账限额 $50,000"
```

#### 阶段 3: Gherkin 规格

```gherkin
Feature: 转账限额控制
  为了符合监管要求
  系统必须限制单笔转账金额

  Background:
    Given 存在以下用户:
      | 用户   | 类型   | 帐户余额   | 转账限额  |
      | Chris  | 标准   | 20000     | 10000    |
      | Lisa   | VIP    | 100000    | 50000    |

  Scenario Outline: 转账限额验证
    Given <用户> 已登录
    When <用户> 尝试转账 $<金额>
    Then 转账应该 <结果>
    And <用户> 应该看到 <消息>

    Examples: 标准用户
      | 用户  | 金额  | 结果 | 消息                           |
      | Chris | 9000  | 成功 | "转账成功"                      |
      | Chris | 10000 | 成功 | "转账成功"                      |
      | Chris | 10001 | 失败 | "超过单笔转账限额 $10,000"      |
      | Chris | 15000 | 失败 | "超过单笔转账限额 $10,000"      |

    Examples: VIP 用户
      | 用户 | 金额  | 结果 | 消息                           |
      | Lisa | 45000 | 成功 | "转账成功"                      |
      | Lisa | 50000 | 成功 | "转账成功"                      |
      | Lisa | 50001 | 失败 | "超过单笔转账限额 $50,000"      |
      | Lisa | 60000 | 失败 | "超过单笔转账限额 $50,000"      |
```

#### 阶段 4 & 5: TDD 实现和自动化

**RED - 写失败的测试**:
```go
func TestTransfer_ValidateLimit(t *testing.T) {
    transfer := &Transfer{
        Amount: decimal.NewFromInt(15000),
        User:   &User{Type: UserTypeStandard},
    }

    err := transfer.ValidateLimit()
    assert.Error(t, err)
    assert.Equal(t, ErrExceedsTransferLimit, err)
}
```

**GREEN - 实现代码**:
```go
func (t *Transfer) ValidateLimit() error {
    limit := t.User.GetTransferLimit()
    if t.Amount.GreaterThan(limit) {
        return fmt.Errorf("超过单笔转账限额 $%s", limit.String())
    }
    return nil
}
```

**REFACTOR - 优化**:
```go
type TransferLimitPolicy interface {
    GetLimit(user *User) decimal.Decimal
}

func (t *Transfer) ValidateWithPolicy(policy TransferLimitPolicy) error {
    limit := policy.GetLimit(t.User)
    if t.Amount.GreaterThan(limit) {
        return &TransferLimitError{
            Limit:  limit,
            Amount: t.Amount,
        }
    }
    return nil
}
```

## 最佳实践总结

### DO ✅

1. **协作优先**
   - 定期进行 Three Amigos 会议
   - 鼓励所有角色参与讨论
   - 使用白板和便利贴进行 Example Mapping

2. **具体化**
   - 使用真实的名字、数据和场景
   - 提供足够的上下文
   - 写明确的预期结果

3. **业务语言**
   - 使用领域词汇，不用技术术语
   - 让非技术人员能理解
   - 关注"做什么"而不是"怎么做"

4. **保持简洁**
   - 每个场景测试一个业务规则
   - 场景独立，不相互依赖
   - 使用 Background 减少重复

5. **持续维护**
   - 及时更新场景以反映需求变化
   - 删除过时的场景
   - 重构重复的步骤定义

### DON'T ❌

1. **跳过协作**
   - 不要开发人员单独写所有场景
   - 不要等到实现完成才写场景
   - 不要忽略测试人员的输入

2. **过于抽象**
   - 避免使用占位符和变量名
   - 不要写"应该正确工作"这样的模糊断言
   - 避免依赖隐式状态

3. **技术细节**
   - 不要在场景中提及类名、方法名
   - 避免描述 UI 交互细节
   - 不要暴露数据库或 API 实现

4. **耦合场景**
   - 避免场景间的依赖
   - 不要假设执行顺序
   - 每个场景应能独立运行

5. **忽略维护**
   - 不要让失败的场景一直失败
   - 不要保留不再相关的场景
   - 不要忽略代码重复

## 总结

BDD 工作流程的核心是**协作**和**具体化**：

```
业务对话 → 理解需求
    ↓
具体范例 → 共同语言
    ↓
Gherkin → 可执行规格
    ↓
自动化测试 → 持续验证
    ↓
活文档 → 知识共享
```

这个流程确保：
- ✅ 构建正确的东西（需求清晰）
- ✅ 正确地构建（测试驱动）
- ✅ 持续验证（自动化测试）
- ✅ 知识保留（活文档）

### 关键要点

1. **BDD 不只是测试工具** - 它是一种协作实践
2. **具体范例是关键** - 抽象需求难以验证
3. **Gherkin 是沟通工具** - 不仅仅是测试脚本
4. **自动化测试有多重价值** - 验收、回归、文档
5. **持续演进** - 随需求变化更新场景

### 下一步

继续学习：
- **Chapter 2**: 编写高质量 Gherkin 场景的技巧
- **Chapter 3**: 高级 Step Definition 模式
- **Chapter 4**: BDD 与领域驱动设计 (DDD) 的结合
- **Chapter 5**: 端到端测试策略

## 参考资源

### 书籍
- "BDD in Action" - John Ferguson Smart
- "Specification by Example" - Gojko Adzic
- "The Cucumber Book" - Matt Wynne & Aslak Hellesøy

### 在线资源
- [Cucumber BDD Guide](https://cucumber.io/docs/bdd/)
- [Example Mapping](https://cucumber.io/blog/bdd/example-mapping-introduction/)
- [Gherkin Reference](https://cucumber.io/docs/gherkin/reference/)
- [Godog Framework](https://github.com/cucumber/godog)

### 工具
- **Godog**: Go 的 BDD 框架
- **Cucumber**: 多语言 BDD 工具
- **SpecFlow**: .NET 的 BDD 框架
- **Behat**: PHP 的 BDD 框架

---

**Ready to practice?** 选择一个简单的功能，召集你的 Three Amigos，开始你的 BDD 之旅！
