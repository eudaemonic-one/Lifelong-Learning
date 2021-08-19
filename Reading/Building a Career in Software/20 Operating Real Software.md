# Chapter 20: Operating Real Software

* The most important difference between academic software and professional software is that the latter must be operated.
* Once professional software works, it needs to keep working across hardware failures, operator errors, bad deploys, and changes in usage.

## Respect Production

* Deploying, releasing, changing configuration when other experts are online is safe-ish.

## On Call

* "On Call": problem-reporting system is routed to you and you commit to responding quickly to address problems.
* **Five Principles**
  * **Protect the customer**: the ultimate purpose of being on-call.
  * **Always pick up the phone**: if you don't, no one will. 
  * **Ensure chain of custody**: own the problem until it's done.
  * **Focus on mitigation**: deep understanding can wait and is very possibly a distraction => roll the bad change back first.
  * **Don't give up.**
* **Just Roll Back**
  * An outage isn't the time for nuance, and the deadline is virtually always less important than the breakage.

## Incident Response

* An indicent response is a coordinated team effort to mitigate a problem.
  * assessing impact, running multiple parallel debugging efforts or mitigations, communicating among ourselves and with stakeholders.
  * some problems and responses are highly amenable to pre-planning.
* **Example Incident Response Playbook**
  * Incident response playbook:
    * **Roles and how to assign them**: Avoid duplication and dropped balls.
    * **Fast mitigation checklist**: Make it impossible to forget the most common and valuable interventions.
    * **After-incident practices**: Use mistakes to make your technology and organization stronger.
  * ***Roles***
    * **Incident commander**: Direct others.
    * **Eyes [on dashboards]**: Monitor dashboards for changes in overall state of issue.
    * **Debugging/operations**: Flip switches, turn dials, and fix stuff.
  * ***Key Mitigation Checklist***
    * Check for recent deploys => If you find any, roll them back.
    * Check for high GC on the XYZ service => Restart if found.
    * Check for HW failure on master node => Fail over to hot standby if needed.
  * ***After-Incident Practices***
    * RCA (root cause analysis), AIR (after-incident report), post-mortem, or a simple 1-2 page writeup circulated to the team to capture key learnings.
* **Fear, Excitement, and Learning**
* **Incident Commands in Detail**
  * ***Telling People What to Do*** => fast, crisp coordination.
  * ***Tracking the Threads*** => a simple text file for main debugging threads.
  * ***Establish a Rhythm*** => observation, analysis, action, communication, and review/summary.

```text
Start:

1. Review status: Verbally review what’s going on; get everyone on the same page.

“We’ve been seeing a 10% rate of HTTP 500s for loading the home page for the last 15 minutes; we believe this is due to partial DB corruption.”

2. Look for mitigations: Do we know what might help? If so, do it. If not, get people digging in the areas that might help.

“Infra team, do you think we can attempt a rollback from snapshot? How long will it take us to assess that?”

3. Pull In others if needed: Do we have everyone we need? If not, go get them.

“If we can’t, is it possible we could write a script to fix the data in-place? Dorian, can you please page someone from the inventory team?”

4. Pause a moment: Let everyone take a breath. If people are actively debugging, give them a few minutes to do it; set a timer on your phone. A small break contributes to a sense of order, which keeps people calm.

“Alright, we’ll pause for five minutes while the inventory team works on the script. Everybody stay calm and take a breath.”

Go to Start
```

## Bus Factor > 1

* You should never be the only person who knows something operational important.
* When the operational environment changes, you need the information to be pushed onto operators' radar.
* **Example: New Component**

```text
Hi Folks,

As of today, we’ve started rolling out a new cache of user data intended to improve page load performance. The design proposal can be found here. We’ll be rolling out incrementally with this experiment config; in the event of a problem (such as seeing stale user data, which we believe to be very unlikely), you can feel free to set the rollout percentage to 0% with this configuration flag and reach out to Carlos and me. You can also see the rollout progress on this dashboard. Please feel free to reach out with any questions. Thanks!

Cheers,
Alice
```

* **Example: Operations Change**

```text
Hi all,

As of today, our previous load shedding configuration flag has been supplanted by a new config; if you need to tune emergency load shedding, please use update the new flag. The interface is unchanged, but the flag controls some cases that were not previously covered. Our operator wiki has been updated as well. Please feel free to reach out with any questions.

Thanks,
Bob
```

## Test Should Always Work

* Broken tests mean guessing at what really works and doesn’t; uncertainty about tests necessitates human testing.
