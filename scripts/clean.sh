rm -rf ./vendor Gopkg.lock

for service in services/*; do
  rm -rf ${service}/bin
done