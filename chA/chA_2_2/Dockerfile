FROM alpine:3.7

ENV VERSION=0.11.5

RUN apk --no-cache add --virtual=build-deps curl gnupg

# terraform과 배포 파일 체크섬 정보를 받음
ADD https://releases.hashicorp.com/terraform/$VERSION/terraform_${VERSION}_linux_amd64.zip .
ADD https://releases.hashicorp.com/terraform/$VERSION/terraform_${VERSION}_SHA256SUMS .
ADD https://releases.hashicorp.com/terraform/$VERSION/terraform_${VERSION}_SHA256SUMS.sig .

# Keybase에 공개된 Hashicorp 사의 PGP 공개키 임포트
RUN curl https://keybase.io/hashicorp/pgp_keys.asc | gpg --import

# 체크섬의 전자서명 검증
RUN gpg --verify terraform_${VERSION}_SHA256SUMS.sig terraform_${VERSION}_SHA256SUMS

# 압축 파일의 체크섬 검증
RUN cat terraform_${VERSION}_SHA256SUMS | grep linux_amd64 | sha256sum -cs

RUN unzip terraform_${VERSION}_linux_amd64.zip
RUN mv terraform /usr/local/bin/
RUN terraform -v
