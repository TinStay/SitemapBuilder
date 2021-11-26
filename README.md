# Sitemap Builder


## Exercise details

A sitemap is basically a map of all of the pages within a specific domain. They are used by search engines and other tools to inform them of all of the pages on your domain.

One way these can be built is by first visiting the root page of the website and making a list of every link on that page that goes to a page on the same domain. For instance, on `calhoun.io` you might find a link to `calhoun.io/hire-me/` along with several other links.

Once you have created the list of links, you could then visit each and add any new links to your list. By repeating this step over and over you would eventually visit every page that on the domain that can be reached by following links from the root page.


Sitemap builder should output the data in the following XML format:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/</loc>
  </url>
  <url>
    <loc>http://www.example.com/dogs</loc>
  </url>
</urlset>
```


### Notes

**1. Be aware that links can be cyclical.**

That is, page `abc.com` may link to page `abc.com/about`, and then the about page may link back to the home page (`abc.com`). These cycles can also occur over many pages, for instance you might have:

```
/about -> /contact
/contact -> /pricing
/pricing -> /testimonials
/testimonials -> /about
```


**2. Following packages are used.**

- [net/http](https://golang.org/pkg/net/http/) - to initiate GET requests to each page in your sitemap and get the HTML on that page
- [encoding/xml](https://golang.org/pkg/encoding/xml/) - to print out the XML output at the end
- [flag](https://golang.org/pkg/flag/) - to parse user provided flags like the website domain


## Bonus

As a bonus exercises you can also add in a `depth` flag that defines the maximum number of links to follow when building a sitemap. For instance, if you had a max depth of 3 and the following links:

```
a->b->c->d
```

Then the sitemap builder would not visit or include `d` because more more than 3 links are followed to to get to the page.

On the other hand, if the links for the page were like this:

```
a->b->c->d
b->d
```

Where there is also a link to page `d` from page `b`, then your sitemap builder should include `d` because it can be reached in 3 links.
