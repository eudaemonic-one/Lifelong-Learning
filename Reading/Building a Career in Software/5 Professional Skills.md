# Chapter 5: Professional Skills

## Project Management

* Project management is organizing technical work to deliver complex projects: planning, tracking, and communication to divide work into coherent pieces, allocate the right resources, do the right things at the right times, remember the loose ends, and unblock engineers.
* **Motivation**
  * Project costs must be estimated and weighed against returns; work must be divided to map onto engineers; schedules must be communicated to stakeholders; key pieces must be tracked so what ships is complete.
* **The Foundation: A List and a Graph**
  * The essence of project management is knowing what needs to be done and how best to deploy people to do it.
  * The essential methods:
    * Keeping a list of all outstanding tasks - a spreadsheet, text file, JIRA board, whatever.
    * Understanding the dependency graph of those tasks - a flowchart or a DAG.
      * => informs prioritization.
      * => let you schedule in the senses of both assigning resources and predicting timelines.
  * => estimate timelines, ask for resources, run meetings, and send updates to stakeholders.
* **Keeping Stakeholders Informed**
  * People care about timeline, what's done, what's not done, implications for downstream systems.
  * Update stakeholders through email or meetings
* **Keeping Close Collaborators Informed**
  * Ping close collaborators freely with updates on minor progress hand hiccups, ideas I'm having, questions, thanks for their progress.
    * “quick fyi—i just landed the build changes, you should be able to adopt the new rule now, please let me know if you have any problems.”
    * “hi anant, sorry i didn’t get to your code review earlier. i just added some comments. i’ll treat it as my top priority for the rest of the day, so feel free to ping me if you update the diff.”
    * “hey evelyn, I just found a bug in my script. it might not rebuild the coordinator config when you do a dev build. i’m working on a fix right now.”
    * “yo, I just saw that you landed a fix to the reconnect problem we were hitting—that’s awesome, thanks so much.”
    * “[to a chat room] hey everybody, if you’re seeing failures when turning up a dev cluster, it’s probably the bug Jeff is working on. I just landed a fix to the workaround script and updated this channel header with a link to the workaround runbook. sorry for the inconvenience, we’ll have a proper fix asap.”
* **Schedules and Estimation**
  * **Four Laws**
    * **They're usually artificial, but you still need to care.**
    * **Estimation is the hardest problem in software.**
    * **Therefore, pad your estimates by 2x.**
    * **Overcommunicate your status.**
      * A prompt, clear notification that a schedule is changing can be a sign of professionalism appreciated by one and all.
      * “Hi Folks,
        Today, we realized that before we can turn on the new session cache, we need a way to shoot down sessions on account deletion. We already have shootdown for password change, but hadn’t previously considered the account deletion case. We estimate that this will add a week to the total time to delivery, meaning that our new estimated date for go-live is March 15th. Please reach out with any concerns.
        Cheers,
        Yun”
  * **Schedules and Estimation: The Basic Algorithm**
    * Sit down with a pen and paper.
    * List the major pieces of your work; draw your dependency graph.
    * Estimate the time for each piece.
    * Sum the estimated time for each work item that cannot or will not be done in parallel.
    * Add time for testing; make it such longer than you expect, because you are guaranteed to underestimate.
    * Double the total.
    * Resist requests to reduce that estimate.
  * **Leading a Project**
    * Deliver the project, full stop.
      * => advocate for the resources
      * => weight the quality trade-offs
      * => estimate the timeline
      * => persuade and coerce the help you need
      * => when things seem bad, you do the exercise of thinking if you were your boss, asking yourself which of your assumptions about resources or requirements could be thrown out.
    * To update the dependency graph
      * Bootstrap your project understanding by researching and meeting with stakeholders and participants.
      * Draw your actual dependency graph and estimate timeline; write an accompanying concise project plan. Ask participants to estimate and commit to their own timelines.
      * Communicate this plan and rough timelines (conservatively estimated) to participants.
      * Set up your tracking.
      * Repeat weekly
        * **Communicate:** weekly meeting led by you and ruthlessly kept short, send a weekly status update up the chain or to other stakeholders.
        * Update your tracking and dependency graph.
        * Follow-up 1:1 to go deeper on areas of concern or uncertainty.
        * Review progress against your timelines, and unblock the pieces that get blocked, whether by nagging, advising, or doing things yourself. Update stakeholders immediately if timelines change.
        * Execute on your personal piece of the project.
