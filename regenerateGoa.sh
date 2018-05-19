mv ./vendor/ ./prevendor/
goagen bootstrap -d github.com/mtenrero/ATQ-Director/http/design
mv ./prevendor/ ./vendor/