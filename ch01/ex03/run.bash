for i in 1000, 5000, 10000, 50000, 100000, 500000; do
    go run main.go `python -c "print 'a ' * $i"`
done