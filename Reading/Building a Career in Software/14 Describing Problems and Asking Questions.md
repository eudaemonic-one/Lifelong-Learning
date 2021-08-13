# Chapter 14: Describing Problems and Asking Questions

* Principles:
  * Try a *reasonable amount* to help yourself before you take others' time/
  * Provide ample context, as you would with any other form of technical writing:
    * **Your intentions**
    * **Your observations**
  * Express your question as precisely as you can.
    * “I’ve been trying since 1 pm to deploy a logging change to the User Service (commit f752e2ee8) with the Prod Deploy Jenkins job. I see that the job fails, and there’s an error message in the Jenkins logs about artifactory not accepting a connection. I think the job is configured correctly—do you know how to check whether artifactory is healthy?”

## Magic Template: Asking For Help Debugging

```text
“I’m trying to <goal>. I’m doing <actions>, but I see <observations>. My best theory is <theory>, because <reason>. Can you <do something specific>?”
```

## Magic Template: Asking For Help with a Production Issue

```text
“We’re observing <observation> starting at <time with time zone>. It is likely causing <experience> for <percentage> of <class of users>. We think <theory> and want help <doing|investigating> <thing> to mitigate>.”
```

## Magic Template: Everything Else

```text
“I’m trying to <do a thing>. I <have done some things | already know some things>. To move forward, I want to <do or know something else>. <specific question>.”
```
