package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOXact_GUC(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 float64
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 float64
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v800 int32
	_ = v800
	var v808 int32
	_ = v808
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v983 int32
	_ = v983
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1054 int32
	_ = v1054
	var v1069 int32
	_ = v1069
	var v1085 int32
	_ = v1085
	var v1101 int32
	_ = v1101
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[0]))
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = l0
	v35 = l1
	v48 = v32
	v49 = int32(_a_F_AtEOXact_GUC_0)
	goto L4
L2:
	;
	v1101 = l1
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[1])) = v1101 - int32(1)
	return
L4:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v66 = v48 - int32(20)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v67 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v1101 = v35
	goto L3
L6:
	;
	if v64 != 0 {
		v48 = v64
		v49 = v1085
		goto L4
	} else {
		goto L220
	}
L7:
	;
	v1085 = v48
	goto L6
L8:
	;
	goto L9
L9:
	;
	v71 = v48 - int32(76)
	v72 = int32(4)
	v73 = v48 + v72
	v75 = v48 - int32(48)
	v83 = v48 - v72
	v85 = v48 - int32(8)
	v87 = v48 - int32(44)
	v89 = v48 - int32(12)
	v91 = v48 - int32(16)
	v99 = v67
	v109 = v48
	goto L10
L10:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v124 < v35 {
		v1085 = v109
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v1085 = v1054
	goto L6
L12:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v34 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v1069 != 0 {
		v99 = v1069
		v109 = v1054
		goto L10
	} else {
		goto L219
	}
L14:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v99)+40))
	v832 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+40)) = v832
	if v831 == v832 {
		goto L175
	} else {
		goto L176
	}
L15:
	;
	F_pfree(m, v772)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L36
	} else {
		goto L174
	}
L16:
	;
	F_discard_stack_value(m, v71, v99+int32(32))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L36
	} else {
		goto L168
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v126
	F_pfree(m, v99)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L36
	} else {
		goto L167
	}
L18:
	;
	F_discard_stack_value(m, v71, v99+int32(32))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L36
	} else {
		goto L166
	}
L19:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v210+v99)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v99+v214)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v221 = *(*float64)(unsafe.Add(mBase, uint32(v212)))
	v223 = base.I32_wrap_i64(base.I64_reinterpret_f64(v221))
	v224 = int32(0)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v48-int32(52))))
	switch v225 {
	case 0:
		goto L52
	case 1:
		goto L51
	case 2:
		goto L50
	case 3:
		goto L49
	case 4:
		goto L48
	default:
		v808 = v224
		goto L14
	}
L20:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v210 = int32(24)
	v212 = v99 + int32(32)
	v213 = v205
	v214 = int32(16)
	v215 = v99 + int32(40)
	goto L19
L21:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	if v129 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v124 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v210 = int32(28)
	v212 = v99 + int32(48)
	v213 = int32(13)
	v214 = int32(20)
	v215 = v99 + int32(56)
	goto L19
L24:
	;
	switch v129 - int32(1) {
	case 0:
		goto L16
	default:
		goto L20
	case 2:
		goto L23
	}
L25:
	;
	goto L26
L26:
	;
	if v126 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v124 - int32(1)
	v1054 = v109
	goto L13
L28:
	;
	goto L29
L29:
	;
	v142 = v124 - int32(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v143 < v142 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v142
	v1054 = v109
	goto L13
L31:
	;
	goto L32
L32:
	;
	switch v129 - int32(1) {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L33
	default:
		goto L17
	}
L33:
	;
	F_discard_stack_value(m, v71, v99+int32(32))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L36
	} else {
		goto L43
	}
L34:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	if v161 != int32(1) {
		goto L18
	} else {
		goto L42
	}
L35:
	;
	F_discard_stack_value(m, v71, v99+int32(32))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	return
L37:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	if v152 == int32(3) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_discard_stack_value(m, v71, v126+int32(48))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L36
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = int32(1)
	goto L17
L41:
	;
	goto L40
L42:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+20)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v99)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+28)) = v166
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v99)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+48)) = v168
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v99)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+56)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = int32(3)
	goto L17