* **The Fine Art of Nagging**
  * **What to Say**
    * We start with a strong grasp of the project plan and frankness about what's needed.
    * We encourage, charm, thank, and come through for others.
    * Acknowledging others' struggles.
      * “Hey Kara, just checking in—are we still on track for rollout in two weeks? Is there anything I can help with?”
      * “Hey Kara, have you fixed that uncaught issue yet? It’s blocking the rollout.”
      * “Hey Youssef. I wanted to check in with you about your progress on the secrets migration. How are things going? I know your team’s usage is a bit unusually complicated—do you folks have the resources to get this done by the end of the month, or would some help from the secrets side be useful?”
  * **When Gentle Requests Don't Work**
    * Increase Intensity => Use a Project Manager If You Can => Go Up the Chain.
      * “Hey Vijay, just following up again about the uncaught bug. Sorry to keep bugging you, but this has now become a major problem for the on-call rotation; would you be able to look at it today?”
      * “Hey Sandra, I just want to put this task on your radar; it’s been lingering for a while, and I think we need to make sure someone follows-up on it, because it’s impacting on-call with several alerts per shift.”
      * Asking a project manager to follow-up is milder than going to an engineer manager, because project managers' essential function is to understand outstanding work and ensure that it happens; their follow-up doesn't involve an implication of negligence, unlike reaching out to someone's manager.
  * **Advanced Unblocking: Doing It Yourself**
    * Do their job for them => solve your problem, might set an inspiring example of vigor and ownership.
      * “Hi Francisco. Since I know your team is slammed, I thought I might take a shot at this project to unblock the bigger effort. I’ll check in with you along the way and make sure your team approves the design and code. What do you think?”
  * **How Often to Follow-up**
    * Err toward less when you can, because it's stressful to be asked for status.
      * for a critical production issue => every 120 seconds.
      * for a broken tool that's blocking engineers => an hour.
      * for a tricky bug blocking a release => once or twice a day.
      * for an important project taht needs to land months in the future => once every week or two.
* **Paying for Information**
  * We prioritize tasks that yield information and reduce uncertainty.
    * Prototyping => building fast and at low quality => answering the most important questions.

## Effective Meetings

* Ingredients of effective meetings:
  * **Throughout**: An owner who takes overall responsibility for making a meeting productive.
  * **Before**: A clear agenda, with context, sent out beforehand.
  * **During**: Punctuality, minimal distraction, and above all, steady focus on reaching conclusions as quickly as possible.
  * **After**: Clear meeting notes sent afterward to confirm the decisions of the meeting.
* **Throughout: An Owner**
  * An owner makes sure that the preparation, conversation, and follow-up happen smoothly.
    * e.g., send the agenda, shepherd an efficient conversation, and write the meeting note.
* **Before: Agenda**
  * We need to push hard to make everyone move fast and to make every else pay attention - try to get the information truly shared and everyone back on task.
  * A meeting agenda example:

```text
Calendar Invitation Title: Meeting re. Database Choice for Next Gen Backend

Calendar Invitation Body:

The goal of this meeting is to make a plan for deciding on the database for our next generation backend.

- How to evaluate performance.

- Sketch budget and possible tradeoffs.

- Confirming level of support from Infrastructure team.

- Back of envelope analysis of impact on system design.

- Owners for follow-up in each area.

Manuel will open the meeting with a brief review of scalability issues with the current system and the known requirements for next generation.

Pooja will then give an overview of the infra team’s initial investigation and the short list of options to consider.

Finally, we’ll have open discussion on the preceding points in that order.
```

* **During: Punctuality**
  * Set a tone of business-like precision => start on time and get done early.
  * Once you have quorum, start the meeting.
* **During: Start with Agenda Review**
  * “Thanks for coming, everyone! As you know, we’ve decided to move off of SimpleDB for our next gen backend, and we need to decide how our teams will work together to make the choice.”
  * “Our goals in this meeting are to confirm a plan for choosing a database, especially identifying an owner and PoCs for driving the decision, but also running through key concerns for performance testing, budget trade-offs, and ownership/support.”
  * “Manuel will start by reviewing the current system’s problems and requirements we know about, and then Pooja will discuss the options the DB team has come up with in their initial investigation. After that we’ll talk about performance, budget, support, and next steps/owners.”
  * “Let’s get started—Manuel?”

* **During: Focus on Reaching Conclusions**
  * Every meeting should achieve actionable conclusions.
    * “What would help us reach a conclusion here?”
    * “Could we table this for now and focus on <the main question>?”
    * “This is definitely an important question, but since we’re short on time, maybe we should move on to the next topic and circle back later?”
    * “I’m not sure if we can decide this right here—<person>, maybe you could analyze this offline and email us your assessment?”
    * “Do we have the right people to figure this out? Should we pause this conversation until we can get input from <an expert>?”

* **During: Review Action Items**
  * The meeting should end with a review of clear action items with clear owners.
    * “Okay, so our action items are: Erika to send guidance to the team and Donald to start the rollout tomorrow morning. I’ll send the meeting notes.”

