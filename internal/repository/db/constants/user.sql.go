package constants

const GetAllUser = `SELECT * FROM users`
const GetUserById = `SELECT * FROM users WHERE id = $1`
const InsertOneUser = `INSERT INTO users("id","username","email","password","address","age") 
VALUES ($1,$2,$3,$4,$5,$6)`
