package domain

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// LinkInfo represents the comprehensive metadata of a web page that is scraped and stored by the service.
// This struct is designed to hold all relevant information about a web page, making it easy to retrieve
// and display metadata for content management.
type LinkInfo struct {
	ID          ulid.ULID     `bson:"_id" json:"id"`                  // The unique identifier for the web page entry, represented as a ULID. This field serves as the MongoDB _id, ensuring that each document is uniquely identifiable and easily sortable in a database context.
	URL         string        `bson:"url" json:"url"`                 // The URL of the web page, which serves as the primary key for identifying the page being scraped. This field is crucial for linking to the actual web content and for ensuring that the correct metadata is associated with each URL.
	Title       string        `bson:"title" json:"title"`             // The title of the web page, typically extracted from the HTML <title> tag. This title is used by search engines and social media platforms to display information about the page, influencing click-through rates and user engagement.
	Description string        `bson:"description" json:"description"` // The meta description of the page, found in the <meta name="description"> tag. This brief summary is critical for SEO, as it provides context for search engine results and is often displayed in previews when the page is shared on social media.
	Keywords    []string      `bson:"keywords" json:"keywords"`       // An array of keywords associated with the page, which are often sourced from the <meta name="keywords"> tag. These keywords can enhance search engine optimization by categorizing the page’s content, making it more discoverable.
	Favicon     string        `bson:"favicon" json:"favicon"`         // The URL of the favicon for the page, typically extracted from the <link rel="icon"> tag. The favicon is a small icon displayed in browser tabs and bookmarks, helping users visually identify the website at a glance.
	OpenGraph   OpenGraphData `bson:"openGraph" json:"openGraph"`     // The Open Graph metadata associated with the page, represented by the OpenGraphData struct. This metadata is essential for controlling how the page appears when shared on social media platforms, impacting user engagement and traffic.
	LastScraped time.Time     `bson:"lastScraped" json:"lastScraped"` // The timestamp of when the page metadata was last retrieved. This information is vital for determining the freshness of the data, allowing the service to schedule regular updates and ensure users access the latest information.
}

// OpenGraphData represents the Open Graph metadata used for enhancing the presentation of the web page
// when shared on social media platforms. This struct includes specific fields that define how the page
// should be displayed, including a title, description, image, and canonical URL.
type OpenGraphData struct {
	Title       string `bson:"title" json:"title"`             // The Open Graph title for the page, typically defined in the <meta property="og:title"> tag. This title may differ from the standard HTML title and is crucial for creating an attractive social media preview.
	Description string `bson:"description" json:"description"` // The Open Graph description, provided in the <meta property="og:description"> tag. This description is used to summarize the page's content when shared, playing a key role in user engagement and click-through rates.
	Image       string `bson:"image" json:"image"`             // The URL of the image specified by the Open Graph protocol, found in the <meta property="og:image"> tag. This image is prominently displayed in social media previews, enhancing visual appeal and providing context about the page's content.
	URL         string `bson:"url" json:"url"`                 // The canonical URL of the page, defined in the <meta property="og:url"> tag. This ensures that all links to the page are consistently referenced by the same URL, even if they contain query parameters or other variations.
}