* **After: Meeting Notes**
  * Crisp meeting notes:
    * the purpose of the meeting
    * who attended
    * what was discussed
    * what was concluded/agreed on
    * what the action items are
  * A write-up minimizes the risk of conflicting memories of what you agreed on.
* **A Note on Meeting Etiqutte**
  * It means yes to
    * Respectfully but proactively introducing ourselves
    * Pay attention
    * Waiting our turn to speak
    * Not dominating the conversation
    * Giving everyone a chance to say their piece; making space for the shy or quiet and soliciting their opinions.
  * People shouldn't open their laptops if it can be avoided.

## Productivity and Organization

* **Time Management and To-Do Lists**
  * To-do app
    * Data must be shared between phone and desktop, so you can create to-dos wherever you are.
    * Creating a simple to do should be frictionless, so you do it.
  * Get Things Done
    * Keeping contextual lists
    * Regularly reviewing and prioritizing your list of tasks => so that you can quickly find the most urgent ones.
* **Effective Calendaring**
  * Always include a video conference link.
  * Give your attendees edit permissions on the invite.
  * Include a crisp agenda in your meeting invite; mention each topic you plan to cover and what the *goals* of the meeting should be.
  * Defragment other people's time; if they have a packed calendar, try to find a slot adjacent to their other meetings, and especially try to avoid creating small gaps between meetings.
  * When scheduling by email, always offer a mewnu of acceptable times.
    * “Meeting next week sounds great. Would one of the following slots work?”
* **Caching Knowledge**
  * An infinitely scrolling Google Doc, your cheat sheet, with a little section for each neat trick you figure out.
  * A text file on your computer, your command cache, with commands you've been using particularly recently.
* **Scripting Everything**
  * If you type a command more than five times, you should script it.
* **Living in the Cloud**
  * All your dot files and scripts in a private Git repository
  * All the documents safely striped to a hundred servers (Google Docs)
  * Never restore from backup.
* **Just Do Something**

| Issue                                                        | Why I Can't Do Something                   | Yes You Can! Try This                                        |
| ------------------------------------------------------------ | ------------------------------------------ | ------------------------------------------------------------ |
| We're blocked waiting for another team's bugfix.             | They haven't commented on the bug.         | Email the engineer or their manager. Chat someone. Call them on the phone. Ask your manager or a more senior person to escalate. |
| Requests are failing.                                        | I don't know what dashboard to look at.    | Search for the service name on the metrics home page. Look at a global dashboard. Call someone and ask. @here in a chat room and ask. Page someone else. |
| We can't land changes because the build is broken.           | I don't know how the build works.          | Figure it out. Ask someone. Go to the Jenkins page and find the logs. Use git bisect, which lets you find the culprit even without semantic understanding. |
| My colleague blocked my diff with code review comments I disagree with. | He's too obstinate, I can't make progress. | Meet face to face with the colleague, and try to understand teh feedback better. Get a second opinion from a colleague you trust, and ask if they have an idea for how to handle it tactfully. Ask your manager for advice or to intervene. Just go around them. |

* **Keep a Worklog**
  * Log every single thing you do at work
    * every bug I fix
    * every proposal I write
    * every outage I work on
    * every presentation I give
    * every person I mentor and what I mentor them about
  * Asynchronously process to-dos and put them into doc.
    * “worklog: incident commander for configuration service outage 10/14”

## Hiring Engineers: Interviewing from the Other Side of the Table

* **Interview Focus Areas**
  * **Coding**
  * **Architecture**
  * **Communication**
  * **Domain Knowledge**
  * **Organized Thinking**
* **The Nuts and Bolts of Interviewing**
* **Decisions**

## Software Methodologies and the Generality of Good Engineering

* Project management is always about understanding dependencies, staying organized, and conservative estimation; collaboration is always about crisp communication, empathy, and efficient use of your colleagues' time; software that works is software that works.
* When you finish something, you absolutely must go find something else to do, asking your boss if necessary.
* **Agile Software Development**
  * short development cycles that integrate requirements gathering, design, and development, constantly shipping small increments of working software, bringing together stakeholders and developers to review and access, and replanning as circumstances change.
* **Scrum**
  * a rhythm of self-contained "sprints", typically periods of 1-4 weeks, and codifies specific roles and meetings to structure those sprints.
  * => collaboration with all team members contributing to planning, implementation.
  * => team introspection and daily heping each other overcome blockers.
  * Structure
    * First, a Sprint Planning Meeting where the team works together to set tractable goals for the sprint.
    * For the duration of the sprint, a short daily "scrum" or "stand-up", help each other resolve "blockers" and cross-pollinate useful information.
    * At the end of the sprint, a Sprint Review and a Retrospective for reviewing progress, demoting results, and identifying opportunities to improve.
    * Asynchronously, the team regularly reviews and prioritizes a "backlog" of future work.
