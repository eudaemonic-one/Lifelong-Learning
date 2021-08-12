# Chapter 7: Shining in an Engineering Organization

## Reputation and Personal Brand

* Good reputation => that people want to work with you, that they listen to your ideas, that they trust you with important work, and that they give you the benefit of the doubt for advancement.
* Steadily doing good work and treating people right are enough to build a good reputation.
* Becoming the go-to person for some technology.

## Helpfulness

* You should help colleagues every chance you get - not grudgingly, but enthusiastically, hustling when they need something fast, digging if you don't know an answer off the top of your head, take an interest in their deadlines and frustrations, apologizing if you accidentally go slower than you could, connecting them to others when you can.

## Running Toward Fires

* When your team's software has an outage, your tests break, or users are hitting some kind of bug, you should almost always take the interrupt and lean in to help.
  * => the right thing to do.
  * => build a brand as a responsible owner and go-to person in your area.

## Politics and Political Capital

* Politics
  * => at the "necessary" end are campaigns for good ideas, like explaining an important idea 1:1 to several people to marshal support a valuable business decision.
  * => at the "evil" end are treacheries like slandering someone to an executive or killing a valuable project by quietly withholding a tiny bit of necessary support.
* **The Necessary Form**
  * We need to practice the mild form of politics to execute on big ideas, because they always require big decisions and the support of many teams and individuals.
  * When you need to get something done, keep in mind the following:
    * **Understand who makes decisions.**
    * **See what matters to other stakeholders in a decision.**
      * => understand how your idea may benefit or harm them.
      * => identify compromises that may balance those concerns for various people invovled.
      * => honestly acknowledging the cost while emphasizing the benefits.
    * **Reason about what different audiences know and don't know.**
    * **Recognize who may be inconvenienced or hamred by a decision.**
      * and how to restore their support.
    * **Understand the strength of our own standing, or *political capital*.**
* **Political Capital**
  * Political capital is the trust of decision-makers; people's inclination to do what you say and tell others to do the same, because they trust you.
  * Managers, executives, peers, and subordinates will do what you say to the extent that
    * They believe that you're saying the right things.
    * They like you personally, that is, they want to.
    * You're backed by real authority, that is, they have to.
  * assess through asking feedback and listening carefully to see how people do in fact respond to us.
* **The Evil Form**
  * The evil formulation either puts the interests of the individual ahead of those of the enterprise (e.g., trying to kill a good idea because it diminishes your influence) or draws on evil methods: bullying or dishonesty.

## Professionals Maximize Business Value: In Which Open Source is Stealing From the Shareholders Sometimes

* You exercise discretion over your own time and budget to maximize business value.
* Maintaining an open source project takes a lot of time and focus, reduces your freedom of movement, and generally distracts you from your primary job.

## Saying No Is Not Your Job

* "platform engineer" is responsible for the ongoing maintenance of highly technical component that teams shipping products depend on.
  * => they often have to say "no".
  * => you should always try to come with an alternative suggestion of how your colleagues can get their jobs done.

## Personal Reliability

* The most important ingredient to the confidence of peers and leaders is good execution.
* Essental practices:
  * **Doing what we say** => sending an email, showing up to a meeting, fixing a bug, reviewing a tiny diff.
  * **Punctuality**
  * **Being available** => answer questions promptly.
  * **Running toward problems** => try to fix tests, scripts, dashboards, alerts, outages, and production code.
  * **Resisting distraction** => spend less time slacking off.
  * **Businesslike language and body language.**

## Knowing Our Limits: The Epistemology of Software Engineering

* Best engineers => a robust theory and caution about the boundaries of one's own knowledge.
  * grasp what they well and truly know, what they believe, what they have good reason to suspect, and what they will and truly do not know.
  * when problem solving, that means that they can economically figure out what new knowledge would advance their understanding of a problem.
  * when collaborating, it means that they seldom embarrass themselves with false assertions.
* Catalog for yourself and others the level of confidence in your statements:
  * **What you have directly observed.**
    * “I saw an error message in the logs saying, “Unable to connect to database: connection refused.”
  * **What you believe.**
    * “I believe this service talks to the database over an HAProxy load balancer, which was true the last time I looked at the code, though I haven’t looked in a few months.”
  * **What you suspect.**
    * “I looked at the database metrics, and I don’t see any incoming requests on this dashboard link, so I’m suspecting that we may be having a problem with the HAProxy upgrade that’s been going on this week.”
