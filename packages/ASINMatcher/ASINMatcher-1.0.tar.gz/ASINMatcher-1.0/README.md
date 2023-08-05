ASIN matcher is used to get the details of the product link on Amazon.

python install ASINMatcher_pkg

Usage:
import ASINMatcher_pkg

ASINMatcher_pkg.url_matcher(url) -- returns the object with all the details like region, id etc
ASINMatcher_pkg.is_valid_link(url) -- returns True if the url is invalid
ASINMatcher_pkg.get_region(url) -- returns the region of the url, if invalid url, return blank value
ASINMatcher_pkg.get_id(url) -- returns the id of the url, if invalid url, return blank value
ASINMatcher_pkg.get_id_type(url) -- returns the id type (e.g. "ISBN" or "ASIN") of the url, if invalid url, return blank value