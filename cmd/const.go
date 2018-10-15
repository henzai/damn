package cmd

// Dump用
const shCmdWithCompress = "docker exec -i %v pg_dump --no-owner -U %v %v | gzip -9 > %v.sql.gz"
const shCmd = "docker exec -i %v pg_dump --no-owner -U %v %v > %v.sql"

// Restore用
