// The cookie name
const COOKIE_NAME = "_btc_oid"

func main() {
  // A memcache instance using Go Memcache (github.com/bradfitz/gomemcache/memcache)
  mc := memcache.New(memcacheHosts)

  // ...

  http.HandleFunc("/cart.json", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Read the cookie
    cookie, err := r.Cookie(COOKIE_NAME)

    // If no cookie set, return empty JSON cart
    if err != nil {
      emptyCart(w)
      return
    }

    // Build the expected cache-key for Memcached, ie. "cart:XXXXX"
    s := []string{"cart", cookie.Value}
    key := strings.Join(s, ":")

    // Get cart data from Memcache
    data, err := mc.Get(key)

    // No data in Memcache. It might have been purged from cache.
    // Notify upstream Nginx so it can forward it to backend Rails app.
    if err != nil {
      log.Println("No data in memcache", err)
      http.Error(w, err.Error(), http.StatusConflict)
      return
    }

    // Data found in memcache. Serve it as JSON.
    w.Header().Set("X-Proxy-Status", "Cache hit")
    w.Write(data.Value)
  })

  // Start the server
  log.Println("Starting HTTP server on", httpHost, "Memcached on", memcacheHosts)
  log.Fatal(http.ListenAndServe(httpHost, nil))
}