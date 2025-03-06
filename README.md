# fastRTB
fast and privacy aware ad bidding protocol 

fastRTB is designed as an faster alternative to openRTB and with features for an privacy aware marketplace.

Current implementation is server to server but a protobuf enabled version that runs within browsers should be possible to implement.


also the goal would be to enhance the following topics:

- metrics/reporting endpoint: the client shall specify endpoints with a well defined grpc service where the server or the browsers can send metrics towards
- a flexible protocol for more granular interaction metrics within the creatives
- bulk bid requests, where a bidRequest could span a bigger number (e.g. 1000) impressions so that the publisher can fill them during a given timeframe and report back. 