L43:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+20)) = v178
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+28)) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	if v182 == int32(3) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_discard_stack_value(m, v71, v126+int32(48))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L36
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v99)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+56)) = v189
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v99)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+48)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = int32(3)
	goto L17
L47:
	;
	goto L46
L48:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v693)))
	if v223 == v694 {
		goto L151
	} else {
		goto L152
	}
L49:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	if v400 != v223 {
		goto L99
	} else {
		goto L100
	}
L50:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v342)))
	if base.F64_eq(v221, v343) != 0 {
		goto L83
	} else {
		goto L84
	}
L51:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	if v223 == v286 {
		goto L68
	} else {
		goto L69
	}
L52:
	;
	v227 = v223 & int32(1)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	if v227 == v229 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v231 == v220 {
		v808 = v224
		goto L14
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v48)+32))
	if v233 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	m.T0[v233].(func(*base.Module, int32, int32))(m, v227, v220)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L36
	} else {
		goto L60
	}
L58:
	;
	v237 = v228
	goto L59
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v237))) = uint8(v227)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v220
	v241 = int32(1)
	if base.B2i32(v239 == int32(0))|base.B2i32(v239 == v220) != 0 {
		v808 = v241
		goto L14
	} else {
		goto L61
	}
L60:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v237 = v236
	goto L59
L61:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v239 == v246 {
		v808 = v241
		goto L14
	} else {
		goto L62
	}
L62:
	;
	v250 = v66
	goto L63
L63:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v278 == int32(0) {
		v772 = v239
		v776 = v241
		goto L15
	} else {
		goto L65
	}
L64:
	;
	v808 = v241
	goto L14
L65:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v278)+40))
	if v239 == v281 {
		v808 = v241
		goto L14
	} else {
		goto L66
	}
L66:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v278)+56))
	if v239 != v283 {
		v250 = v278
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v288 == v220 {
		v808 = v224
		goto L14
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	if v290 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	m.T0[v290].(func(*base.Module, int32, int32))(m, v223, v220)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L36
	} else {
		goto L75
	}
L73:
	;
	v294 = v285
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v223
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v220
	v298 = int32(1)
	if base.B2i32(v296 == int32(0))|base.B2i32(v296 == v220) != 0 {
		v808 = v298
		goto L14
	} else {
		goto L76
	}
L75:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v294 = v293
	goto L74
L76:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v296 == v303 {
		v808 = v298
		goto L14
	} else {
		goto L77
	}
L77:
	;
	v307 = v66
	goto L78
L78:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	if v335 == int32(0) {
		v772 = v296
		v776 = v298
		goto L15
	} else {
		goto L80
	}
L79:
	;
	v808 = v298
	goto L14
L80:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v335)+40))
	if v296 == v338 {
		v808 = v298
		goto L14
	} else {
		goto L81
	}
L81:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v335)+56))
	if v296 != v340 {
		v307 = v335
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v345 == v220 {
		v808 = v224
		goto L14
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v347 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	m.T0[v347].(func(*base.Module, float64, int32))(m, v221, v220)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L36
	} else {
		goto L90
	}
L88:
	;
	v351 = v342
	goto L89
L89:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v351))) = v221
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v220
	v355 = int32(1)
	if base.B2i32(v353 == int32(0))|base.B2i32(v353 == v220) != 0 {
		v808 = v355
		goto L14
	} else {
		goto L91
	}
L90:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v351 = v350
	goto L89
L91:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v353 == v360 {
		v808 = v355
		goto L14
	} else {
		goto L92
	}
L92:
	;
	v364 = v66
	goto L93
L93:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	if v392 == int32(0) {
		v772 = v353
		v776 = v355
		goto L15
	} else {
		goto L95
	}
L94:
	;
	v808 = v355
	goto L14
L95:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v392)+40))
	if v353 == v395 {
		v808 = v355
		goto L14
	} else {
		goto L96
	}
L96:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v392)+56))
	if v353 != v397 {
		v364 = v392
		goto L93
	} else {
		goto L97
	}
L97:
	;
	goto L94
L98:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v99)+32))
	v566 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+32)) = v566
	if v565 == v566 {
		goto L129
	} else {
		goto L130
	}
