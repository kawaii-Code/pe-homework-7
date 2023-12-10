package main;

import (
    "fmt"
)

func main() {
    Отвратительный код

    proof := `An Overview of the Proof of Fermat’s Last Theorem
Glenn Stevens
The principal aim of this article is to sketch the proof of the following
famous assertion.
Fermat’s Last Theorem. For n > 2, we have
FLT(n) : an + bn = cn
a, b, c 2 Z

=) abc = 0.
Many special cases of Fermat’s Last Theorem were proved from the
17th through the 19th centuries. The first known case is due to Fermat
himself, who proved FLT(4) around 1640. FLT(3) was proved by Euler
between 1758 and 1770. Since FLT(d) =) FLT(n) whenever d|n, the re-
sults of Euler and Fermat immediately reduce our theorem to the following
assertion.
Theorem. If p  5 is prime, and a, b, c 2 Z, then ap + bp + cp = 0 =)
abc = 0.
The proof of this theorem is the result of the combined e↵orts of innumer-
able mathematicians who have worked over the last century (and more!)
to develop a rich and powerful arithmetic theory of elliptic curves, modular
forms, and galois representations. It seems appropriate to emphasize the
names of five individuals who had the insight to see how this theory could
be used to prove Fermat’s Last Theorem and to supply the final crucial
ingredients of the proof:
Gerhart Frey (1985), who first suggested that the existence of a solu-
tion of the Fermat equation might contradict the Modularity Conjecture
of Taniyama, Shimura, and Weil;
Jean-Pierre Serre (1985-6), who formulated and (with J.-F. Mestre)
tested numerically a precise conjecture about modular forms and galois rep-
resentations mod p and who showed how a small piece of this conjecture—
the so-called epsilon conjecture—together with the Modularity Conjecture
would imply Fermat’s Last Theorem;
Ken Ribet (1986), who proved Serre’s epsilon conjecture, thus reducing
the proof of Fermat’s Last Theorem to a proof of the Modularity Conjecture
for semistable elliptic curves;
Richard Taylor (1994), who collaborated with Wiles to complete the
proof of Wiles’s numerical criterion in the minimal case;
Andrew Wiles (1994), who had the vision to identify the crucial numer-
ical criterion from which the Modularity Conjecture for semistable elliptic
curves would follow, and who finally supplied a proof of this criterion, thus
completing the proof of Fermat’s Last Theorem.
1
To prove the theorem we follow the program outlined by Serre in [16].
Fix a prime p  5 and suppose a, b, c 2 Z satisfy ap + bp + cp = 0 but
abc 6 = 0. The triple (ap, bp, cp) is what Gerhard Frey has called a “remark-
able” triple of integers, so remarkable in fact, that we suspect it does not
exist. To derive a contradiction, we will transform this triple into another
object with remarkable properties, namely a very special modular form
fap,bp,cp , something firmly rooted in the fertile grounds of modern number
theory. The construction of this modular form is a two-step process. First,
by a simple but insightful construction due independently to Yves Helle-
gouarch and Gerhard Frey, we obtain a certain semistable elliptic curve
Eap,bp,cp defined over Q. Then, by Wiles’s semistable modularity theorem,
we deduce the existence of a modular form fap,bp,cp associated to Eap,bp,cp
by the correspondence of Eichler and Shimura.
With fap,bp,cp in hand, we seek a contradiction within the realm of
modular forms. The crucial ingredients that finally lead to a contradiction
are encoded in a certain galois representation ⇢ap,bp,cp : G ! GL2(Fp)
mod p associated to fap,bp,cp . As noted by Frey and Serre, the remarkable-
ness of the triple (ap, bp, cp) is reflected by some remarkable local properties
of ⇢ap,bp,cp . Indeed, they noted that ⇢ap,bp,cp can ramify only at 2 and p,
and that the ramification at p is rather mild (what Serre called peu ram-
ifi´ee). But experience with galois representations shows that it is dicult
to make large galois representations with so little ramification. As Serre
conjectured and Ribet proved, the existence of such a galois representation
has untenable consequences in the theory of modular forms. Fermat’s Last
Theorem follows.
§1. A Remarkable Elliptic Curve
In this section we describe the crucial construction of an elliptic curve
Eap,bp,cp out of a hypothetical solution of the Fermat equation ap +bp +cp =
0. For any triple (A, B, C) of coprime integers satisfying A + B + C =
0, Gerhart Frey [8] considered the elliptic curve EA,B,C defined by the
Weierstrass equation
EA,B,C : y2 = x(x  A)(x + B)
and explained some of the ways in which the arithmetic properties of
EA,B,C are related to the diophantine properties of the triple (A, B, C).
Especially interesting are the connections with the Masser-Oesterle A-B-C
conjecture and its generalizations. For a discussion of this line of thought
including connections with modular curves, we refer the reader to [7] and
to Frey’s article in this volume (chapter XX).
For our purposes it suces to consider only the special case where
(A, B, C) = (ap, bp, cp) corresponds to a hypothetical solution of the Fermat
equation. Without loss of generality, we may assume a ⌘ 1 modulo 4 and
2|b. It is not hard to calculate both the minimal discriminant ap,bp,cp and
the conductor Nap,bp,cp of the elliptic curve Eap,bp,cp .
2
(1.1) Proposition. Let p  5 be prime and let a, b, c be coprime integers
satisfying abc 6 = 0, a ⌘ 1 modulo 4, 2|b, and ap + bp + cp = 0. Then
Eap,bp,cp is a semistable elliptic curve whose minimal discriminant and
conductor are given by the formulas
(a) ap,bp,cp = 28 · (abc)2p, and
(b) Nap,bp,cp = Q
'|abc '.
For definitions of semistability and of the conductor and minimal dis-
criminant see Silverman’s article in this volume (chapter II, especially §14
and §17). In general the primes dividing the minimal discriminant of an
elliptic curve over Q are the same as those dividing the conductor and this
might lead us to suspect that the discriminant and conductor should be
close to one another. Indeed, Szpiro has formulated the following conjec-
ture (see [19] where a slightly stronger form of the conjecture is formulated).
Conjecture. (Szpiro) For any ✏ > 0 there is a constant C > 0 such that
the minimal discriminant E and conductor NE of any elliptic curve E/Q
satisfy the inequality
|E | < C · N 6+✏
E .
On the other hand, proposition 1.1 shows that a counterexample to
F LT (p) for suciently large p gives rise to an elliptic curve whose mini-
mal discrimant and conductor are so far apart that they would contradict
Szpiro’s conjecture. We might thus hope to uncover a contradiction within
the field of diophantine geometry. We will follow a di↵erent but related
path and examine certain galois representations attached to Eap,bp,cp .
The idea of using elliptic curves to study Fermat’s Last Theorem and
vice versa goes back at least to the work of Y. Hellegouarch [9] (1972) who
studied connections between the Fermat equation and torsion points on
elliptic curves. Gerhart Frey seems to have been the first to suspect that a
counterexample to Fermat’s Last Theorem might contradict the Modularity
Conjecture and to investigate various approaches based on this idea.
§2. Galois Representations.
In this section we collect the basic definitions and conventions from the
theory of galois representations that we will need later. For more details
we refer the reader to the article by Mazur in this volume (chapter VIII).
Let Q be the algebraic closure of Q in C. We endow the galois group
GQ := Gal(Q/Q) with the Krull topology in which a basis of neighbor-
hoods of the origin is given by the collection of subgroups H ✓ GQ of finite
index in GQ. With this topology, GQ is a profinite group and in particular
is a compact topological group.
By a two dimensional galois representation over a topological ring A
we mean a continuous group homomorphism
⇢ : GQ ! GL2(A).
3
In this paper, the topological ring A will always be what Mazur calls a
coecient ring (in chapter VIII). Since these rings will play an important
role in what follows, we make a formal definition.
(2.1) Definition. A coecient ring is a complete noetherian local ring
with finite residue field of characteristic p (our fixed prime).
Whenever we write that ⇢ : GQ ! GL2(A) is a galois representation, it
is understood that A is a coecient ring and that ⇢ is continuous.
(2.2) Residual representations and deformations. Let A be a co-
ecient ring with maximal ideal mA and let kA := A/mA be the resid-
ual field. We define the residual representation of a galois representation
⇢ : GQ ! GL2(A) to be the representation
⇢ : GQ ! GL2(kA)
obtained by composing ⇢ with the reduction map GL2(A) ! GL2(kA).
Conversely, if ⇢0 : GQ ! GL2(k) is a two dimensional galois representa-
tion over a finite field k, then we say that ⇢ is a lifting of ⇢0 to A if k = kA
and ⇢ = ⇢0. Two liftings ⇢, ⇢0 of ⇢0 to A are said to be equivalent if ⇢0 can
be conjugated to ⇢ by a matrix in GL2(A) that is congruent to the identity
matrix modulo mA.
A deformation of ⇢0 to A is an equivalence class of liftings of ⇢0 to A.
For a given lifting ⇢ of ⇢0, we will abuse notation and also write ⇢ to denote
the deformation to which it belongs. This should not cause confusion in
our discussion.
(2.3) The determinant of a galois representation. If ⇢ is a two
dimensional galois representation over A then
det(⇢) : GQ ! A⇥
will denote the composition of ⇢ with the determinant homomorphism det :
GL2(A) ! A⇥. In the applications it is sometimes convenient to restrict
our attention to representations with prescribed determinant.
For example, let p : GQ ! Z⇥
p denote the cyclotomic character,
which is characterized by the property (⇣) = ⇣p() for any p-power root of
unity ⇣ and any  2 GQ. Any coecient ring A admits a unique continuous
ring homomorphism Zp ! A and we therefore have a canonical group
homomorphism Z⇥
p ! A⇥. We say that ⇢ has determinant p if det(⇢) is
the composition of p with the canonical homomorphism Z⇥
p ! A⇥.
(2.4) Local galois groups. For each prime ', we let Q' denote the field of
'-adic rationals, i.e. the completion of Q with respect to the '-adic absolute
value | · |'. We fix once and for all an algebraic closure Q' of Q' as well as
an embedding of Q into Q'. For ' = 1 we let Q1 := R, the completion
4
of Q with respect to the usual absolute value | · |1, and we take Q1 := C.
For each ' (' prime, or ' = 1), the local galois group at ' is the group
GQ' := Gal(Q'/Q').
For ' = 1, we have
GQ1 := Gal(C/R) = hci,
the cyclic group of order 2 generated by complex conjugation c. It is well-
known that for each ' there is a unique absolute value | · |' on Q' extending
the given absolute value on Q'. From this it follows easily that the elements
of GQ' are continuous automorphisms of Q'.
Using our fixed embeddings Q ✓ Q', we may restrict any automor-
phism of Q' to obtain an automorphism of Q. Since Q is dense in Q', the
induced homomorphisms GQ' ! GQ are injective and we will regard them
as inclusions:
GQ' ✓ GQ.
These subgroups are often called the decomposition subgroups of GQ. Of
course, strictly speaking, they are not well-defined since their definition
depends on our choice of the fixed embeddings of Q into Q'. However,
changing any one of these embeddings has the e↵ect of conjugating the
corresponding decomposition subgroup by an element of GQ. This ambi-
guity will not be important to us.
(2.5) Inertia groups. For ' 6 = 1, GQ' preserves the ring Z' of integers
in Q' and also preserves the maximal ideal  ✓ Z'. Thus, GQ' acts
naturally on the residual field F' = Z'/ and we obtain a natural map
GQ' ! Gal(F'/F'), which is easily seen to be surjective. Its kernel I'
is called the inertia group at '. Thus for each ' 6 = 1, we have an exact
sequence
1 ! I' ! GQ' ! Gal(F'/F') ! 1.
(2.6) Local properties of galois representations. Given a global galois
representation ⇢ : GQ ! GL2(A), we may restrict ⇢ to the decomposition
groups GQ' and obtain the family {⇢|G' } of local galois representations
⇢|G' : GQ' ! GL2(A).
In many important examples from number theory one knows that the global
representation ⇢ is determined up to isomorphism by the family of local
representations {⇢|G' }'62S , where ' ranges over the complement of any finite
set S of primes. By the local properties at ' of a galois representation ⇢
we mean the properties of the local representation ⇢|G' . The next three
definitions describe three local properties that play a special role in what
follows.
5
(2.7) Definition. We say that ⇢ is odd if det ⇢(c) = 1, where c is the
complex conjugation generating GQ1 .
(2.8) Definition. We say that ⇢ is unramified at a prime ' if I' ✓ ker ⇢|G' .
Since the galois group Gal(F'/F') is a topologically cyclic group gen-
erated by the 'th power Frobenius automorphism F rob', when ⇢ is unrami-
fied at ', ⇢|G' may be viewed as a homomorphism Gal(F'/F') ! GL2(A)
and is thus determined by its value on any representative of F rob' in GQ' .
When ' = p we need the following weaker condition.
(2.9) Definition. We say that ⇢ is flat at p if, for every ideal I ✓ A for
which A/I is finite, the representation GQp ! GL2(A/I), obtained by
reducing ⇢|GQp mod I, extends to a finite flat group scheme over Zp (see
Tate’s article in this volume (chapter V)).
(2.10) Examples from number theory. The Galois representations
that arise naturally in number theory have the especially nice property of
being unramified almost everywhere, that is, they are unramified at all but
finitely many primes '. For example, let E/Q be an elliptic curve. Then for
each n  0 the galois group GQ acts on the group E[pn] ⇠= (Z/pnZ)2 of pn-
torsion points on E. Since the action of GQ commutes with multiplication
by p on E, GQ acts naturally on the Tate module T ap(E) := lim
 E[pn] ⇠= Z2
p
and we obtain the p-adic galois representation
⇢E,p : GQ ! GL2(Zp)
associated to E. The residual representation ⇢E,p : GQ ! GL2(Fp).
describes the action of GQ on E[p] ⇠= F2
p. We have the following basic
result concerning the properties of these representations.
(2.11) Theorem. Let ⇢E,p be the p-adic galois representation associated
to an elliptic curve E/Q and let NE be the conductor of E. Then
• the determinant of ⇢E,p is p, and
• ⇢E,p is unramified outside of pNE .
In particular, ⇢E,p is odd. If E is semistable with minimal discriminant
E , then the residual representation ⇢E,p has the following local properties.
• If ' 6 = p, then ⇢E,p is unramified at ' () p|ord'(E ).
• ⇢E,p is flat at p () p|ordp(E ).
§3. A Remarkable Galois Representation.
Let E := Eap,bp,cp be as in §1 and consider the galois representation
⇢ap,bp,cp : GQ ! GL2(Fp)
given by ⇢ap,bp,cp = ⇢E,p. Gerhart Frey [7,8] and Jean-Pierre Serre [16]
noted that this representation has some remarkable local properties. More
precisely they proved the following theorem.
6
(3.1) Theorem Let p  5 be prime and a, b, c 2 Z satisfy ap + bp + cp = 0
and abc 6 = 0. Assume further that a ⌘ 1 modulo 4 and 2|b. Then
(a) ⇢ap,bp,cp is absolutely irreducible;
(b) ⇢ap,bp,cp is odd;
(c) ⇢ap,bp,cp is unramified outside 2p and is flat at p.
One suspects that there are no galois representations ⇢0 : GQ ! GL2(Fp)
satisfying properties (a), (b) and (c), but this suspicion remains unproven.
On the other hand, by a theorem of Ribet, we do know that no such galois
representation lives in the world of modular forms, in a sense that we will
make precise in the next section.
§4. Modular Galois Representations.
The theory of modular forms o↵ers a rich source of galois representa-
tions. Using the Hecke operators, these “modular” galois representations
can be constructed out of the torsion groups on the modular jacobians
J1(N ), N > 0 by the method of Eichler and Shimura. For an introduction
to the theory of modular forms and the Eichler-Shimura theory, see David
Rohrlich’s article in this volume (chapter III).
(4.1) Galois representations associated to newforms. Fix, once and
for all, a prime } of Q lying over p. Let f = P
n1 anqn be a weight two
(normalized) newform of conductor N and character ✏ (in (3.5) of chapter
III, newforms are called primitive forms). We let Kf be the completion at }
of the number field generated by the values of ✏ and the fourier coecents
an (n  1), and we let Of ✓ Kf be the ring of integers in Kf . The
theory of Eichler and Shimura associates to f an odd two dimensional
galois representation
⇢f : GQ ! GL2(Of )
such that for all suciently large primes ', ⇢f is unramified at ' and
Trace⇢f (F rob') = a' and det⇢f (F rob') = ✏(')'.
For the details of the Eichler-Shimura construction, we refer to section 3.7
of Rohrlich’s chapter III in this volume, where ⇢f appears as ⇢. By the
work of Carayol and others, we now have a good understanding of the local
structure of ⇢f at all primes. In particular we know that ⇢f is unramified
outside pN and that the above conditions on the trace and determinant of
⇢f (F rob') are satisfied for these primes.
By the work of Deligne [3] and Deligne-Serre [4], we know that similar
assertions hold for newforms of any weight w  1. Indeed, if f is a weight
w newform of conductor N then Deligne has constructed an odd two di-
mensional p-adic galois representation ⇢f , which is unramified outside pN
and satisfies Trace(⇢f (F rob')) = a' and det(⇢f (F rob')) = ✏(')'w1 for all
'6 |pN . In this paper, we will be concerned almost exclusively with the case
w = 2.
7
(4.2) Hecke algebras. Let N > 0 be an integer and let S2(N ) denote
the space of weight 2 cusp forms for 1(N ) (see (3.2) of chapter III). We
let
T0(N ) := Z[T', hdi] ✓ EndS2(N )
be the Z-subalgebra of End(S2(N )) generated by the Hecke operators T'
and the diamond operators hdi where ' runs over all primes not dividing
pN , and d runs over (Z/N Z)⇥ (see (3.3) of chapter III).
(4.3) Modularity of galois representations. Motivated by (4.1) we
say that a galois representation
⇢ : GQ ! GL2(A)
over a coecient ring A is modular if there exists an integer N > 0 and a
homomorphism ⇡ : T0(N ) ! A such that ⇢ is unramified outside N p and
for every prime ' 6 |pN we have
Trace(⇢(F rob')) = ⇡(T') and det(⇢(F rob')) = ⇡(h'i)'.
Remark: In view of the above restriction on the determinant it might
be more appropriate to call these modular representations of weight 2.
However, since all of our representations will have weight 2, we will drop
that modifier from our language.
(4.4) Serre’s Conjectures. In the special case where A = k is a finite
field, Serre [16] has formulated some precise conjectures about modularity
of galois representations over k. One consequence of Serre’s conjectures is
the following conjecture.
Conjecture. Every odd absolutely irreducible galois representation
⇢0 : GQ ! GL2(k)
is modular (in the sense of (4.3)).
In fact, Serre’s conjectures are much more precise. They predict—
in terms of the local structure of ⇢—the optimal weight, conductor and
character of a newform f for which ⇢f = ⇢0. For precise statements of
Serre’s conjectures and an account of what is known about them today,
see the article by Edixhoven in this volume (chapter VII). An important
special case of these conjectures, which Serre called the epsilon conjecture
in [16], is the following theorem of Ribet [13] (see §3 of chapter VII for a
sketch of the proof).
(4.5) Ribet’s Theorem. Let f be a weight two newform of conductor N '
where '6 |N is a prime. Suppose ⇢f is absolutely irreducible and that one of
the following is true:
• ⇢f is unramified at '; or
8
• ' = p and ⇢f is flat at p.
Then there is a weight two newform g of conductor N such that ⇢f ⇠= ⇢g .
§5. The Modularity Conjecture and Wiles’s Theorem.
We say that an elliptic curve E/Q is modular if there is a weight two
newform f of conductor NE and trivial character for which
L(f, s) = L(E, s).
There are a number of equivalent ways of defining modularity of elliptic
curves. Here are a few.
(5.1) Theorem. The following assertions are equivalent for an elliptic
curve E/Q.
(a) E is modular;
(b) for some prime p, ⇢E,p is modular;
(c) for every prime p, ⇢E,p is modular;
(d) there is a non-constant morphism ⇡ : X0(NE ) ! E of algebraic
curves defined over Q;
(e) E is isogenous to the modular abelian variety Af associated to some
weight two newform f of conductor NE .
We have the following profound conjecture developed between 1957
and 1967 by Shimura, Taniyama, and Weil.
(5.2) The Modularity Conjecture. Every elliptic curve over Q is mod-
ular.
The Modularity Conjecture is still open in general, but thanks to the work
of Wiles [20] and Taylor–Wiles [18], we know at least that it is true for a
large and important class of elliptic curves, namely the semistable ones.
(5.3) Wiles’s Theorem. Every semistable elliptic curve over Q is mod-
ular.
We will sketch the proof in §7. In fact, by improving Wiles’s methods,
Fred Diamond [5] has proven the much stronger result that every elliptic
curve E/Q that is semistable at 3 and 5 is modular. The proof is outlined
in chapter XVII by Diamond.
§6. The proof of Fermat’s Last Theorem.
Returning to the situation of §1 and §3 we suppose p  5 and assume
a, b, c 2 Z satisfy ap + bp + cp = 0 but abc 6 = 0. We derive a contradiction
by the method described in [16] (see also [8]). Without loss of generality,
we may assume a ⌘ 1 (mod 4) and 2|b. Let Eap,bp,cp be the elliptic curve
y2 = x(x  ap)(x + bp) and let ⇢ap,bp,cp be the associated p-adic galois
representation.
9
By proposition 1.1, Eap,bp,cp is semistable and has conductor Nap,bp,cp =Q
'|abc '. Hence, by Wiles’s theorem, Eap,bp,cp is modular and there is a
weight two newform fap,bp,cp of conductor Nap,bp,cp associated to Eap,bp,cp .
In particular, we have ⇢ap,bp,cp ⇠= ⇢fap,bp,cp . But according to theorem 2.11
⇢ap,bp,cp is absolutely irreducible and is unramified outside 2p and flat at p.
Applying Ribet’s Theorem we conclude that there is a weight two newform
g of conductor 2 such that ⇢g ⇠= ⇢ap,bp,cp . But the dimension of S2(0(2))
is equal to the genus of X0(2), which is easily seen to be zero. Thus there
are no weight two newforms of conductor 2. This is a contradiction and
Fermat’s Last Theorem is proved.
§7. The proof of Wiles’s Theorem.
In this final section, we describe the structure of the proof of Wiles’s
Theorem [18,20]. For other surveys of the proof, we recommend [2,12,14,17].
Here we assume that the distinguished prime p is  3. Let k be a finite
field of characteristic p and let
⇢0 : GQ ! GL2(k)
be a galois representation. As we move through this section we will impose
a number of cumulative hypotheses on ⇢0. The first of these is the following.
Hypothesis A. ⇢0 has determinant p.
(7.1) Semistable galois representations. We say that a galois repre-
sentation
⇢ : GQ ! GL2(A)
is ordinary at p if the restriction of ⇢ to the inertia group Ip at p has the
form ⇢|Ip =
✓ p ⇤
0 1
◆
for a suitable choice of basis. We say that ⇢ is
semistable at a prime ' if one of the following two conditions is satisfied.
• ' = p and ⇢ is either flat at p or ordinary at p (or both).
• ' 6 = p and ⇢|I' =
✓ 1 ⇤
0 1
◆
for a suitable choice of basis.
We say that a two dimensional galois representation ⇢ is semistable if it is
semistable at every prime. From now on, we impose the following additional
hypothesis on ⇢0.
Hypothesis B. ⇢0 is semistable.
The use of the word semistable in this context is motivated by the
simple fact that if E/Q is a semistable elliptic curve, then the p-adic galois
representation ⇢E,p : GQ ! GL2(Zp) is semistable in the above sense.
(7.2) Deformation types. A deformation type D is a list of conditions
to be imposed on deformations of ⇢0, satisfying certain properties. Using
more sophisticated terminology, a deformation type may be regarded as
10
a functor from the category of coecient rings to the category of sets,
where, for a given coecient ring A, D(A) is the set of two dimensional
galois representations over A that satisfy the conditions of D. For more
discussion of deformation types we refer the reader to Mazur’s chapter VIII
in this volume.
Wiles considers a variety of di↵erent deformation types, but for the
application to the semistable modularity conjecture it suces to restrict
to the following special cases. Let S := { ' 6 = p | ⇢0 is ramified at ' }. A
deformation type D is associated to a finite set of primes ⌃D disjoint from
S [ { p }. We say that a deformation ⇢ of ⇢0 is of type D if the following
conditions are satisfied.
• ⇢ has determinant p,
• ⇢ is unramified outside S [ { p } [ ⌃D,
• ⇢ is semistable outside ⌃D, and
• if p 6 2 ⌃D and if ⇢0 is flat at p, then ⇢ is also flat at p.
Roughly speaking, the last three conditions say that ⇢ has the same local
properties as ⇢0 at primes not in ⌃D. We remark that in any case, if ⇢0 is
ordinary at p then ⇢ is also ordinary at p.
(7.3) Universal deformation rings and Hecke rings. In addition to
hypotheses A and B above we suppose ⇢0 satisfies the following hypothesis.
Hypothesis C. ⇢0 is absolutely irreducible.
Using Mazur’s theory of deformations of galois representations [10], Wiles
associates to each deformation type D a universal deformation ring RD
(which is, in particular, a coecient ring) and a universal deformation
⇢D : GQ ! GL2(RD)
of ⇢0 of type D. The representation ⇢D,mod satisfies the following universal
property: for every deformation ⇢ : GQ ! GL2(A) of ⇢0 of type D there
is a unique homomorphism ⇡A : RD ! A such that the diagram
GQ
⇢D
! GL2(RD)
⇢ & . ⇡A
GL2(A)
is commutative. For details on the properties and construction of RD see
chapter VIII by Mazur and chapter XIII by Brian Conrad. An explicit
approach to constructing deformation rings is given in chapter IX by de
Smit, Rubin, and Schoof.
Hypothesis D. ⇢0 is modular, and ⇢0|GQ(p3) is absolutely irreducible.
Under this hypothesis, Wiles defines another coecient ring TD, the uni-
versal modular deformation ring and a universal modular deformation
⇢D,mod : GQ ! GL2(TD)
11
of ⇢0 of type D. The representation ⇢D,mod satisfies the analogous uni-
versal property for modular deformations of type D. Namely, for every
modular deformation ⇢ : GQ ! GL2(A) of ⇢0 of type D there is a unique
homomorphism ⇡A : TD ! A such that the obvious diagram commutes.
The constructions of TD and ⇢D,mod are quite dicult. The algebra
TD is defined in chapter XII by Diamond and Ribet. It’s existence depends
on the highly non-trivial fact (described in chapter VII by Edixhoven) that
⇢0 admits at least one modular deformation of type D. The representation
⇢D,mod is cut out of the Tate module of a modular Jacobian using the Hecke
operators. Wiles’s proof that this representation is a free rank two TD-
module depends on the Gorenstein property of TD (see Tilouine’s chapter X
in this volume). Later, other proofs of this fact were given that do not make
explicit use of the Gorenstein property, but rather have the Gorenstein
property as a by-product (for example, see [6]).
(7.4) The main theorem. By the universal property of ⇢D there is a
unique homomorphism 'D : RD ! TD such that ⇢D,mod = 'D  ⇢D. The
following theorem is a special case of the main theorem of Wiles [20].
Theorem. Suppose ⇢0 satisfies hypotheses A-D. Then the canonical map
'D : RD ! TD is an isomorphism of complete intersection rings.
For the definition of complete intersection rings, we refer to chapter IX by
Schoof, Rubin, and de Smit in this volume. For our purposes what matters
is the conclusion that 'D is an isomorphism. The proof of the theorem is
based on the numerical criterion of Wiles described in the next secion, which
reduces the proof to an inequality between two numbers. The theorem has
the following important corollary as an immediate consequence.
Corollary. Suppose ⇢0 satisfies hypotheses A-D. Then every deformation
of ⇢0 of type D is modular.
(7.5) Wiles’s numerical criterion. Let R and T be coecient rings and
suppose we have a commutative diagram
R '
! T
⇡R & . ⇡T
O
in which O is a complete discrete valuation ring and all the arrows are
surjective. Let IR := ker ⇡R, IT := ker ⇡T , and let ⌘T := ⇡T
AnnT (IT ).
Then the following three assertions are equivalent.
• ' is an isomorphism of complete intersection rings;
• IR/I2
R is finite and #(IR/I2
R)  #(O/⌘T );
• IR/I2
R is finite and #(IR/I2
R) = #(O/⌘T ).
12
This is a special case of Criterion I given in chapter IX by Schoof, Rubin,
and de Smit.
(7.6) Selmer groups and congruence modules. Now let f be a weight
two newform and suppose ⇢f : GQ ! GL2(Of ) is a deformation of ⇢0
of type D. By the universality of TD there is a unique homomorphism
⇡TD : TD ! Of such that ⇢f = ⇡TD  ⇢D,mod. Let ⇡RD := ⇡TD  'D so
that we have the following commutative diagram:
RD
'D
! TD
⇡RD & . ⇡TD
Of .
To prove that 'D is an isomorphism, Wiles establishes the middle inequality
in the above criterion. For this, he first interprets the two sides of the
inequality in terms of other objects that have been studied in some detail
in the literature. More precisely, Wiles interprets the “tangent space”
HomO(IRD /I2
RD , K/O) as a Selmer group H1
D(GQ, ad0(⇢f ) ⌦ K/O) (i.e.
as a certain subgroup of the galois cohomology group H1(GQ, ad0(⇢f ) ⌦
K/O) determined by local conditions associated to D), and he interprets
O/⌘TD as a congruence module classifying congruences between f and other
newforms of type D. For precise definitions, see chapter XII by Diamond
and Ribet, sections 4.2 and 4.3, chapter VIII by Mazur, and chapter IV by
Washington. The isomorphism between tangent spaces and Selmer groups
is described in chapter VIII.
The proof of the crucial numerical inequality divides into two parts.
The case where ⌃D = , which is called the minimal case, is proved by
Wiles with Taylor in [18]. Their original proof has been simplified by
making use of another criterion due to Faltings, a generalization of which
is given as criterion II in chapter XI. This is the method followed by de
Shalit in chapter XIV. The non-minimal case is proved by induction on the
number of primes in ⌃D. The proof is accomplished by analyzing how the
Selmer groups and congruence modules grow as ⌃D is enlarged to conclude
that if the numerical inequality is satisfied for one D then it is also satisfied
when more primes are included in ⌃D. See chapter XII by Diamond and
Ribet for more details.
(7.7) The Proof of Wiles’s Theorem. We prepare for the proof by
noting that hypotheses A and B are satisfied by ⇢E,p for every prime p.
Indeed hypothesis A is contained in theorem 2.11 and hypothesis B is a
consequence of the semistability of E.
Moreover, by a theorem of Serre ([15], prop. 21, and [17], §3.1), the
semistability of E guarantees that ⇢E,p is either surjective or reducible
for every prime p  3. Hence for p  3, absolute irreducibility of ⇢E,p is
equivalent to irreducibility of ⇢E,p, and if p = 3 this is equivalent to absolute
13
irreducibility of ⇢E,3|GQ(p3) . Thus the following lemma is a consequence
of corollary 7.4.
(7.8) Lemma. Let E/Q be a semistable elliptic curve and suppose ⇢E,p is
both modular and irreducible for some prime p  3. Then E is modular.
Wiles gave an ingenious argument to show that for E semistable, the
hypotheses of the lemma are satisfied by either p = 3 or p = 5. The proof
is based on the following three theorems.
(7.9) Theorem. Let E be an arbitrary elliptic curve and suppose ⇢E,3 is
irreducible. Then ⇢E,3 is modular.
This follows from a deep theorem of Langlands and Tunnell and de-
pends in a crucial way on the theory of Langlands for GL2. For an expo-
sition of the Langlands theory and the proof of Theorem 7.9, see chapter
VI by Stephen Gelbart in this volume.
(7.10) Theorem. Let E/Q be a semistable elliptic curve and suppose ⇢E,5
is irreducible. Then there is another semistable elliptic curve E0
/Q for which
(a) ⇢E0,3 is irreducible, and
(b) ⇢E0,5 ⇠= ⇢E,5.
Indeed, proposition 11 and the argument in section 4 of Rubin’s chap-
ter XVI in this volume provide us with a family of elliptic curves E0
/Q
satisfying conditions (a) and (b). All of these curves are semistable away
from 5. By taking E0 in this family suciently close 5-adically to E, we
obtain the desired semistable curve.
(7.11) Theorem. Let E/Q be a semistable elliptic curve. Then at least
one of the representations ⇢E,3 or ⇢E,5 is irreducible.
Indeed, if both ⇢E,3 and ⇢E,5 were reducible then E[15] would contain
a galois invariant subgroup of order 15. This contradicts Lemma 9 (iv) of
chapter XVI by Karl Rubin (see also [11]).
(7.12) Conclusion of the proof. Let E/Q be a semistable elliptic curve.
If ⇢E,3 is irreducible then, according to theorem 7.9, ⇢E,3 is also modular,
so E is modular by lemma 7.8. If ⇢E,3 is not irreducible, then by theorem
7.11, ⇢E,5 is irreducible. Then there is another semistable elliptic curve E0
/Q
satisfying (a) and (b) of theorem 7.10. In particular, ⇢E0,3 is irreducible.
Repeating the above argument we see that E0 is modular. Hence ⇢E0,5 is
modular and by (b) of 7.10, ⇢E,5 is modular. Once again we use lemma 7.8
to conclude E is modular.`

    fmt.Println(proof);
}
