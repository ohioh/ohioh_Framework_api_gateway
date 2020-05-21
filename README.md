<hr />
<hr />
<hr />
------------------ WORK IN PROGRESS
<hr />
<hr />
<hr />


<p align="center">
 <a href="https://ohioh.de"><img src="https://github.com/https-github-com-ohioh/Documentation/blob/master/images/solution_architecture/csm_logo_ohioh_6edae6fe4c.png" width="400"></a>
</p>

<hr />
<p align="center">
    <a href="#about-this-pct">About this Project</a> • 
    <a href="#who-we-are">Who We Are</a> •
    <a href="#credits">Credits</a> •
    <a href="#data-privacy">Data Privacy</a> •
    <a href="#code-of-conduct">Code of Conduct</a> •
    <a href="#working-language">Working Language</a> •
    <a href="#our-documentation">Our Documentation</a> •
    <a href="#licensing">Licensing</a> •
    <a href="#how-to-contribute">How to Contribute</a> •
    <a href="https://ohioh.de">Web Site</a>
</p>
<hr />

NOTE: This README is also available [in German](translations/README.de.md). Thank you for understanding that the German version might not always be up-to-date with the English one.

HINWEIS: Diese README ist ebenfalls [auf Deutsch](translations/README.de.md) verfügbar. Bitte haben Sie Verständnis, dass die deutsche Version nicht immer auf dem gleichen Stand wie die englische Version ist.

## About this Project

OHIOH APP is an open source implementation in the OHIOH Framework.

OHIOH App is a part of the OHIOH Products, which uses privacy-preserving protocol for community-driven contact tracing across borders, with the opportunity to select geolocation tracing, like GPS.

Its aim is to flatten the curve of pandemics and potentially detect them before they spread widely. To effectively contain the spread of a virus health officials use the approach of test, trace and quarantine.

When an individual is determined to be infected , the official’s task is to determine who might have been in close enough proximity to this person, to be infected. Testing/ quarantining those individuals quickly is one of the effective strategies to curb the spread of infection further. However, such a method, without information is ineffective as well as inefficient.
	
The OhIOH APP, allows participating devices to log encounters with each other, in order to facilitate epidemiological contact tracing, while protecting the user's personal data and privacy. 
	
