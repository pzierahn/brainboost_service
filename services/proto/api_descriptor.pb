
�
google/protobuf/empty.protogoogle.protobuf"
EmptyB}
com.google.protobufB
EmptyProtoPZ.google.golang.org/protobuf/types/known/emptypb��GPB�Google.Protobuf.WellKnownTypesJ�
 2
�
 2� Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


  

" E
	
" E

# ,
	
# ,

$ +
	
$ +

% "
	

% "

& !
	
$& !

' ;
	
%' ;

( 
	
( 
�
 2 � A generic empty message that you can re-use to avoid defining duplicated
 empty messages in your APIs. A typical example is to use it as the request
 or the response type of an API method. For instance:

     service Foo {
       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
     }




 2bproto3
�1
google/protobuf/timestamp.protogoogle.protobuf";
	Timestamp
seconds (Rseconds
nanos (RnanosB�
com.google.protobufBTimestampProtoPZ2google.golang.org/protobuf/types/known/timestamppb��GPB�Google.Protobuf.WellKnownTypesJ�/
 �
�
 2� Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


  

" 
	
" 

# I
	
# I

$ ,
	
$ ,

% /
	
% /

& "
	

& "

' !
	
$' !

( ;
	
%( ;
�
 � �� A Timestamp represents a point in time independent of any time zone or local
 calendar, encoded as a count of seconds and fractions of seconds at
 nanosecond resolution. The count is relative to an epoch at UTC midnight on
 January 1, 1970, in the proleptic Gregorian calendar which extends the
 Gregorian calendar backwards to year one.

 All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
 second table is needed for interpretation, using a [24-hour linear
 smear](https://developers.google.com/time/smear).

 The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
 restricting to that range, we ensure that we can convert to and from [RFC
 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.

 # Examples

 Example 1: Compute Timestamp from POSIX `time()`.

     Timestamp timestamp;
     timestamp.set_seconds(time(NULL));
     timestamp.set_nanos(0);

 Example 2: Compute Timestamp from POSIX `gettimeofday()`.

     struct timeval tv;
     gettimeofday(&tv, NULL);

     Timestamp timestamp;
     timestamp.set_seconds(tv.tv_sec);
     timestamp.set_nanos(tv.tv_usec * 1000);

 Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.

     FILETIME ft;
     GetSystemTimeAsFileTime(&ft);
     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;

     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
     Timestamp timestamp;
     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));

 Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.

     long millis = System.currentTimeMillis();

     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
         .setNanos((int) ((millis % 1000) * 1000000)).build();

 Example 5: Compute Timestamp from Java `Instant.now()`.

     Instant now = Instant.now();

     Timestamp timestamp =
         Timestamp.newBuilder().setSeconds(now.getEpochSecond())
             .setNanos(now.getNano()).build();

 Example 6: Compute Timestamp from current time in Python.

     timestamp = Timestamp()
     timestamp.GetCurrentTime()

 # JSON Mapping

 In JSON format, the Timestamp type is encoded as a string in the
 [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
 format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
 where {year} is always expressed using four digits while {month}, {day},
 {hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
 seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
 are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
 is required. A proto3 JSON serializer should always use UTC (as indicated by
 "Z") when printing the Timestamp type and a proto3 JSON parser should be
 able to accept both UTC and other timezones (as indicated by an offset).

 For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
 01:30 UTC on January 15, 2017.

 In JavaScript, one can convert a Date object to this format using the
 standard
 [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
 method. In Python, a standard `datetime.datetime` object can be converted
 to this format using
 [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with
 the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use
 the Joda Time's [`ISODateTimeFormat.dateTime()`](
 http://joda-time.sourceforge.net/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime()
 ) to obtain a formatter capable of generating timestamps in this format.



 �
�
  �� Represents seconds of UTC time since Unix epoch
 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
 9999-12-31T23:59:59Z inclusive.


  �

  �

  �
�
 �� Non-negative fractions of a second at nanosecond resolution. Negative
 second values with fractions must still have non-negative nanos values
 that count forward in time. Must be from 0 to 999,999,999
 inclusive.


 �

 �

 �bproto3
�
account_service.protochatbot.account.v1google/protobuf/empty.protogoogle/protobuf/timestamp.proto"�
Overview7
payments (2.chatbot.account.v1.PaymentRpayments4
usage (2.chatbot.account.v1.ModelUsageRusage
balance (Rbalance"�

ModelUsage
model (	Rmodel
input (Rinput
output (Routput
costs (Rcosts
requests (Rrequests"?
Usage6
models (2.chatbot.account.v1.ModelUsageRmodels"a
Payment
id (	Rid.
date (2.google.protobuf.TimestampRdate
amount (Ramount"=
Payments1
items (2.chatbot.account.v1.PaymentRitems2�
Account=
GetUsage.google.protobuf.Empty.chatbot.account.v1.UsageC
GetPayments.google.protobuf.Empty.chatbot.account.v1.PaymentsC
GetOverview.google.protobuf.Empty.chatbot.account.v1.OverviewB	Z./protoJ�
  *

  

 
	
 

 
	
  %
	
 )


 	 


 	

  
6

  


  
$

  
/4

 <

 

 '

 2:

 <

 

 '

 2:


  


 

   

  


  

  

  

  

 


 

 

 

 

 

 

 


 




 

 

 	

 





	







	







	







	




  




 !

 


 

 

  


" &


"

 #

 #

 #	

 #

$%

$

$ 

$#$

%

%

%	

%


( *


(

 )

 )


 )

 )

 )bproto3
�&
chat_service.protochatbot.chat.v1google/protobuf/empty.proto"
CollectionId
id (	Rid"�
CompletionRequest
document_id (	R
documentId
prompt (	RpromptB
model_options (2.chatbot.chat.v1.ModelOptionsRmodelOptions"4
CompletionResponse

completion (	R
completion"�
Prompt
	thread_id (	RthreadId#
collection_id (	RcollectionId
prompt (	RpromptB
model_options (2.chatbot.chat.v1.ModelOptionsRmodelOptionsN
retrieval_options (2!.chatbot.chat.v1.RetrievalOptionsRretrievalOptions 
attachments (	Rattachments"
ModelOptions
model_id (	RmodelId 
temperature (Rtemperature

max_tokens (R	maxTokens
top_p (RtopP"h
RetrievalOptions
enabled (Renabled
	threshold (R	threshold
	documents (R	documents"�
Source
document_id (	R
documentId
name (	Rname>
	fragments (2 .chatbot.chat.v1.Source.FragmentR	fragmentsf
Fragment
id (	Rid
content (	Rcontent
position (Rposition
score (Rscore"�
Message
	thread_id (	RthreadId
prompt (	Rprompt

completion (	R
completion1
sources (2.chatbot.chat.v1.SourceRsources"N
Thread
id (	Rid4
messages (2.chatbot.chat.v1.MessageRmessages"
ThreadID
id (	Rid"A
MessageIndex
	thread_id (	RthreadId
index (Rindex"
	ThreadIDs
ids (	Rids2�
Chat@
PostMessage.chatbot.chat.v1.Prompt.chatbot.chat.v1.Message?
	GetThread.chatbot.chat.v1.ThreadID.chatbot.chat.v1.ThreadJ
ListThreadIDs.chatbot.chat.v1.CollectionId.chatbot.chat.v1.ThreadIDsA
DeleteThread.chatbot.chat.v1.ThreadID.google.protobuf.EmptyP
DeleteMessageFromThread.chatbot.chat.v1.MessageIndex.google.protobuf.EmptyU

Completion".chatbot.chat.v1.CompletionRequest#.chatbot.chat.v1.CompletionResponseB	Z./protoJ�
  l

  

 
	
 

 
	
  %


  


 

  	,

  	

  	

  	#*

 
+

 


 


 
#)

 6

 

  

 +4

 =

 

 

 &;

 L

 

 *

 5J

 A

 

 "

 -?


  


 

  

  

  	

  


 




 

 

 	

 





	



!





 


 




 

 

 	

 


 1



/
 !" Thread ID to post the message to


 !

 !	

 !
3
$& Collection ID to post the message to


$

$	

$
,
' Prompt to generate completion


'

'	

'

*! Model options


*

*

* 

-) Search options


-

-$

-'(
(
0" Attachments to the prompt


0


0

0

0 !


3 8


3

 4

 4

 4	

 4

5

5

5

5

6

6

6	

6

7

7

7

7


: >


:

 ;

 ;

 ;

 ;

<

<

<

<

=

=

=	

=


@ L


@

 A

 A

 A	

 A

B

B

B	

B

 DI

 D


  E

  E


  E

  E

 F

 F


 F

 F

 G

 G


 G

 G

 H

 H	

 H


 H

K"

K


K

K

K !


N Z


N
'
 P Unique ID of the message


 P

 P	

 P
2
S% Prompt used to generate the message


S

S	

S
#
V Generated completion


V

V	

V
6
Y) Sources used to generate the completion


Y


Y

Y

Y


\ _


\

 ]

 ]

 ]	

 ]

^ 

^


^

^

^


	a c


	a

	 b

	 b

	 b	

	 b



e h



e


 f


 f


 f	


 f


g


g


g	


g


j l


j

 k

 k


 k

 k

 kbproto3
�
collection_service.protochatbot.collections.v1google/protobuf/empty.proto"0

Collection
id (	Rid
name (	Rname"J
CollectionList8
items (2".chatbot.collections.v1.CollectionRitems2�
CollectionsF
List.google.protobuf.Empty&.chatbot.collections.v1.CollectionListD
Insert".chatbot.collections.v1.Collection.google.protobuf.EmptyD
Update".chatbot.collections.v1.Collection.google.protobuf.EmptyD
Delete".chatbot.collections.v1.Collection.google.protobuf.EmptyB	Z./protoJ�
  

  

 
	
 

 
	
  %


  


 

  	;

  	


  	 

  	+9

 
9

 


 


 
"7

 9

 

 

 "7

 9

 

 

 "7


  


 

  

  

  	

  

 

 

 	

 


 




  

 


 

 

 bproto3
�
crashlytics.protocrashlytics.v1google/protobuf/empty.proto"g
Error
	exception (	R	exception
stack_trace (	R
stackTrace
app_version (	R
appVersion2K
Crashlytics<
RecordError.crashlytics.v1.Error.google.protobuf.EmptyB	Z./protoJ�
  

  

 
	
 

 
	
  %


  



 

  	9

  	

  	

  	"7


  


 

  

  

  	

  

 

 

 	

 

 

 

 	

 bproto3
�#
document_service.protochatbot.documents.v1google/protobuf/empty.protogoogle/protobuf/timestamp.proto"4
RenameDocument
id (	Rid
name (	Rname"A

DocumentID
id (	Rid#
collection_id (	RcollectionId"�
DocumentListC
items (2-.chatbot.documents.v1.DocumentList.ItemsEntryRitems`

ItemsEntry
key (	Rkey<
value (2&.chatbot.documents.v1.DocumentMetadataRvalue:8"z
SearchQuery
text (	Rtext#
collection_id (	RcollectionId
	threshold (R	threshold
limit (Rlimit"{
Chunk
id (	Rid
text (	Rtext
score (Rscore
postion (Rpostion

documentId (	R
documentId"�
SearchResults3
chunks (2.chatbot.documents.v1.ChunkRchunks]
document_names (26.chatbot.documents.v1.SearchResults.DocumentNamesEntryRdocumentNames@
DocumentNamesEntry
key (	Rkey
value (	Rvalue:8"C
IndexProgress
status (	Rstatus
progress (Rprogress"K
DocumentFilter
query (	Rquery#
collection_id (	RcollectionId"
DocumentMetadata0
file (2.chatbot.documents.v1.FileH Rfile1
web (2.chatbot.documents.v1.WebpageH RwebB
data"6
File
path (	Rpath
filename (	Rfilename"1
Webpage
url (	Rurl
title (	Rtitle"�
DocumentHeader
id (	Rid#
collection_id (	RcollectionId9

created_at (2.google.protobuf.TimestampR	createdAtB
metadata (2&.chatbot.documents.v1.DocumentMetadataRmetadata"�
IndexJob
id (	Rid#
collection_id (	RcollectionIdB
document (2&.chatbot.documents.v1.DocumentMetadataRdocument2�
DocumentP
List$.chatbot.documents.v1.DocumentFilter".chatbot.documents.v1.DocumentListF
Rename$.chatbot.documents.v1.RenameDocument.google.protobuf.EmptyB
Delete .chatbot.documents.v1.DocumentID.google.protobuf.EmptyN
Index.chatbot.documents.v1.IndexJob#.chatbot.documents.v1.IndexProgress0P
Search!.chatbot.documents.v1.SearchQuery#.chatbot.documents.v1.SearchResultsB	Z./protoJ�
  Z

  

 
	
 

 
	
  %
	
 )


 	 


 	

  
2

  



  


  
$0

 =

 

 

 &;

 9

 

 

 "7

 5

 

 

 %

 &3

 2

 

 

 #0


  


 

  

  

  	

  

 

 

 	

 


 




 

 

 	

 





	




 




 * Id to filename


 

  %

 ()


  %


 

 !

 !

 !	

 !

"

"

"	

"

#

#

#

#

$

$

$	

$


' -


'

 (

 (

 (	

 (

)

)

)	

)

*

*

*

*

+

+

+	

+

,

,

,	

,


/ 2


/

 0

 0


 0

 0

 0

1)

1

1$

1'(


4 7


4

 5

 5

 5	

 5

6

6

6

6


9 <


9

 :

 :

 :	

 :

;

;

;	

;


> C


>

 ?B

 ?

 @

 @

 @	

 @

A

A

A

A


	E H


	E

	 F

	 F

	 F	

	 F

	G

	G

	G	

	G



J M



J


 K


 K


 K	


 K


L


L


L	


L


O T


O

 P

 P

 P	

 P

Q

Q

Q	

Q

R+

R

R&

R)*

S 

S

S

S


V Z


V

 W

 W

 W	

 W

X

X

X	

X

Y 

Y

Y

Ybproto3
�
notion_service.protochatbot.notion.v2google/protobuf/empty.protochat_service.proto" 
NotionApiKey
key (	Rkey"�
NotionPrompt
database_id (	R
databaseId#
collection_id (	RcollectionId
prompt (	RpromptA
modelOptions (2.chatbot.chat.v1.ModelOptionsRmodelOptions"-
ExecutionResult
document (	Rdocument"
DatabasesID
id (	Rid"p
	Databases7
items (2!.chatbot.notion.v2.Databases.ItemRitems*
Item
id (	Rid
name (	Rname2�
NotionG
InsertAPIKey.chatbot.notion.v2.NotionApiKey.google.protobuf.EmptyG
UpdateAPIKey.chatbot.notion.v2.NotionApiKey.google.protobuf.Empty>
DeleteAPIKey.google.protobuf.Empty.google.protobuf.EmptyD
	GetAPIKey.google.protobuf.Empty.chatbot.notion.v2.NotionApiKeyE
ListDatabases.google.protobuf.Empty.chatbot.notion.v2.DatabasesV
ExecutePrompt.chatbot.notion.v2.NotionPrompt".chatbot.notion.v2.ExecutionResult0B	Z./protoJ�	
  -

  

 
	
 

 
	
  %
	
 


 
 


 


  A

  

  

  *?

 A

 

 

 *?

 J

 

 (

 3H

 >

 

 %

 0<

 ?

 

 )

 4=

 C

 

  

 +1

 2A


  


 

  

  

  	

  


 




 

 

 	

 





	







	



(



#

&'


  




 

 

 	

 


" $


"

 #

 #

 #	

 #


& -


&

 '*

 '


  (

  (


  (

  (

 )

 )


 )

 )

 ,

 ,


 ,

 ,

 ,bproto3