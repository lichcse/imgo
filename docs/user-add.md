# API: Thêm User Mới

## Tổng quan

| Thuộc tính | Giá trị |
|------------|---------|
| **Method** | POST |
| **Endpoint** | `/identity/v1/user` |
| **Mô tả** | Tạo mới một user trong hệ thống |
| **Tags** | identity |

---

## Mục đích (Dành cho Business/Non-tech)

API này cho phép hệ thống **thêm mới một người dùng** vào cơ sở dữ liệu. Khi người dùng đăng ký tài khoản hoặc admin thêm user mới, hệ thống sẽ gọi API này để lưu thông tin user vào database.

**Ví dụ thực tế:**
- Người dùng điền form đăng ký trên website → API này lưu thông tin vào DB
- Admin thêm nhân viên mới vào hệ thống quản lý

---

## Request Parameters

### Headers

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| Content-Type | string | Yes | `application/json` |
| lang | string | No | Ngôn ngữ trả về: `en` hoặc `vi` |

### Body

```json
{
  "full_name": "Lich Truong",
  "username": "lichtv",
  "email": "example@imgo.com",
  "password": "W3^&(80)&&^x"
}
```

### Parameters Detail

| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| full_name | string | Yes | 3-50 ký tự | Họ và tên đầy đủ |
| username | string | Yes | Chỉ chứa a-z, A-Z, 0-9, dấu gạch ngang | Tên đăng nhập |
| email | string | Yes | Định dạng email hợp lệ | Địa chỉ email |
| password | string | Yes | Tối thiểu 8 ký tự | Mật khẩu |

---

## Response

### Success (200)

```json
{
  "code": 200,
  "data": {
    "id": 1,
    "full_name": "Lich Truong",
    "username": "lichtv",
    "email": "example@imgo.com",
    "created_at": "1991-02-13 10:10:10",
    "modified_at": "2020-07-15 10:10:10",
    "status": 1
  },
  "message": "success"
}
```

### Error

| Code | Message | Description |
|------|---------|-------------|
| 400 | invalid_email | Email không đúng định dạng |
| 400 | invalid_username | Username không hợp lệ |
| 400 | password_too_short | Mật khẩu quá ngắn |
| 400 | full_name_required | Họ tên là bắt buộc |

---

## Sequence Diagram

### Dành cho Developer (Technical)

```mermaid
sequenceDiagram
    participant Client as Client (Frontend/Mobile)
    participant API as API Gateway
    participant Handler as UserHandler
    participant Service as UserService
    participant Repo as UserRepository
    participant DB as MySQL Database

    Client->>API: POST /identity/v1/user<br/>(JSON: full_name, username, email, password)
    API->>Handler: Gọi handler.Add()
    Handler->>Handler: Validate request<br/>(email, username, password, full_name)
    
    alt Validation failed
        Handler-->>Client: Return error 400
    else Validation passed
        Handler->>Service: Gọi service.Add()
        Service->>Repo: Gọi repo.Create()
        Repo->>DB: INSERT INTO users (...)
        DB-->>Repo: Return inserted user
        Repo-->>Service: Return user object
        Service-->>Handler: Return UserDetailResponse
        Handler-->>Client: Return 200 + user data
    end
```

### Dành cho Business/Non-tech

```mermaid
flowchart TD
    A[Người dùng điền<br/>form đăng ký] --> B[Hệ thống gửi<br/>API thêm user]
    B --> C{Hướng dẫn<br/>hợp lệ?}
    C -->|Không| D[Trả về lỗi<br/>cho người dùng]
    C -->|Có| E[Lưu vào<br/>CSDL]
    E --> F[Trả về thông tin<br/>user đã tạo]
    F --> G[Hoàn tất<br/>đăng ký]
    
    style A fill:#e1f5fe
    style E fill:#c8e6c9
    style G fill:#c8e6c9
    style D fill:#ffcdd2
```

---

## Ví dụ sử dụng (cURL)

```bash
# Thêm user mới
curl -X POST http://localhost:8080/identity/v1/user \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Lich Truong",
    "username": "lichtv",
    "email": "example@imgo.com",
    "password": "W3^&(80)&&^x"
  }'
```

---

## Lưu ý

1. **Bảo mật**: Password được hash trước khi lưu vào database
2. **Validation**: Tất cả các trường đều được validate phía server
3. **Ngôn ngữ**: Có thể chọn ngôn ngữ trả về (en/vi) qua query param `lang`