L99:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v48)+32))
	if v405 != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v402 != v220 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v542 = int32(0)
	goto L98
L102:
	;
	m.T0[v405].(func(*base.Module, int32, int32))(m, v223, v220)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L36
	} else {
		goto L105
	}
L103:
	;
	v410 = v399
	v411 = v400
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = v223
	if v411 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	v410 = v408
	v411 = v409
	goto L104
L106:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v220
	v491 = int32(1)
	if base.B2i32(v489 == int32(0))|base.B2i32(v489 == v220) != 0 {
		v542 = v491
		goto L98
	} else {
		goto L119
	}
L107:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	if v411 == v416 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	if v411 == v418 {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v411 == v420 {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v425 = v66
	goto L111
L111:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	if v452 != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	F_pfree(m, v411)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L36
	} else {
		goto L118
	}
L113:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+32))
	if v411 == v453 {
		goto L106
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	goto L112
L116:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v452)+48))
	if v411 != v455 {
		v425 = v452
		goto L111
	} else {
		goto L117
	}
L117:
	;
	goto L106
L118:
	;
	goto L106
L119:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v489 == v496 {
		v542 = v491
		goto L98
	} else {
		goto L120
	}
L120:
	;
	v501 = v66
	goto L121
L121:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	if v528 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	F_pfree(m, v489)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L36
	} else {
		goto L128
	}
L123:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)+40))
	if v489 == v529 {
		v542 = v491
		goto L98
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	goto L122
L126:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v528)+56))
	if v489 != v531 {
		v501 = v528
		goto L121
	} else {
		goto L127
	}
L127:
	;
	v542 = v491
	goto L98
L128:
	;
	v542 = v491
	goto L98
L129:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v99)+48))
	v645 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+48)) = v645
	if v644 == v645 {
		v808 = v542
		goto L14
	} else {
		goto L142
	}
L130:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	if v565 == v571 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	if v565 == v573 {
		goto L129
	} else {
		goto L132
	}
L132:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v565 == v575 {
		goto L129
	} else {
		goto L133
	}
L133:
	;
	v580 = v66
	goto L134
L134:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v580)))
	if v607 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	F_pfree(m, v565)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L36
	} else {
		goto L141
	}
L136:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)+32))
	if v565 == v608 {
		goto L129
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	goto L135
L139:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v607)+48))
	if v565 != v610 {
		v580 = v607
		goto L134
	} else {
		goto L140
	}
L140:
	;
	goto L129
L141:
	;
	goto L129
L142:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	if v644 == v650 {
		v808 = v542
		goto L14
	} else {
		goto L143
	}
L143:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	if v644 == v652 {
		v808 = v542
		goto L14
	} else {
		goto L144
	}
L144:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v644 == v654 {
		v808 = v542
		goto L14
	} else {
		goto L145
	}
L145:
	;
	v658 = v66
	goto L146
L146:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v658)))
	if v686 == int32(0) {
		v772 = v644
		v776 = v542
		goto L15
	} else {
		goto L148
	}
L147:
	;
	v808 = v542
	goto L14
L148:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v686)+32))
	if v644 == v689 {
		v808 = v542
		goto L14
	} else {
		goto L149
	}
L149:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v686)+48))
	if v644 != v691 {
		v658 = v686
		goto L146
	} else {
		goto L150
	}
L150:
	;
	goto L147
L151:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v696 == v220 {
		v808 = v224
		goto L14
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v48)+36))
	if v698 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	goto L153
L155:
	;
	m.T0[v698].(func(*base.Module, int32, int32))(m, v223, v220)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L36
	} else {
		goto L158
	}
L156:
	;
	v702 = v693
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v702))) = v223
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v220
	v706 = int32(1)
	if base.B2i32(v704 == int32(0))|base.B2i32(v704 == v220) != 0 {
		v808 = v706
		goto L14
	} else {
		goto L159
	}
L158:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v702 = v701
	goto L157
L159:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v704 == v711 {
		v808 = v706
		goto L14
	} else {
		goto L160
	}
L160:
	;
	v715 = v66
	goto L161
L161:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	if v743 == int32(0) {
		v772 = v704
		v776 = v706
		goto L15
	} else {
		goto L163
	}
