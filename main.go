package main

import (
	"bytes"
	"fmt"
	"unicode"
)

const s = `
CrYPtogRapHY iS a ScIEnce Of “seCrET wriTinG”. FOr aT Least Two
THoUsAnd yeaRS ThErE haVE bEeN peOPlE WHo WAnTeD to
SEnd MESsaGes WHiCh coUlD oNLY bEeN reAd bY tHe pEOplE FOR
WhOm thEy were INtEndeD. a lOT oF different MEtHODs For
CONceaLiNG MESSAgeS weRE InvENted startINg WiTh ancient
cIPhERS lIke “skYtAle” AnD “AtbasH” AND EndINg WITh mOdERn
SYMmetRIc And puBLiC key eNCryptiON algOriThMS SuCH AS aES
AnD rSa. THe deVelOPMeNt oF CRyPTOgrAPhY coNtiNUeS aND
neveR SToPs! DecRYpT the MessAGe ThAt iS hIDDeN iN tHE TeXt of
This Task! thE alPhaBeT For thE mESsagE cOnsIsTS of alL TWENty
SIx ENGliSh letTErs fROM “a” tO “Z” and sIx puNCtUAtiON MARkS “
”, “.”, “,”, “!”, “?”, “’”.
`

// const s = `
// CrYPtogRapHY iS a ScIEnce Of «seCrET wriTinG». FOr aT
// Least Two THoUsANd yeaRS ThErE haVE bEeN peOPlE WHo WAnTeD
// to SEnd MESsaGes WHiCh coUlD oNly bEen rEAd bY tHe pEOPLe
// FoR whOm tHey were iNteNdeD. a loT oF different MEtHODs FoR
// coNcEalING mEssageS WerE invENtED stARTING WIth AnCIeNt
// cIPHerS lIKE «SkytaLE» and «ATBAsH» and ending wiTH MOdErn
// SymmeTRiC ANd PubliC-kEy enCRYptioN ALGOriTHmS SUch aS AeS
// and Rsa. the dEVELopMENT Of crYPtOgRaPHy cOntiNueS And
// NEVER sTopS! decrYPt THe mESsaGe tHat iS hIDdEn in thE teXT oF
// this TASk! tHE aLphabet FoR THE mEssAGE ConsisTs of ALl tWEnTy
// six enGliSh letTERS from «a» To «z» ANd Six puNCTuaTIoN MARkS
// « », «.», «,», «!», «?», «’».
// `

const letters = "abcdefghijklmnopqrstuvwxyz .,!?'"

func main() {
	bits := bitSlice(s)
	fmt.Printf("# of bits: %d\n\n", len(bits))
	fmt.Printf("bits: %v\n\n", bits)
	nums := bitsToNums(bits)
	fmt.Printf("numbers: %v\n\n", nums)
	msg := numsToS(nums)
	fmt.Printf("message: %s\n", msg)
}

func bitSlice(s string) []int {
	bits := []int{}
	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}
		if unicode.IsLower(c) {
			bits = append(bits, 0)
		} else {
			bits = append(bits, 1)
		}
	}
	return bits
}

func bitsToNums(bits []int) []int {
	nums := []int{}
	if len(bits)%5 != 0 {
		return nil
	}
	for i := 0; i < len(bits); i += 5 {
		n := 0
		for j := 0; j < 5; j++ {
			n += (1 << uint(4-j)) * bits[i+j]
		}
		nums = append(nums, n)
	}
	return nums
}

func numsToS(nums []int) string {
	b := &bytes.Buffer{}
	for _, code := range nums {
		fmt.Fprintf(b, "%s", string(letters[code]))
	}
	return b.String()
}
