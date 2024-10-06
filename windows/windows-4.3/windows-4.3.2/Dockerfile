# Create a Windows-like environment in Docker
FROM golang:1.17

WORKDIR /app

COPY main.go .

RUN go build -o main main.go

RUN mkdir -p /C/PerfLogs \
    && mkdir -p /C/Program\ Files \
    && mkdir -p /C/Program\ Files\ \(x86\) \
    && mkdir -p /C/ProgramData \
    && mkdir -p /C/Users \
    && mkdir -p /C/Windows \
    && mkdir -p /C/Windows/Temp \
    && mkdir -p "/C/Long Journey/Of Unknown Perils" \
    && mkdir -p /C/.hidden \
    && mkdir -p /C/RumbleInTheJungle
    # && echo "OS{what-you_C:Lis_what-you-get}" > /C/flag.txt \
    # && echo "OS{a-flag-from-beyond}" > "/C/Long Journey/Of Unknown Perils/flag.txt" \
    # && echo "OS{now-you-see-me-now-you-dont}" > /C/.hidden/flag.txt \
    # && echo "OS{last_man_standing_597263}" > /C/RumbleInTheJungle/true_flag.txt

CMD ["./main"]
