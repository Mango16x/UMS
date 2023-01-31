wrk.method = "POST"
wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

function request()
--     local n = math.random(1, 10000000) -- no cache
    local n = math.random(1, 200) -- cache
    local body = "username=" .. n .. "&password=" .. n
    wrk.body = body
    return wrk.format(wrk.method, "/login", wrk.headers, wrk.body)
end

--  wrk -t20 -c2000 -d20s -s ./scripts/wrk.lua http://localhost:8080/login