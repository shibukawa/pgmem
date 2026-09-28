package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_spgDeformLeafTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if l4 == int32(0) {
		F_index_deform_tuple_internal(m, l1, l2, l3, l0+int32(16), l0+int32(12), int32(base.Ui32(v7&int32(_a_F_spgDeformLeafTuple_0))>>(uint(int32(15))%32)))
		mBase = m.M
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v10 != int32(1) {
			F_index_deform_tuple_internal(m, l1, l2, l3, l0+int32(16), l0+int32(12), int32(base.Ui32(v7&int32(_a_F_spgDeformLeafTuple_0))>>(uint(int32(15))%32)))
			mBase = m.M
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
			v15 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v15)
			return
		}
	}
}
func F_spgWalk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int64
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int64
	_ = v326
	var v329 int64
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v459 int32
	_ = v459
	var v468 int32
	_ = v468
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v689 int64
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int64
	_ = v692
	var v693 int32
	_ = v693
	var v695 int64
	_ = v695
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v758 int32
	_ = v758
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v843 int32
	_ = v843
	var v855 int32
	_ = v855
	v5 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(96)
	m.G0 = v26
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)) = uint8(v5)
	v52 = v5
	goto L3
L1:
	;
	if v843 != 0 {
		goto L146
	} else {
		goto L147
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L6
	} else {
		goto L143
	}
L3:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v64 == int32(0) {
		v843 = v52
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L6
	} else {
		goto L140
	}
L5:
	;
	v67 = F_pairingheap_remove_first(m, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	if v67 == int32(0) {
		v843 = v52
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v72 = v67 + int32(36)
	v85 = v52
	goto L9
L9:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[0]))
	if v97 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L4
L11:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+43)))
	if v100 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L13
L15:
	;
	goto L10
L16:
	;
	v798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v268)+10)))
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)) = uint16(v798)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v268)+6))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v800
	v85 = v147
	goto L9
L17:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+43)))
	if v769 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L18:
	;
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v67)+16))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+42)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+44)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+45)))
	m.T0[l3].(func(*base.Module, int32, int32, int64, int32, int32, int32, int32, int32))(m, l1, v72, v103, v104, v105, v106, v107, v67+int32(48))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+38)))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+36)))
	v118 = v114 | v115<<(uint(int32(16))%32)
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+40)))
	if v85 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)) = uint8(v112)
	v758 = v85
	goto L17
L22:
	;
	if v147 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L23:
	;
	if v85 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	v142 = F_ReadBuffer(m, l0, v118)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L32
	}
L26:
	;
	if v138 == v118 {
		v147 = v85
		goto L22
	} else {
		goto L30
	}
L27:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[1]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v123+(v85^int32(-1))*int32(56))+16))
	v138 = v129
	goto L26
L28:
	;
	goto L29
L29:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[2]))
	v132 = int32(56)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131+v85*v132-v132)+16))
	v138 = v137
	goto L26
L30:
	;
	F_UnlockReleaseBuffer(m, v85)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	goto L25
L32:
	;
	F_LockBufferInternal(m, v142, int32(1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v147 = v142
	goto L22
L34:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+16)))
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166+v165))))
	v170 = v168 & int32(8)
	if v168&int32(4) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[3]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v151+(v147^int32(-1))<<(uint(int32(2))%32))))
	v165 = v157
	goto L34
L36:
	;
	goto L37
L37:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[4]))
	v165 = v159 + v147<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L38:
	;
	v173 = int32(1)
	if base.Ui32(v118-v173) <= base.Ui32(v173) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v165+v119<<(uint(int32(2))%32))+20))
	v268 = v165 + v265&int32(_a_F_spgWalk_0)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v271 = v269 & int32(3)
	if v271 == int32(1) {
		goto L16
	} else {
		goto L55
	}
L41:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+12)))
	if base.Ui32(v177) < base.Ui32(int32(25)) {
		v758 = v147
		goto L17
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v232 = v119
	goto L50
L44:
	;
	v186 = int32(base.Ui32(v177+int32(_a_F_spgWalk_1))>>(uint(int32(2))%32)) & int32(_a_F_spgWalk_2)
	if v186 == int32(0) {
		v758 = v147
		goto L17
	} else {
		goto L45
	}
L45:
	;
	v195 = int32(1)
	goto L46
L46:
	;
	v219 = F_spgTestLeafTuple(m, l1, v67, v165, v195&int32(_a_F_spgWalk_2), base.B2i32(v170 != int32(0)), int32(1), v26+int32(15), l3)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L48
	}
L47:
	;
	v758 = v147
	goto L17
