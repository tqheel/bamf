id: 20200528
title: CLI Tool For my Blog Written in Go
postDate: May 28, 2020

I recently moved from app development at LexisNexis to the DevOps group, officially known as Platform Engineering.

DevOpsy type things is really about the only types of things that I have done in about two years, so the change really makes sense.

The Plaftorm Engineering organization has been doing some work in [Go](https://golang.org), so I thought that I would take a look at the language and see if it's something worth learning. To make that determination, I figured I would use golang to do something that I have been wanting to do for a while as a personal project: create a command line tool for 
this blog.

Why do I need a command line tool for a blog, you say? Well, I don't really. This blog has never been about what I should do or what I need to do. I mean there are a gazillion (real word) blog platforms and content management systems out there. Did that stop me from writing my own blogging platform from scratch? Nooooooo it didn't.

So I wrote a command line tool for this blog too. Just for the fun of it and to learn a little bit about golang.

A short overview of how this blog works:

1. I drop a hand coded HTML partial (a new post) into a certain directory and push to an AzureDevOps Git repo.
2. The commit to the Git master branch triggers the Node app build.
3. A Docker image is built that contains the new blog post.
4. The new Docker image for the blog back end is pushed to an Azure image registry.
5. The final step in the build reloads the container with the latest image on Azure.
6. A browser gets a request for the Angular.js blog route from the user.
7. My new post shows up on the blog page at the top.

I was driven to create this completely text-based blog, because I wanted complete control of my content, and I had grown tired of maintaing databases and database servers. Schema updates and migrations are for the birds. Yeah, I know. NoSQL. Sure but I didn't want to have to host that either. So, I came up with the idea to just push the new post to source control, and let the cloud-build take care of the rest. No SQL. No NoSQL. [No stinking badges](https://www.youtube.com/watch?v=Dln7yj8MDWE).

So then it became cumbersome to have to hand-code the HTML for the post. I developed a template generator script using Typescript, but I still had to hand-code the inner-HTML. I figured I would start coding my new posts in Markdown, and then have the blog engine convert the Markdown to HTML before sending it over to the browser.

That would have actually have been fairly trivial to implement in the Node backend. But why do it the easy way, when I could learn a whole new language and over-engineer the whole process?

So the [bamf-cli tool](https://github.com/tqheel/bamf), written in golang was born.

Is Go something worth learning? The jury's still out on that one. I learned a lot, but I could have done the same thing in using Node with Commander, or with C# in a lot less time.

But it was a fun diversion, and now I can write blog posts (like this one), in Markdown and then let Git take the converted file all the way to the cloud in under two minutes.