The OHIOH reference implementation comprises:

 [Xcode developer tools](https://developer.apple.com/xcode/downloads/) 
 <hr />
 
 Frontend:

- Android app: [https-github-com-ohioh/
ohioh-app-kotlin](https://github.com/https-github-com-ohioh/ohioh-app-kotlin/)


- iOS app: [https-github-com-ohioh/ohioh-app-swift](https://github.com/https-github-com-ohioh/ohioh-app-swift)

- API Gateway: [https-github-com-ohioh/ohioh_api_gateway](https://github.com/https-github-com-ohioh/ohioh_api_gateway)

<hr />

Cluster 1:
1. Flask Service: [https-github-com-ohioh/ohioh_Framework_Cluster_1_Flask](https://github.com/https-github-com-ohioh/ohioh_Framework_Cluster_1_Flask)

2. MongoDV Service: [https-github-com-ohioh/ohioh_Framework_Cluster_1_Mong_DBo](https://github.com/https-github-com-ohioh/ohioh_Framework_Cluster_1_Mong_DBo)

3. Go Service: [Code coming soon](https://ohioh.de)

4. Firewall Service: [Code coming soon ](https://ohioh.de)

<hr />

Cluster 2:
1. Flask Service: [https-github-com-ohioh/ohioh_Framework_Cluster_2_Flask](https://github.com/https-github-com-ohioh/ohioh_Framework_Cluster_2_Flask)

2. MongoDV Service: [https-github-com-ohioh/ohioh_Framework_Cluster_2_Mong_DBo](https://github.com/https-github-com-ohioh/ohioh_Framework_Cluster_2_Mong_DBo)

3. Go Service: [Code coming soon](https://ohioh.de)

4. Firewall Service: [Code coming soon ](https://ohioh.de)
<hr />

Cluster 3:
1. Flask Service: [https-github-com-ohioh/ohioh_Framework_Cluster_3_Flask](https://github.com/https-github-com-ohioh/ohioh_Framework_Cluster_3_Flask)

2. MongoDV Service: [https-github-com-ohioh/ohioh_Framework_Cluster_3_Mong_DBo](https://github.com/https-github-com-ohioh/ohioh_Framework_Cluster_3_Mong_DBo)

3. Go Service: [Code coming soon](https://ohioh.de)

4. Firewall Service: [Code coming soon ](https://ohioh.de)
<hr />

Cluster 4:
1. Flask Service: [https-github-com-ohioh/ohioh_Framework_Cluster_4_Flask](https://github.com/https-github-com-ohioh/ohioh_Framework_Cluster_4_Flask)

2. MongoDV Service: [https-github-com-ohioh/ohioh_Framework_Cluster_4_Mong_DBo](https://github.com/https-github-com-ohioh/ohioh_Framework_Cluster_4_Mong_DBo)

3. Go Service: [Code coming soon](https://ohioh.de)

4. Firewall Service: [Code coming soon ](https://ohioh.de)
<hr />

Cluster 5:
1. Flask Service: [https-github-com-ohioh/ohioh_Framework_Cluster_5_Flask](https://github.com/https-github-com-ohioh/ohioh_Framework_Cluster_5_Flask)

2. MongoDV Service: [https-github-com-ohioh/ohioh_Framework_Cluster_5_Mong_DBo](https://github.com/https-github-com-ohioh/ohioh_Framework_Cluster_5_Mong_DBo)

3. Go Service: [Code coming soon](https://ohioh.de)

4. Firewall Service: [Code coming soon ](https://ohioh.de)
<hr />

Cluster AI:
1. AI_Accumulator_Service: [https-github-com-ohioh/ohioh_AI_Accumulator_Service](https://github.com/https-github-com-ohioh/ohioh_AI_Accumulator_Service)

2. AI_MongoDB_Service: [Code coming soon](https://ohioh.de)

3. AI_Go_Service: [Code coming soon](https://ohioh.de)

4. AI_Firewall_Service: [Code coming soon ](https://ohioh.de)
<hr />



## Who We Are

we are many :-) https://ohioh.de/?page_id=661

## Credits

We'd like to thank all the partners who have been involved in this enormous project from the beginning. The project would not be where it is today without all the exploration and great work that had already been done around the tcn-coalition

## Data Privacy

In this project we are strictly observing the principles of the General Data Protection Regulation (GDPR) to protect the users’ privacy. We are processing necessary data only - exclusively for the purpose to let users know if they have come into close contact with other infected users, without revealing each other's identity. The compliance with these regulations is safeguarded by several procedures, e.g. by implementing technical and organizational measures adhering diligently to the high standards of the GDPR. Of course, the app will provide users with a comprehensive privacy statement to be as transparent and clear as possible. As we are developing the app as an open source project, the community can review it. We appreciate your feedback!

## Code of Conduct

This project has adopted the [Contributor Covenant](https://www.contributor-covenant.org/) in version 2.0 as our code of conduct. Please see the details in our [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md). All contributors must abide by the code of conduct.

## Working Language

We are building this application for Germany. We want to be as open and transparent as possible, also to interested parties in the global developer community who do not speak German. Consequently, all content will be made available primarily in _English_. We also ask all interested people to use English as language to create issues, in their code (comments, documentation etc.) and when you send requests to us. The application itself, documentation and all end-user facing content will - of course - be made available in German (and probably other languages as well). We also try to make developer documentation available in German, but please understand that focussing on the _Lingua Franca_ of the global developer community makes the development of this application as efficient as possible.

## Our Documentation

This repository contains the developer documentation and related content.

### Project Scope

- [OHIOH-APP - Scoping Document](scoping_document.md)

### Technical Documentation
The technical documents are intended for a technical audience and represent the most recent state of the architecture. The solution architecture and concepts might change over time as external dependencies (e.g. the framework provided by Apple/Google) are still changing and new requirements need to be included or existing ones change. We appreciate feedback to all elements of these technical documents. 
* [OHIOH-APP - Solution Architecture](solution_architecture.md)

To be published:
* OHIOH-APP - Mobile Architecture for iOS and Android
* OHIOH-APP - Server Architecture
* Verification Server Architecture
* Portal Server Architecture
* System Operation
* Technical Security Concept
* Data Privacy Concept

### Glossary
For an easier understanding of the used acronyms and special terms in our documents please see our [glossary](glossary.md).

## Licensing

Copyright (c) 2020 Tjark Ziehm,.

Licensed under the **Apache License, Version 2.0** (the "License"); you may not use this file except in compliance with the License. 

You may obtain a copy of the License at https://www.apache.org/licenses/LICENSE-2.0.

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the [LICENSE](./LICENSE) for the specific language governing permissions and limitations under the License.

## How to Contribute

Please see our [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute, our team setup, the project structure and additional details which you need to know to work with us.