L48:
	;
	v222 = v195 + int32(1)
	if base.Ui32(v222&int32(_a_F_spgWalk_2)) <= base.Ui32(v186) {
		v195 = v222
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v250 = v232 & int32(_a_F_spgWalk_2)
	if v250 == int32(0) {
		v758 = v147
		goto L17
	} else {
		goto L52
	}
L51:
	;
	v85 = v147
	goto L9
L52:
	;
	v253 = int32(0)
	v258 = F_spgTestLeafTuple(m, l1, v67, v165, v250, base.B2i32(v170 != v253), v253, v26+int32(15), l3)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	if v258 != int32(2049) {
		v232 = v258
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	if v271 != 0 {
		goto L15
	} else {
		goto L56
	}
L56:
	;
	v274 = int32(_a_F_spgWalk_3)
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_spgWalk[5]))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, _c_F_spgWalk[5])) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v280 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+88)) = v280
	*(*int64)(unsafe.Add(mBase, uint32(v26)+80)) = v280
	*(*int64)(unsafe.Add(mBase, uint32(v26)+72)) = v280
	v287 = int32(base.Ui32(v279) >> (uint(int32(3)) % 32))
	v289 = v287 & int32(_a_F_spgWalk_4)
	if v170 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	v514 = int32(0)
	if base.B2i32(base.B2i32(v511&int32(4) == v514)|base.B2i32(v510 == v514) == v514)&base.B2i32(v510 != v289) != 0 {
		goto L2
	} else {
		goto L80
	}
L58:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v294
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v296
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v298
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v67)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v300
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v304
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v306
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+208)))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+54)) = uint8(base.B2i32(base.Ui32(int32(_a_F_spgWalk_2)) < base.Ui32(v279)))
	v315 = int32(base.Ui32(v279)>>(uint(int32(2))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+53)) = uint8(v315)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+52)) = uint8(v308)
	if base.Ui32(v279) < base.Ui32(int32(_a_F_spgWalk_5)) {
		v329 = int64(0)
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v289
	v341 = F_palloc_mul(m, int32(4), v289)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L6
	} else {
		goto L68
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v289
	*(*int64)(unsafe.Add(mBase, uint32(v26)+56)) = v329
	v332 = F_spgExtractNodeLabels(m, l1, v268)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L6
	} else {
		goto L66
	}
L62:
	;
	v322 = v268 + int32(8)
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+50)))
	if v323 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v322)))
	v329 = v326
	goto L61
L64:
	;
	goto L65
L65:
	;
	v329 = base.I64_extend_i32_u(v322)
	goto L61
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v332
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v336 = F_FunctionCall2Coll(m, l1+int32(132), v335, base.I64_extend_i32_u(v26+int32(16)), base.I64_extend_i32_u(v26+int32(72)))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
	v510 = v338
	goto L57
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v341
	v344 = int32(0)
	if v289 == v344 {
		v510 = v344
		goto L57
	} else {
		goto L69
	}
L69:
	;
	v348 = v287 & int32(7)
	v349 = int32(0)
	if base.Ui32(int32(8)) <= base.Ui32(v289) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v360 = v349
	v367 = int32(0)
	goto L73
L71:
	;
	v436 = v349
	goto L72
L72:
	;
	v459 = v436
	v468 = v349
	goto L77
L73:
	;
	v379 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v341+v360<<(uint(v379)%32)))) = v360
	v384 = v360 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v341+v384<<(uint(v379)%32)))) = v384
	v390 = v360 | v379
	*(*int32)(unsafe.Add(mBase, uint32(v341+v390<<(uint(v379)%32)))) = v390
	v396 = v360 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v341+v396<<(uint(v379)%32)))) = v396
	v402 = v360 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v341+v402<<(uint(v379)%32)))) = v402
	v408 = v360 | int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v341+v408<<(uint(v379)%32)))) = v408
	v414 = v360 | int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v341+v414<<(uint(v379)%32)))) = v414
	v420 = v360 | int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v341+v420<<(uint(v379)%32)))) = v420
	v425 = int32(8)
	v426 = v360 + v425
	v428 = v367 + v425
	if v428 != v287&int32(_a_F_spgWalk_6) {
		v360 = v426
		v367 = v428
		goto L73
	} else {
		goto L75
	}
L74:
	;
	if v348 == int32(0) {
		v510 = v289
		goto L57
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	v436 = v426
	goto L72
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341+v459<<(uint(int32(2))%32)))) = v459
	v482 = int32(1)
	v485 = v468 + v482
	if v485 != v348 {
		v459 = v459 + v482
		v468 = v485
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v510 = v289
	goto L57
L79:
	;
	goto L78
L80:
	;
	if v510 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spgWalk[5])) = v275
	v758 = v147
	goto L17
L82:
	;
	v526 = F_palloc_mul(m, int32(4), v289)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if v528&int32(_a_F_spgWalk_7) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v541 = v268 + int32(base.Ui32(v528)>>(uint(int32(16))%32)) + int32(8)
	v544 = int32(0)
	goto L87
L85:
	;
	goto L86