L162:
	;
	v808 = v706
	goto L14
L163:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v743)+40))
	if v704 == v746 {
		v808 = v706
		goto L14
	} else {
		goto L164
	}
L164:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v743)+56))
	if v704 != v748 {
		v715 = v743
		goto L161
	} else {
		goto L165
	}
L165:
	;
	goto L162
L166:
	;
	goto L17
L167:
	;
	v1054 = v109
	goto L13
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v126
	if v126 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v64
	F_pfree(m, v99)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L36
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	F_pfree(m, v99)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L36
	} else {
		goto L173
	}
L172:
	;
	v1054 = v49
	goto L13
L173:
	;
	v1054 = v109
	goto L13
L174:
	;
	v808 = v776
	goto L14
L175:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v99)+56))
	v908 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+56)) = v908
	if v907 == v908 {
		goto L187
	} else {
		goto L188
	}
L176:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v831 == v836 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v831 == v838 {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	v843 = v66
	goto L179
L179:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v843)))
	if v870 != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	F_pfree(m, v831)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L36
	} else {
		goto L186
	}
L181:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v870)+40))
	if v831 == v871 {
		goto L175
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	goto L180
L184:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v870)+56))
	if v831 != v873 {
		v843 = v870
		goto L179
	} else {
		goto L185
	}
L185:
	;
	goto L175
L186:
	;
	goto L175
L187:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v983 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L188:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v907 == v912 {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v907 == v914 {
		goto L187
	} else {
		goto L190
	}
L190:
	;
	v919 = v66
	goto L191
L191:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v919)))
	if v946 != 0 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	F_pfree(m, v907)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L36
	} else {
		goto L198
	}
L193:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v946)+40))
	if v907 == v947 {
		goto L187
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	goto L192
L196:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v946)+56))
	if v907 != v949 {
		v919 = v946
		goto L191
	} else {
		goto L197
	}
L197:
	;
	goto L187
L198:
	;
	goto L187
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v48-int32(36)))) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v48-int32(28)))) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v126
	if v126 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L200:
	;
	if v213 == int32(0) {
		goto L199
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if v213 != 0 {
		goto L199
	} else {
		goto L208
	}
L203:
	;
	v989 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[2]))
	if v989 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v997 = int32(_a_F_AtEOXact_GUC_1)
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v997
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v996)+4)) = v85
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[3])) = v85
	goto L199
L205:
	;
	v991 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[3]))
	v996 = v991
	goto L204
L206:
	;
	goto L207
L207:
	;
	v993 = int32(_a_F_AtEOXact_GUC_1)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[2])) = v993
	v996 = v993
	goto L204
L208:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+4)) = v1004
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v1004))) = v1006
	goto L199
L209:
	;
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48-int32(56)))))
	if v1023&int32(64) == int32(0) {
		v1054 = v1022
		goto L13
	} else {
		goto L217
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v64
	F_pfree(m, v99)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L36
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	F_pfree(m, v99)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L36
	} else {
		goto L215
	}
L213:
	;
	if v808 != 0 {
		v1022 = v49
		goto L209
	} else {
		goto L214
	}
L214:
	;
	v1054 = v49
	goto L13
L215:
	;
	if v808 == int32(0) {
		v1054 = v109
		goto L13
	} else {
		goto L216
	}
L216:
	;
	v1022 = v109
	goto L209
L217:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v1028&int32(4) != 0 {
		v1054 = v1022
		goto L13
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v1028 | int32(4)
	v1034 = int32(_a_F_AtEOXact_GUC_2)
	v1035 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v1035
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[4])) = v73
	v1054 = v1022
	goto L13
L219:
	;
	goto L11
L220:
	;
	goto L5
}
func F_AtEOXact_MultiXact(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[0]))
	v3 = int32(_a_F_AtEOXact_MultiXact_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[1]))
	v5 = int32(2)
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2+v4<<(uint(v5)%32)))) = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[2]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v11+v13<<(uint(v5)%32)))) = v8
	v20 = int32(_a_F_AtEOXact_MultiXact_1)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[3])) = v20
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[4])) = v20
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[5])) = v8
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[6])) = v8
	return
}
