export interface User {
  ID: number
  username: string
  firstname: string
  lastname: string
  email: string
}

export interface Inbox {
  date: string
  body: string
  from: string
  message_id: string
  subject: string
  to: string
  _timestamp: number
}
