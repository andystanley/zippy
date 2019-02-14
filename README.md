# zippy
zippy is a persistent in-memory key-value store

## Getting Started
**Clone the repository**  
```
git clone https://github.com/andystanley/zippy.git
```

**Run the zippy CLI**  
```cd zippy 
bin/zippy-cli
```

By default zippy will create a log file in the same directory that you run it. If you prefer to use a different path try something like:  
```
bin/zippy-cli -path /Users/andystanley/Desktop/log
```

**Let's store some  data!**
```
Welcome to zippy!
>>> set name andrew
>>> get name
andrew
>>> set sport soccer
>>> keys
andrew
sport
>>> exit
```
Awesome we can store some data and access it really quickly! But will our data be there if we open it again?
```
Welcome to zippy!
>>> keys
andrew
sport
>>> exit
```
Of course it will! If you're curious to see how this works check out the **How Does Zippy Work?** section

## Available Commands
`set key value` Stores a key with the value provided  
`get key`       Retrieves a key's value  
`keys`          Displays a list of all keys  
`delete key`    Deletes a key and it's value  
`clear`         Clears the zippy store and deletes the log file  
`exit`          Close zippy  

## How Does Zippy Work?
Zippy is inspired by [Redis](https://redis.io/). Everytime a a `set` or `delete` operation is called a line is appended to the `log` file. This is called an Append Only Log. When the `zippy-cli` is launched, zippy reads the `log` file into memory. This allows zippy to be both performant and persistent.
