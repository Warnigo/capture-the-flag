FROM golang:1.17 AS build

WORKDIR /build

COPY main.go .

RUN go build -o main main.go

FROM debian:latest

WORKDIR /app

COPY --from=build /build/main .

RUN mkdir -p /app/c/Long\ Journey/Of\ Unknown\ Perils \
    && mkdir -p /app/c/.hidden \
    && mkdir -p /app/c/RumbleInTheJungle \
    && echo "OS{what-you_C:Lis_what_you_get}" > /app/c/flag.txt \
    && echo "OS{a-flag-from-beyond}" > "/app/c/Long Journey/Of Unknown Perils/flag.txt" \
    && echo "OS{now-you-see-me-now-you-dont}" > /app/c/.hidden/flag.txt \
    && echo "FAKE_FLAG{not_this_one}" > /app/c/RumbleInTheJungle/flag1.txt \
    && echo "FAKE_FLAG{try_again}" > /app/c/RumbleInTheJungle/flag2.txt \
    && echo "OS{last_man_standing_597263}" > /app/c/RumbleInTheJungle/true_flag.txt \
    && touch -t 202301010000 /app/c/RumbleInTheJungle/flag1.txt \
    && touch -t 202301010001 /app/c/RumbleInTheJungle/flag2.txt \
    && touch -t 202301010002 /app/c/RumbleInTheJungle/true_flag.txt

CMD ["./main"]
