# API: Xóa User

## Tổng quan

| Thuộc tính | Giá trị |
|------------|---------|
| **Method** | DELETE |
| **Endpoint** | `/identity/v1/user/:id` |
| **Mô tả** | Xóa một user khỏi hệ thống theo ID |
| **Tags** | identity |

---

## Mục đích (Dành cho Business/Non-tech)

API này dùng để **xóa vĩnh viễn một người dùng** khỏi hệ thống. Khi user yêu cầu xóa tài khoản hoặc admin cần xóa user không còn hoạt động, hệ thống sẽ gọi API này.

**Ví dụ thực tế:**
- User yêu cầu xóa tài khoản (xóa vĩnh viễn)
- Admin xóa user vi phạm quy định
- Dọn dẹp user test trong môi trường dev

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
| id | uint64 | Yes | ID của user cần xóa |

---

## Response

### Success (200)

```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### Error

| Code | Message | Description |
|------|---------|-------------|
| 400 | not_allow | ID không hợp lệ (0 hoặc không phải số) |
| 404 | user_not_found | Không tìm thấy user với ID này |

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

    Client->>API: DELETE /identity/v1/user/1
    API->>Handler: Gọi handler.Delete(id=1)
    Handler->>Handler: Validate id (phải là số > 0)
    
    alt ID không hợp lệ
        Handler-->>Client: Return error 400 (not_allow)
    else ID hợp lệ
        Handler->>Service: Gọi service.Delete(id)
        Service->>Repo: Gọi repo.GetByID(id)
        alt User không tồn tại
            Repo-->>Service: Return nil
            Service-->>Handler: Return error
            Handler-->>Client: Return 404 (user_not_found)
        else User tồn tại
            Service->>Repo: Gọi repo.Delete()
            Repo->>DB: DELETE FROM users WHERE id = ?
            DB-->>Repo: Return success
            Repo-->>Service: Return success
            Service-->>Handler: Return success
            Handler-->>Client: Return 200 (success)
        end
    end
```

### Dành cho Business/Non-tech

```mermaid
flowchart TD
    A[Admin/Người dùng<br/>yêu cầu xóa user] --> B[Hệ thống gọi API<br/>xóa user]
    B --> C{ID user<br/>hợp lệ?}
    C -->|Không| D[Hiển thị lỗi<br/>không hợp lệ]
    C -->|Có| E[Tìm kiếm<br/>trong CSDL]
    E --> F{Tìm thấy<br/>user?}
    F -->|Không| G[Hiển thị thông báo<br/>không tìm thấy]
    F -->|Có| H[Xóa user<br/>khỏi CSDL]
    H --> I[Xác nhận đã xóa<br/>thành công]
    
    style A fill:#e1f5fe
    style H fill:#ff9800
    style I fill:#c8e6c9
    style D fill:#ffcdd2
    style G fill:#ffcdd2
```

---

## Ví dụ sử dụng (cURL)

```bash
# Xóa user có ID = 1
curl -X DELETE http://localhost:8080/identity/v1/user/1
```

---

## Lưu ý

1. **Xóa vĩnh viễn**: Đây là xóa cứng (hard delete), dữ liệu sẽ bị xóa vĩnh viễn khỏi database
2. **ID bắt buộc**: Phải cung cấp ID hợp lệ (số nguyên dương)
3. **Không tìm thấy**: Trả về 404 nếu user không tồn tại
4. **Cân nhắc**: Nên cân nhắc sử dụng soft delete (đánh dấu status = 0) thay vì xóa vĩnh viễn để phục hồi nếu cần