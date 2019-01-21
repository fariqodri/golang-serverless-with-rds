for service in services/*; do
  (cd $service && serverless deploy --verbose)
done