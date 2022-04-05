package golangweb

// secara default golang akan aman dari xss, atau cross site scripting
// kalau tidak hati" kita bisa terkena cross site scripting ini dan bahanya dia bisa mengambil data cookie dari user
// jika iningin menon aktifkan fitue xss di golang cukup gunakan perintah template.HTML / template.CSS / template.JS
// jika benar benar butuh kamu bisa disable fitur ini, sperti ketika membuat blog artikel dll, yang harus merender file html/css/js
// jangan merender bulat" data yang sensitif dari user
