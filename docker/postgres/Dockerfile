FROM postgres:16

# ตั้ง timezone ให้ตรงไทย (optional แต่แนะนำ)
ENV TZ=Asia/Bangkok

# Copy SQL schema (รันครั้งแรกเท่านั้น)
COPY ./db/schema /docker-entrypoint-initdb.d/

EXPOSE 5432