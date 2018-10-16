cd user
docker build -t hub:5000/user .

cd ../order
docker build -t hub:5000/order .


cd ../product
docker build -t hub:5000/product .

docker push hub:5000
docker push order:5000
docker push product:5000
