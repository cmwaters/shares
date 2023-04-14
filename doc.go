/*
The shares package outlines a serialization protocol that encodes arbitrary sets
of data that belong to an identifier, known as a namespace, to evenly sized shares.

The protocol contains multiple versioned implementations that contain different
tradeoffs. New implementations can be created so long as they pick a unique version
and implement the Schema interface.
*/
package shares