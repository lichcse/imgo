# API: Cập Nhật User

## Tổng quan

| Thuộc tính | Giá trị |
|------------|---------|
| **Method** | PUT |
| **Endpoint** | `/identity/v1/user/:id` |
| **Mô tả** | Cập nhật thông tin của một user theo ID |
| **Tags** | identity |

---

## Mục đích (Dành cho Business/Non-tech)

API này dùng để **cập nhật thông tin của một người dùng** đã có trong hệ thống. Khi user muốn thay đổi thông tin cá nhân (họ tên, email) hoặc admin cần cập nhật thông tin nhân viên, hệ thống sẽ gọi API này.

**Ví dụ thực tế:**
- User thay đổi email cá nhân
- Admin cập nhật thông tin nhân viên
- User thay đổi username

---

## Request Parameters

### Headers

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| Content-Type | string | Yes | `application/json` |
| lang | string | No | Ngôn ngữ trả về: `en` hoặc `vi` |

### Path Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | uint64 | Yes | ID của user cần cập nhật |

### Body

```json
{
  "full_name": "Lich Truong Updated",
  "username": "lichtv_new",
  "email": "newemail@imgo.com",
  "password": "newpassword123"
}
```

### Parameters Detail

| Field | Type | Required | Constraints | Description |
|-------|------|----------|-------------|-------------|
| full_name | string | No | 3-50 ký tự | Họ và tên đầy đủ |
| username | string | No | Chỉ chứa a-z, A-Z, 0-9, dấu gạch ngang | Tên đăng nhập |
| email | string | No | Định dạng email hợp lệ | Địa chỉ email |
| password | string | No | Tối thiểu 8 ký tự | Mật khẩu mới |

---

## Response

### Success (200)

```json
{
  "code": 200,
  "data": {
    "id": 1,
    "full_name": "Lich Truong Updated",
    "username": "lichtv_new",
    "email": "newemail@imgo.com",
    "created_at": "1991-02-13 10:10:10",
    "modified_at": "2024-01-15 10:10:10",
    "status": 1
  },
  "message": "success"
}
```

### Error

| Code | Message | Description |
|------|---------|-------------|
| 400 | not_allow | ID không hợp lệ |
| 400 | invalid_email | Email không đúng định dạng |
| 400 | user_not_found | Không tìm thấy user |
| 400 | email_already_exists | Email đã được sử dụng bởi user khác |

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

    Client->>API: PUT /identity/v1/user/1<br/>(JSON: full_name, username, email, password)
    API->>Handler: Gọi handler.Update(id=1)
    Handler->>Handler: Validate id và request body
    
    alt Validation failed
        Handler-->>Client: Return error 400
    else Validation passed
        Handler->>Service: Gọi service.Update(id, request)
        Service->>Repo: Gọi repo.GetByID(id)
        alt User not found
            Repo-->>Service: Return nil
            Service-->>Handler: Return error 404
            Handler-->>Client: Return 404 (user_not_found)
        else User found
            Service->>Repo: Gọi repo.Update()
            Repo->>DB: UPDATE users SET ... WHERE id = ?
            DB-->>Repo: Return updated user
            Repo-->>Service: Return user object
            Service-->>Handler: Return UserDetailResponse
            Handler-->>Client: Return 200 + updated user data
        end
    end
```

### Dành cho Business/Non-tech

```mermaid
flowchart TD
    A[Người dùng/Màn hình<br/>yêu cầu cập nhật] --> B[Hệ thống gọi API<br/>cập nhật user]
    B --> C{Dữ liệu<br/>hợp lệ?}
    C -->|Không| D[Hiển thị lỗi<br/>validation]
    C -->|Có| E[Tìm user<br/>trong CSDL]
    E --> F{Tìm thấy<br/>user?}
    F -->|Không| G[Hiển thị thông báo<br/>không tìm thấy]
    F -->|Có| H[Cập nhật<br/>thông tin]
    H --> I[Trả về thông tin<br/>đã cập nhật]
    
    style A fill:#e1f5fe
    style H fill:#c8e6c9
    style I fill:#c8e6c9
    style D fill:#ffcdd2
    style G fill:#ffcdd2
```

---

## Ví dụ sử dụng (cURL)

```bash
# Cập nhật user có ID = 1
curl -X PUT http://localhost:8080/identity/v1/user/1 \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Lich Truong Updated",
    "username": "lichtv_new",
    "email": "newemail@imgo.com"
  }'
```

---

## Lưu ý

1. **Cập nhật một phần**: Chỉ cần gửi các trường cần thay đổi, không bắt buộc gửi tất cả
2. **Validation**: Các trường gửi lên vẫn được validate đầy đủ
3. **ModifiedAt**: Trường modified_at sẽ được cập nhật tự động