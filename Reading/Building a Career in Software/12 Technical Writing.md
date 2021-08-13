# Chapter 12: Technical Writing

* Eight foundational principles common to all technical writing:
  * **Be concise** => holding precision constant.
  * **Emphasize conclusions** => readers usually care most about high-level conclusions.
  * **Put yourself in your readers' place** => ask yourself what matters to them, what they know, and whether they will understand what you've written.
  * **Guide readers with clear structure** => help readers follow your writing.
    * start with an introduction/summary, follow with clearly titled sections addressing specific areas, and close with a conclusion.
    * use bullted lists liberally.
  * **Use simple language** => for readability.
  * **Provide ample context** => for them to reason about and know the relevance.
  * **Lead with your most important points** => clearest for readers.
  * **Edit for precision** => does it say exactly what I mean?
* Two most common problems for engineers:
  * => failing to provide enough context.
  * => giving too much detail.

## A Note About Strutcure

* Our structure tool:
  * A razor-sharp introduction that summarizes everything else.
  * Sections delineated by headers or paragraph breaks.
  * Section ordering that sorts first for comprehensibility - necessary context first - and then by importance - most important points in early, before the reader gives up.

## Editing

* **Remove Words**
  * “This is the simplest, ~~easiest,~~ and ~~by far the~~ most effective editing technique ~~of which~~ I know; we ~~go back through our text and~~ reread every ~~single~~ sentence ~~we’ve written,~~ identify ~~all the words we’ve added that aren’t strictly necessary~~ unnecessary words, and ~~take out (or, in some cases, simplify)~~ remove them.”
* **Simplify Sentences**
  * Subject-Action-Object structure
  * “The infrastructure team is migrating the data to the cloud. This migration blocks big schema changes. They expect to finish by the end of the week, and we can launch this feature at that time.”
* **Get the Context Right**
  * We can ask:
    * Does the reader know everything they need to know to understand what I've said?
    * Will the reader believe me, or do I need to justify myself?
    * Am I explaining something the reader already understands?
  * “For financial transaction data, we will need greater consistency than OurDistributedDataStore guarantees. Therefore, those records will be stored in OurReplicatedACIDDatabase, which should offer sufficient throughput for this workload. Our application will be written in Java, which supports Object Oriented Programming; inheritance will allow good code reuse.” => “For financial transaction data, we will need greater consistency than OurDistributedDataStore guarantees; account balances for each user must be stored in separate rows, and transfers between them must be atomic. OurDistributedDataStore can only guarantee consistency for transactions on a single row. Therefore, those records will be stored in OurReplicatedACIDDatabase. While it does not offer the same throughput, we have load tested it up to 15,000 transactions per second, 10x the maximum projected QPS in 1 region.”
* **Spelling, Grammar, Capitalization, and Punctuation**
  * Trust your spellchecker and dictionary.
  * Start sentences with capitals, end them with periods.
  * Use trustworthy editor tools.