L86:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	*(*int32)(unsafe.Add(mBase, _c_F_spgWalk[5])) = v600
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
	if v602 <= int32(0) {
		goto L81
	} else {
		goto L90
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526+v544<<(uint(int32(2))%32)))) = v541
	v564 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v541)+6)))
	v565 = int32(_a_F_spgWalk_4)
	v569 = v544 + int32(1)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if base.Ui32(v569) < base.Ui32(int32(base.Ui32(v570)>>(uint(int32(3))%32))&v565) {
		v541 = v541 + v564&v565
		v544 = v569
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L86
L89:
	;
	goto L88
L90:
	;
	v606 = int32(base.Ui32(v170) >> (uint(int32(3)) % 32))
	v612 = int32(0)
	v614 = v602
	goto L91
L91:
	;
	v631 = int32(2)
	v632 = v612 << (uint(v631) % 32)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v26)+76))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v632+v633)))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v526+v635<<(uint(v631)%32))))
	if v639 == int32(0) {
		v713 = v614
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L81
L93:
	;
	v719 = v612 + int32(1)
	if v719 < v713 {
		v612 = v719
		v614 = v713
		goto L91
	} else {
		goto L118
	}
L94:
	;
	v642 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v639)+4)))
	if v642 == int32(0) {
		v713 = v614
		goto L93
	} else {
		goto L95
	}
L95:
	;
	if v170 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v674 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v639)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v671)+40)) = uint16(v674)
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v639)))
	*(*int32)(unsafe.Add(mBase, uint32(v671)+36)) = v676
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v26)+80))
	if v679 != 0 {
		goto L107
	} else {
		goto L108
	}
L97:
	;
	v646 = F_palloc(m, int32(48))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L6
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
	if v649 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v646)+42)) = uint8(v606)
	v671 = v646
	goto L96
L101:
	;
	v651 = v649 + v632
	goto L103
L102:
	;
	v651 = l1 + int32(192)
	goto L103
L103:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v658 = F_palloc(m, v653<<(uint(int32(3))%32)+int32(48))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+42)) = uint8(v606)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if v661 <= int32(0) {
		v671 = v658
		goto L96
	} else {
		goto L105
	}
L105:
	;
	v665 = v661 << (uint(int32(3)) % 32)
	if v665 == int32(0) {
		v671 = v658
		goto L96
	} else {
		goto L106
	}
L106:
	;
	base.MemoryCopy(m, v658+int32(48), v652, v665)
	v671 = v658
	goto L96
L107:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v632+v679)))
	v683 = v681 + v678
	goto L109
L108:
	;
	v683 = v678
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v671)+32)) = v683
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v26)+84))
	if v685 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v689 = *(*int64)(unsafe.Add(mBase, uint32(v685+v612<<(uint(int32(3))%32))))
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	v691 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+36)))
	v692 = F_datumCopy(m, v689, v690, v691)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L6
	} else {
		goto L113
	}
L111:
	;
	v695 = int64(0)
	goto L112
L112:
	;
	v696 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v671)+24)) = v696
	*(*int64)(unsafe.Add(mBase, uint32(v671)+16)) = v695
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
	if v700 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v695 = v692
	goto L112
L114:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v632+v700)))
	v703 = v702
	goto L116
L115:
	;
	v703 = v696
	goto L116
L116:
	;
	v704 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v671)+45)) = uint8(v704)
	*(*uint16)(unsafe.Add(mBase, uint32(v671)+43)) = uint16(v704)
	*(*int32)(unsafe.Add(mBase, uint32(v671)+28)) = v703
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	F_pairingheap_add(m, v709, v671)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
	v713 = v712
	goto L93
L118:
	;
	goto L92
L119:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	if v782 != 0 {
		goto L128
	} else {
		goto L129
	}
L120:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v776 == int32(0) {
		goto L119
	} else {
		goto L126
	}
L121:
	;
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v772 == int32(0) {
		goto L120
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v775 != 0 {
		goto L119
	} else {
		goto L125
	}
L124:
	;
	goto L119
L125:
	;
	goto L120
L126:
	;
	F_pfree(m, v776)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	goto L119
L128:
	;
	F_pfree(m, v782)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L6
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
	if v785 != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	goto L130
L132:
	;
	F_pfree(m, v785)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L6
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	F_pfree(m, v67)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L6
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	F_MemoryContextReset(m, v790)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	if l2 != 0 {
		v52 = v758
		goto L3
	} else {
		goto L138
	}
L138:
	;
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+15)))
	if v793&int32(1) == int32(0) {
		v52 = v758
		goto L3
	} else {
		goto L139
	}
L139:
	;
	v843 = v758
	goto L1
L140:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v806 & int32(3)
	F_errmsg_internal(m, int32(_a_F_spgWalk_8), v26)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_spgWalk_9), int32(906), int32(_a_F_spgWalk_10))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errmsg_internal(m, int32(_a_F_spgWalk_11), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L6
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_spgWalk_9), int32(695), int32(_a_F_spgWalk_12))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_UnlockReleaseBuffer(m, v843)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L6
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	m.G0 = v26 + int32(96)
	return
L149:
	;
	goto L148
}
func F_spg_bbox_quad_config(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2)+12)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(603)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9783935500891)
	return int64(0)
}
func F_spg_box_quad_config(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2)+12)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9783935500891)
	return int64(0)
}
func F_spg_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = l0 - int32(16)
	if base.Ui32(v3) <= base.Ui32(int32(127)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v3)>>(uint(int32(2))%32))&int32(1073741820))+uint32(_c_F_spg_identify[0])))
		v12 = v10
	} else {
		v12 = int32(0)
	}
	return v12
}
