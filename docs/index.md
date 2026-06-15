# API Documentation Index

## Tổng quan

Tài liệu này cung cấp thông tin chi tiết về tất cả các API trong hệ thống **imgo**.

---

## Danh sách API

### Identity Module (`/identity/v1`)

| # | Method | Endpoint | Mô tả | Tài liệu |
|---|--------|----------|-------|----------|
| 1 | POST | `/identity/v1/user` | Thêm mới một user | [Xem chi tiết](./user-add.md) |
| 2 | GET | `/identity/v1/user/:id` | Lấy thông tin chi tiết của user | [Xem chi tiết](./user-detail.md) |
| 3 | PUT | `/identity/v1/user/:id` | Cập nhật thông tin user | [Xem chi tiết](./user-update.md) |
| 4 | DELETE | `/identity/v1/user/:id` | Xóa user khỏi hệ thống | [Xem chi tiết](./user-delete.md) |

---

## Cấu trúc tài liệu API

Mỗi file tài liệu API bao gồm:

1. **Tổng quan** - Thông tin cơ bản về API (method, endpoint, mô tả)
2. **Mục đích (Non-tech)** - Giải thích dễ hiểu cho business/non-technical
3. **Request Parameters** - Chi tiết các tham số yêu cầu
   - Headers
   - Path Parameters
   - Query Parameters
   - Body (nếu có)
4. **Response** - Các trường hợp response thành công và lỗi
5. **Sequence Diagram**
   - Technical: Dành cho developer
   - Non-tech: Dành cho business/stakeholder
6. **Ví dụ sử dụng** - Curl command để test API
7. **Lưu ý** - Các lưu ý quan trọng khi sử dụng

---

## Base URL

```
Development: http://localhost:8080
Production:  https://api.imgo.com (ví dụ)
```

---

## Authentication

Hiện tại các API trong module Identity chưa yêu cầu authentication. Các API trong tương lai có thể sẽ yêu cầu Bearer Token.

---

## Response Format

Tất cả API trả về JSON với cấu trúc chung:

```json
{
  "code": 200,
  "data": { ... },
  "message": "success"
}
```

| Trường | Type | Mô tả |
|--------|------|-------|
| code | int | Mã HTTP response |
| data | object | Dữ liệu trả về (có thể là null) |
| message | string | Thông báo trạng thái |

---

## Ngôn ngữ (i18n)

Các API hỗ trợ đa ngôn ngữ qua header `lang`:

| Giá trị | Ngôn ngữ |
|---------|----------|
| en | English (mặc định) |
| vi | Tiếng Việt |

**Ví dụ:**
```bash
curl -X GET http://localhost:8080/identity/v1/user/1 \
  -H "lang: vi"
```

---

## Liên hệ & Hỗ trợ

- **Email:** support@imgo.com
- **GitHub:** https://github.com/lichcse/imgo

---

## Changelog

| Ngày | Mô tả |
|------|-------|
| 2024-01-15 | Thêm tài liệu API user (Add, Detail, Update, Delete) |