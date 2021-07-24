<p align="center">
  <h1 align="center">Stripe GraphQL</h1>
  <p align="center">
    <img src="https://github.com/nhost/stripe-graphql/blob/master/logo.png" height="170">
  </p>
</p>

An open-source Stripe GraphQL integration with Nhost. Add Stripe to your app with just a couple clicks and use the Stripe API through the GraphQL schema.

## Documentation

GraphQL Schema documentation coming soon!

## Usage with Nhost

## Local Usage

```bash
git clone https://github.com/nhost/stripe-graphql.git
cd stripe-graphql
```

After cloning the repo and going to the project directory, run the following command to generate models and logic from the GraphQL Schemas:

```bash 
go generate ./graph
```

Now we can run the application either with Docker or the go cli tool.

#### Docker

```bash
docker build -t stripe-graphql .
docker run -d -p 8080:8080 -t stripe-graphql
```

#### Go

```bash
go run main.go
```

After running the application you can test out the GraphQL API with your favorite GraphQL client. *You must specify your Stripe Secret key as an http header:*

```HTTP
"X-Stripe-Secret-Key": "YOUR_STRIPE_KEY"
```

To get your Stripe secret key, go to [Stripe](https://stripe.com) and log in to your account. Then copy your secret key from the dashboard.
