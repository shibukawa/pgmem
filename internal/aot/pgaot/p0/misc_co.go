package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_ConditionVariableBroadcast(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[0]))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[1]))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v71 != 0 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	F_s_lock(m, v11, int32(_a_F_ConditionVariableBroadcast_0), int32(238), int32(_a_F_ConditionVariableBroadcast_1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[0]))
	v27 = v22 + v24*int32(640)
	v29 = v27 + int32(84)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
	if v30 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v62
	*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[1])) = v62
	goto L3
L10:
	;
	if v47 == int32(-1) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22+v30*int32(640))+84)) = v41
	v46 = v30
	v47 = v41
	goto L10
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v30 != int32(-1) {
		v41 = v36
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v41 = v33
	goto L11
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v46 = v40
	v47 = v36
	goto L10
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(0)
	goto L9
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v46
	goto L17
L19:
	;
	goto L20
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(v53+v47*int32(640))+88)) = v46
	goto L17
L21:
	;
	F_s_lock(m, l0, int32(_a_F_ConditionVariableBroadcast_0), int32(317), int32(_a_F_ConditionVariableBroadcast_2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v79 == int32(-1) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return
L26:
	;
	goto L27
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v89 = v86 + v79*int32(640)
	v91 = v89 + int32(84)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+88))
	if v93 == int32(-1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v92 == int32(-1) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v92
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v102 = v97
	goto L28
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86+v93*int32(640))+84)) = v92
	v102 = v93
	goto L28
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = int64(0)
	v117 = v86 + v79*int32(640)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v118 != int32(-1) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v102
	goto L32
L34:
	;
	goto L35
L35:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v108+v92*int32(640))+88)) = v102
	goto L32
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v128 = v123 + v9*int32(640) + int32(84)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v129 == int32(-1) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	if v117 != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
	goto L38
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	goto L39
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v129
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v138+v129*int32(640))+84)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = int32(-1)
	goto L39
L43:
	;
	F_SetLatch(m, v117+int32(20))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if v118 != int32(-1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	goto L50
L48:
	;
	goto L49
L49:
	;
	return
L50:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v165 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	F_s_lock(m, l0, int32(_a_F_ConditionVariableBroadcast_0), int32(351), int32(_a_F_ConditionVariableBroadcast_2))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v173 == int32(-1) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L54
L56:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v216 = v215 + v9*int32(640)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+88))
	if v217 != 0 {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	v208 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v182 = v179 + v173*int32(640)
	v184 = v182 + int32(84)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)+88))
	if v186 == int32(-1) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v185 == int32(-1) {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v185
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v195 = v190
	goto L60
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179+v186*int32(640))+84)) = v185
	v195 = v186
	goto L60
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v184))) = int64(0)
	v208 = v182
	goto L56
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v195
	goto L64
L66:
	;
	goto L67
L67:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	*(*int32)(unsafe.Add(mBase, uint32(v201+v185*int32(640))+88)) = v195
	goto L64
L68:
	;
	v222 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v222
	if v208 == v222 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v221 = int32(1)
	goto L68
L70:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)+84))
	if v218 != 0 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v221 = int32(0)
	goto L68
L72:
	;
	if v221 != 0 {
		goto L50
	} else {
		goto L76
	}
L73:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[3]))
	if v208 == v227 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	F_SetLatch(m, v208+int32(20))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	goto L51
}
func F_ConditionVariablePrepareToSleep(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[0]))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1]))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
		if v12 != 0 {
			F_s_lock(m, v11, int32(_a_F_ConditionVariablePrepareToSleep_0), int32(238), int32(_a_F_ConditionVariablePrepareToSleep_1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[0]))
				v27 = v22 + v24*int32(640)
				v29 = v27 + int32(84)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
				if v30 == int32(0) {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					if v33 == int32(0) {
					} else {
						v41 = v33
						*(*int32)(unsafe.Add(mBase, uint32(v22+v30*int32(640))+84)) = v41
						v46 = v30
						v47 = v41
						if v47 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v46
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							*(*int32)(unsafe.Add(mBase, uint32(v53+v47*int32(640))+88)) = v46
						}
						*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(0)
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					if v30 != int32(-1) {
						v41 = v36
						*(*int32)(unsafe.Add(mBase, uint32(v22+v30*int32(640))+84)) = v41
						v46 = v30
						v47 = v41
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v36
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
						v46 = v40
						v47 = v36
					}
					if v47 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v46
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
						*(*int32)(unsafe.Add(mBase, uint32(v53+v47*int32(640))+88)) = v46
					}
					*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
				*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1])) = l0
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
				if v70 != 0 {
					F_s_lock(m, l0, int32(_a_F_ConditionVariablePrepareToSleep_0), int32(75), int32(_a_F_ConditionVariablePrepareToSleep_2))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
						v85 = v80 + v9*int32(640) + int32(84)
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v86 == int32(-1) {
							*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
							v94 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
							*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
							*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
						return
					}
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
					v85 = v80 + v9*int32(640) + int32(84)
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v86 == int32(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
						v94 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
						*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
						*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
					return
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[0]))
			v27 = v22 + v24*int32(640)
			v29 = v27 + int32(84)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
			if v30 == int32(0) {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				if v33 == int32(0) {
				} else {
					v41 = v33
					*(*int32)(unsafe.Add(mBase, uint32(v22+v30*int32(640))+84)) = v41
					v46 = v30
					v47 = v41
					if v47 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v46
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
						*(*int32)(unsafe.Add(mBase, uint32(v53+v47*int32(640))+88)) = v46
					}
					*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(0)
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				if v30 != int32(-1) {
					v41 = v36
					*(*int32)(unsafe.Add(mBase, uint32(v22+v30*int32(640))+84)) = v41
					v46 = v30
					v47 = v41
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v36
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					v46 = v40
					v47 = v36
				}
				if v47 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v46
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
					*(*int32)(unsafe.Add(mBase, uint32(v53+v47*int32(640))+88)) = v46
				}
				*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1])) = l0
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
			if v70 != 0 {
				F_s_lock(m, l0, int32(_a_F_ConditionVariablePrepareToSleep_0), int32(75), int32(_a_F_ConditionVariablePrepareToSleep_2))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
					v85 = v80 + v9*int32(640) + int32(84)
					v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v86 == int32(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
						v94 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
						*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
						*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
					return
				}
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
				v85 = v80 + v9*int32(640) + int32(84)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v86 == int32(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
					v94 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
					*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
					*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1])) = l0
		v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
		if v70 != 0 {
			F_s_lock(m, l0, int32(_a_F_ConditionVariablePrepareToSleep_0), int32(75), int32(_a_F_ConditionVariablePrepareToSleep_2))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
				v85 = v80 + v9*int32(640) + int32(84)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v86 == int32(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
					v94 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
					*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
					*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
				return
			}
		} else {
			v79 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
			v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
			v85 = v80 + v9*int32(640) + int32(84)
			v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v86 == int32(-1) {
				*(*int64)(unsafe.Add(mBase, uint32(v85))) = int64(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v86
				v94 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
				*(*int32)(unsafe.Add(mBase, uint32(v95+v86*int32(640))+84)) = v9
				*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(-1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
			return
		}
	}
}
func F_ConditionalLockBufferForCleanup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 < v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return v176
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(l0^int32(-1))<<(uint(int32(2))%32))))
	v176 = base.B2i32(v19 == int32(1))
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[1]))
	if l0 == v24 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v72 != int32(1) {
		v176 = v2
		goto L1
	} else {
		goto L34
	}
L6:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[2]))
	if l0 == v28 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_1)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[3]))
	if l0 == v32 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_2)
	goto L5
L13:
	;
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[4]))
	if l0 == v36 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_3)
	goto L5
L16:
	;
	goto L17
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[5]))
	if l0 == v40 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_4)
	goto L5
L19:
	;
	goto L20
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[6]))
	if l0 == v44 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_5)
	goto L5
L22:
	;
	goto L23
L23:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[7]))
	if l0 == v48 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_6)
	goto L5
L25:
	;
	goto L26
L26:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[8]))
	if l0 == v52 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v71 = int32(_a_F_ConditionalLockBufferForCleanup_7)
	goto L5
L28:
	;
	goto L29
L29:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[9]))
	if v56 == int32(0) {
		v176 = v2
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[10]))
	v63 = int32(0)
	v65 = F_hash_search(m, v60, v8+int32(8), v63, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	if v65 == int32(0) {
		v176 = v2
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v71 = v65
	goto L5
L34:
	;
	v76 = l0 << (uint(int32(6)) % 32)
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[11]))
	v83 = F_LWLockConditionalAcquire(m, v76+v78-int32(16), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	if v83 == int32(0) {
		v176 = v2
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(_a_F_ConditionalLockBufferForCleanup_8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(_a_F_ConditionalLockBufferForCleanup_9)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_ConditionalLockBufferForCleanup_10)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v101 = v76 + v88 - int32(40)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = int32(_a_F_ConditionalLockBufferForCleanup_11)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v102 | v103
	if v102&v103 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	goto L40
L38:
	;
	v124 = v102
	goto L39
L39:
	;
	v132 = int32(_a_F_ConditionalLockBufferForCleanup_12)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[12]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v135 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L31
	} else {
		goto L42
	}
L41:
	;
	v124 = v117
	goto L39
L42:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v118 = int32(_a_F_ConditionalLockBufferForCleanup_11)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v117 | v118
	if v117&v118 != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if v124&int32(_a_F_ConditionalLockBufferForCleanup_13) == int32(1) {
		goto L55
	} else {
		goto L56
	}
L45:
	;
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[12])) = v150
	goto L45
L47:
	;
	if int32(999) < v133 {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v133 < int32(11) {
		goto L45
	} else {
		goto L54
	}
L50:
	;
	v140 = int32(900)
	if v140 <= v133 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v143 = v140
	goto L53
L52:
	;
	v143 = v133
	goto L53
L53:
	;
	v150 = v143 + int32(100)
	goto L46
L54:
	;
	v150 = v133 - int32(1)
	goto L46
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v124 & int32(-4456447)
	v176 = int32(1)
	goto L1
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v124 & int32(-4194305)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[11]))
	F_LWLockRelease(m, v164+l0<<(uint(int32(6))%32)-int32(16))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L31
	} else {
		goto L58
	}
L58:
	;
	v176 = int32(0)
	goto L1
}
func F_CountChildren(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CountChildren[0]))
	if v12 == v2 {
		v101 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v101
L2:
	;
	if v12 == int32(_a_F_CountChildren_0) {
		v101 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = v12
	v28 = v2
	goto L4
L4:
	;
	if int32(base.Ui32(l0&int32(64))>>(uint(int32(6))%32))^int32(base.Ui32(l0&int32(2))>>(uint(int32(1))%32)) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v101 = v94
	goto L1
L6:
	;
	v56 = v27 - int32(12)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if int32(base.Ui32(l0)>>(uint(v57)%32))&int32(1) != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v35 = v27 - int32(12)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 != int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v27-int32(16))))
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_CountChildren[1]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v41<<(uint(int32(2))%32))+44))
	goto L9
L9:
	;
	if base.B2i32(v47 == int32(3)) == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(6)
	goto L6
L11:
	;
	v63 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v94 = v28
	goto L13
L13:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v96 != int32(_a_F_CountChildren_0) {
		v27 = v96
		v28 = v94
		goto L4
	} else {
		goto L25
	}
L14:
	;
	return int32(0)
L15:
	;
	if v63 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if base.Ui32(v67) <= base.Ui32(int32(17)) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	v94 = v28 + int32(1)
	goto L13
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v27-int32(20))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v77
	F_errmsg_internal(m, int32(_a_F_CountChildren_1), v9)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L14
	} else {
		goto L23
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67<<(uint(int32(2))%32))+uint32(_c_F_CountChildren[2])))
	v77 = v76
	goto L22
L21:
	;
	v77 = int32(_a_F_CountChildren_2)
	goto L22
L22:
	;
	goto L19
L23:
	;
	F_errfinish(m, int32(_a_F_CountChildren_3), int32(3942), int32(_a_F_CountChildren_4))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	goto L5
}
func F_CountDBBackends(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[1]))
	v18 = F_LWLockAcquire(m, v14+int32(512), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v22 <= int32(0) {
			v107 = v2
		} else {
			v25 = int32(1)
			v28 = v12 + int32(36)
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[2]))
			v31 = int32(0)
			if v22 != v25 {
				v38 = v31
				v39 = v2
				v40 = int32(0)
				for {
					v49 = v28 + v38<<(uint(int32(2))%32)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					v53 = v30 + v50*int32(640)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
					if v54 == int32(0) {
						v61 = v39
					} else {
						if l0 != 0 {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+60))
							if v57 != l0 {
								v61 = v39
							} else {
								v61 = v39 + int32(1)
							}
						} else {
							v61 = v39 + int32(1)
						}
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
					v65 = v30 + v62*int32(640)
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+44))
					if v66 == int32(0) {
						v73 = v61
					} else {
						if l0 != 0 {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+60))
							if v69 != l0 {
								v73 = v61
							} else {
								v73 = v61 + int32(1)
							}
						} else {
							v73 = v61 + int32(1)
						}
					}
					v74 = int32(2)
					v75 = v38 + v74
					v77 = v40 + v74
					if v77 != v22&int32(2147483646) {
						v38 = v75
						v39 = v73
						v40 = v77
						continue
					} else {
						break
					}
					break
				}
				v80 = v75
				v81 = v73
			} else {
				v80 = v31
				v81 = v2
			}
			if v22&v25 == int32(0) {
				v107 = v81
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v28+v80<<(uint(int32(2))%32))))
				v97 = v30 + v94*int32(640)
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+44))
				if v98 == int32(0) {
					v107 = v81
				} else {
					if l0 != 0 {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+60))
						if v101 != l0 {
							v107 = v81
						} else {
							v107 = v81 + int32(1)
						}
					} else {
						v107 = v81 + int32(1)
					}
				}
			}
		}
		v116 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[1]))
		F_LWLockRelease(m, v116+int32(512))
		mBase = m.M
		v120 = m.ExcPending
		if v120 != 0 {
			return int32(0)
		} else {
			return v107
		}
	}
}
func F_codepoint_range_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v6) <= base.Ui32(v5) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v10 = base.B2i32(base.Ui32(v8) < base.Ui32(v5))
	} else {
		v10 = int32(-1)
	}
	return v10
}
func F_colname_is_unique(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v8 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v250
L2:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v187 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v11 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v174 = int32(0)
	v176 = F_hash_search(m, v8, l0, v174, v174)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L59
	} else {
		goto L60
	}
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v19 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if int32(0) < v66 {
		goto L23
	} else {
		goto L24
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14+v19<<(uint(int32(2))%32))))
	if v26 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v57 = v19 + int32(1)
	if v57 != v11 {
		v19 = v57
		goto L9
	} else {
		goto L22
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v32 == int32(0) {
		v51 = v31
		v52 = v32
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v52-v51 != 0 {
		goto L11
	} else {
		goto L21
	}
L14:
	;
	goto L13
L15:
	;
	if v31 != v32 {
		v51 = v31
		v52 = v32
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v36 = v26
	v37 = l0
	goto L17
L17:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v41 == int32(0) {
		v51 = v40
		v52 = v41
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v51 = v40
	v52 = v41
	goto L14
L19:
	;
	v44 = int32(1)
	if v40 == v41 {
		v36 = v36 + v44
		v37 = v37 + v44
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	return int32(0)
L22:
	;
	goto L10
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v74 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v121 == int32(0) {
		goto L2
	} else {
		goto L40
	}
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v69+v74<<(uint(int32(2))%32))))
	if v81 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v112 = v74 + int32(1)
	if v112 != v66 {
		v74 = v112
		goto L26
	} else {
		goto L39
	}
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v87 == int32(0) {
		v106 = v86
		v107 = v87
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v107-v106 != 0 {
		goto L28
	} else {
		goto L38
	}
L31:
	;
	goto L30
L32:
	;
	if v86 != v87 {
		v106 = v86
		v107 = v87
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v91 = v81
	v92 = l0
	goto L34
L34:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if v96 == int32(0) {
		v106 = v95
		v107 = v96
		goto L31
	} else {
		goto L36
	}
L35:
	;
	v106 = v95
	v107 = v96
	goto L31
L36:
	;
	v99 = int32(1)
	if v95 == v96 {
		v91 = v91 + v99
		v92 = v92 + v99
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	return int32(0)
L39:
	;
	goto L27
L40:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v124 <= int32(0) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v127 = int32(0)
	if v127 < v124 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v131 = v124
	goto L44
L43:
	;
	v131 = v127
	goto L44
L44:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v136 = v127
	goto L45
L45:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v132+v136<<(uint(int32(2))%32))))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v147 == int32(0) {
		v166 = v146
		v167 = v147
		goto L48
	} else {
		goto L49
	}
L46:
	;
	return int32(0)
L47:
	;
	if v167-v166 != 0 {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	goto L47
L49:
	;
	if v146 != v147 {
		v166 = v146
		v167 = v147
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v151 = v143
	v152 = l0
	goto L51
L51:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+1)))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	if v156 == int32(0) {
		v166 = v155
		v167 = v156
		goto L48
	} else {
		goto L53
	}
L52:
	;
	v166 = v155
	v167 = v156
	goto L48
L53:
	;
	v159 = int32(1)
	if v155 == v156 {
		v151 = v151 + v159
		v152 = v152 + v159
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v170 = v136 + int32(1)
	if v131 != v170 {
		v136 = v170
		goto L45
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	goto L46
L58:
	;
	goto L2
L59:
	;
	return int32(0)
L60:
	;
	if v176 != 0 {
		v250 = v4
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L2
L62:
	;
	return int32(1)
L63:
	;
	goto L64
L64:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v193 <= int32(0) {
		v250 = int32(1)
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v196 = int32(0)
	if v196 < v193 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v199 = v193
	goto L68
L67:
	;
	v199 = v196
	goto L68
L68:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v205 = int32(0)
	goto L69
L69:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v200+v205<<(uint(int32(2))%32))))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if v216 == int32(0) {
		v235 = v215
		v236 = v216
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v250 = v239
	goto L1
L71:
	;
	v238 = int32(0)
	v239 = base.B2i32(v237 != v238)
	if v237 == v238 {
		v250 = v239
		goto L1
	} else {
		goto L79
	}
L72:
	;
	v237 = v236 - v235
	goto L71
L73:
	;
	if v215 != v216 {
		v235 = v215
		v236 = v216
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v220 = v212
	v221 = l0
	goto L75
L75:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+1)))
	if v225 == int32(0) {
		v235 = v224
		v236 = v225
		goto L72
	} else {
		goto L77
	}
L76:
	;
	v235 = v224
	v236 = v225
	goto L72
L77:
	;
	v228 = int32(1)
	if v224 == v225 {
		v220 = v220 + v228
		v221 = v221 + v228
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v243 = v205 + int32(1)
	if v243 != v199 {
		v205 = v243
		goto L69
	} else {
		goto L80
	}
L80:
	;
	goto L70
}
func F_comp_trgm(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_comp_trgm[0]))
	v5 = m.T0[v4].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_compact_trigram(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	v9 = int32(255)
	switch l2 {
	case 0:
		v61 = l2
		*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v61)
		v69 = int32(base.Ui32(v61) >> (uint(int32(16)) % 32))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v69)
		return
	default:
		v20 = l1
		v21 = l2
		v23 = v9
		v24 = v9
		v25 = v9
		v26 = v9
		for {
			v27 = int32(8)
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			v35 = *(*int32)(unsafe.Add(mBase, uint32((v29^v23)<<(uint(int32(2))%32))+uint32(_c_F_compact_trigram[0])))
			v38 = int32(16)
			v43 = int32(24)
			v46 = v35 ^ (v24<<(uint(v27)%32)&int32(_a_F_compact_trigram_0) | v25<<(uint(v38)%32)&int32(16711680) | v26<<(uint(v43)%32))
			v53 = int32(1)
			v56 = v21 - v53
			if v56 != 0 {
				v20 = v20 + v53
				v21 = v56
				v23 = int32(base.Ui32(v46) >> (uint(v43) % 32))
				v24 = v35
				v25 = int32(base.Ui32(v46) >> (uint(v27) % 32))
				v26 = int32(base.Ui32(v46) >> (uint(v38) % 32))
				continue
			} else {
				break
			}
			break
		}
		v61 = v46 ^ int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v61)
		v69 = int32(base.Ui32(v61) >> (uint(int32(16)) % 32))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v69)
		return
	case 3:
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v13)
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v15)
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v17)
		return
	}
}
func F_compare_lexeme_textfreq(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v8 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v38 < v6 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v11 = int32(4)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v13&int32(254) == int32(2) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v26 = int32(1)
	if v8&v26 != 0 {
		v38 = int32(base.Ui32(v8)>>(uint(v26)%32)) - v26
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v22 = v11
	goto L7
L6:
	;
	v22 = base.B2i32(v13 == int32(18)) << (uint(v11) % 32)
	goto L7
L7:
	;
	if v13 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = v11
	goto L10
L9:
	;
	v25 = v22
	goto L10
L10:
	;
	v38 = v25
	goto L1
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v38 = int32(base.Ui32(v32)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	return int32(1)
L13:
	;
	goto L14
L14:
	;
	if v6 < v38 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(-1)
L16:
	;
	goto L17
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = int32(1)
	if v8&v46 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v50 = v46
	goto L20
L19:
	;
	v50 = int32(4)
	goto L20
L20:
	;
	v51 = v7 + v50
	if v6 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	return v95
L22:
	;
	v95 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v57 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v58 = v45
	v59 = v51
	v60 = v6
	v61 = v57
	goto L29
L26:
	;
	v83 = v51
	v87 = int32(0)
	goto L27
L27:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v95 = v87 - v88
	goto L21
L28:
	;
	v83 = v78
	v87 = v80
	goto L27
L29:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v61 != v63 {
		v78 = v59
		v80 = v61
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v78 = v72
	v80 = int32(0)
	goto L28
L31:
	;
	if v63 == int32(0) {
		v78 = v59
		v80 = v61
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v68 = v60 - int32(1)
	if v68 == int32(0) {
		v78 = v59
		v80 = v61
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v71 = int32(1)
	v72 = v59 + v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v73 != 0 {
		v58 = v58 + v71
		v59 = v72
		v60 = v68
		v61 = v73
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
}
func F_compare_scalars_simple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v9 = m.T0[v8].(func(*base.Module, int32, int32, int32) int32)(m, v6, v7, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 < int32(0) {
			v16 = int32(1)
		} else {
			v16 = int32(0) - v9
		}
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
		if v17 != 0 {
			v18 = v16
		} else {
			v18 = v9
		}
		return v18
	}
}
func F_compute_bucket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v20 = base.I32_extend16_s(v19)
	v22 = base.B2i32(int32(0) <= v20)
	if int32(0) <= v20 {
		v23 = int32(-8)
	} else {
		v23 = int32(-6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(base.Ui32(int32(base.Ui32(v14)>>(uint(int32(2))%32))+v23) >> (uint(int32(1)) % 32))
	if int32(0) <= v20 {
		v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
		v38 = v28
	} else {
		v38 = v19<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v19&int32(63)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v38
	v40 = int32(_a_F_compute_bucket_0)
	v41 = v19 & v40
	if v41 != v40 {
		if v41 != int32(_a_F_compute_bucket_1) {
			v52 = v41
		} else {
			v52 = v19 << (uint(int32(1)) % 32) & int32(_a_F_compute_bucket_2)
		}
	} else {
		v52 = v19 & int32(_a_F_compute_bucket_3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v52
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v54
	v63 = base.B2i32(v20 < v54)
	if v20 < v54 {
		v64 = int32(base.Ui32(v19)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v64 = v19 & int32(_a_F_compute_bucket_4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v64
	if v20 < v54 {
		v68 = int32(6)
	} else {
		v68 = int32(8)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = l1 + v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v77 = base.I32_extend16_s(v76)
	v79 = base.B2i32(int32(0) <= v77)
	if int32(0) <= v77 {
		v80 = int32(-8)
	} else {
		v80 = int32(-6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(base.Ui32(int32(base.Ui32(v71)>>(uint(int32(2))%32))+v80) >> (uint(int32(1)) % 32))
	if int32(0) <= v77 {
		v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
		v95 = v85
	} else {
		v95 = v76<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v76&int32(63)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v95
	v97 = int32(_a_F_compute_bucket_0)
	v98 = v76 & v97
	if v98 != v97 {
		if v98 != int32(_a_F_compute_bucket_1) {
			v109 = v98
		} else {
			v109 = v76 << (uint(int32(1)) % 32) & int32(_a_F_compute_bucket_2)
		}
	} else {
		v109 = v76 & int32(_a_F_compute_bucket_3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v109
	v111 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v111
	v120 = base.B2i32(v77 < v111)
	if v77 < v111 {
		v121 = int32(base.Ui32(v76)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v121 = v76 & int32(_a_F_compute_bucket_4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v121
	if v77 < v111 {
		v125 = int32(6)
	} else {
		v125 = int32(8)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l2 + v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v134 = base.I32_extend16_s(v133)
	v136 = base.B2i32(int32(0) <= v134)
	if int32(0) <= v134 {
		v137 = int32(-8)
	} else {
		v137 = int32(-6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(base.Ui32(int32(base.Ui32(v128)>>(uint(int32(2))%32))+v137) >> (uint(int32(1)) % 32))
	if int32(0) <= v134 {
		v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
		v152 = v142
	} else {
		v152 = v133<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v133&int32(63)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v152
	v154 = int32(_a_F_compute_bucket_0)
	v155 = v133 & v154
	if v155 != v154 {
		if v155 != int32(_a_F_compute_bucket_1) {
			v166 = v155
		} else {
			v166 = v133 << (uint(int32(1)) % 32) & int32(_a_F_compute_bucket_2)
		}
	} else {
		v166 = v133 & int32(_a_F_compute_bucket_3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v166
	v168 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v168
	v173 = base.B2i32(v134 < v168)
	if v134 < v168 {
		v174 = int32(6)
	} else {
		v174 = int32(8)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l0 + v174
	if v134 < v168 {
		v183 = int32(base.Ui32(v133)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v183 = v133 & int32(_a_F_compute_bucket_4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v183
	v186 = v12 + int32(8)
	F_sub_var(m, v186, v12+int32(56), v186)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		return
	} else {
		v194 = v12 + int32(32)
		F_sub_var(m, v194, v12+int32(56), v194)
		mBase = m.M
		v200 = m.ExcPending
		if v200 != 0 {
			return
		} else {
			v202 = v12 + int32(8)
			v205 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			F_mul_var(m, v202, l3, v202, v205+v206)
			mBase = m.M
			v209 = m.ExcPending
			if v209 != 0 {
				return
			} else {
				v214 = int32(0)
				F_div_var(m, v12+int32(8), v12+int32(32), l4, v214, v214, int32(1))
				mBase = m.M
				v218 = m.ExcPending
				if v218 != 0 {
					return
				} else {
					F_add_var(m, l4, int32(_a_F_compute_bucket_5), l4)
					mBase = m.M
					v221 = m.ExcPending
					if v221 != 0 {
						return
					} else {
						v222 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
						if v222 != 0 {
							F_pfree(m, v222)
							mBase = m.M
							v224 = m.ExcPending
							if v224 != 0 {
								return
							} else {
								v225 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
								if v225 != 0 {
									F_pfree(m, v225)
									mBase = m.M
									v227 = m.ExcPending
									if v227 != 0 {
										return
									} else {
										m.G0 = v12 + int32(80)
										return
									}
								} else {
									m.G0 = v12 + int32(80)
									return
								}
							}
						} else {
							v225 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
							if v225 != 0 {
								F_pfree(m, v225)
								mBase = m.M
								v227 = m.ExcPending
								if v227 != 0 {
									return
								} else {
									m.G0 = v12 + int32(80)
									return
								}
							} else {
								m.G0 = v12 + int32(80)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_compute_distinct_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 float64
	_ = v88
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 float64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 float64
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v412 int32
	_ = v412
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v439 float64
	_ = v439
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v453 float64
	_ = v453
	var v455 float32
	_ = v455
	var v458 float64
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 float64
	_ = v564
	var v568 float64
	_ = v568
	var v575 float64
	_ = v575
	var v576 float64
	_ = v576
	var v578 float64
	_ = v578
	var v584 float64
	_ = v584
	var v585 float64
	_ = v585
	var v588 float64
	_ = v588
	var v590 float64
	_ = v590
	var v623 float32
	_ = v623
	var v625 float64
	_ = v625
	var v631 float32
	_ = v631
	var v633 float32
	_ = v633
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v825 float32
	_ = v825
	var v826 float64
	_ = v826
	var v827 float32
	_ = v827
	var v829 float64
	_ = v829
	var v835 int32
	_ = v835
	var v841 float64
	_ = v841
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v867 int32
	_ = v867
	var v869 float64
	_ = v869
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v895 float64
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v905 int32
	_ = v905
	var v907 float64
	_ = v907
	var v925 int32
	_ = v925
	var v927 float64
	_ = v927
	var v935 int32
	_ = v935
	var v942 int32
	_ = v942
	var v944 float64
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v956 float64
	_ = v956
	var v968 float64
	_ = v968
	var v977 int32
	_ = v977
	var v982 float64
	_ = v982
	var v994 float64
	_ = v994
	var v995 float64
	_ = v995
	var v999 float64
	_ = v999
	var v1002 float64
	_ = v1002
	var v1005 float64
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1009 float64
	_ = v1009
	var v1013 float64
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1019 float64
	_ = v1019
	var v1021 float64
	_ = v1021
	var v1027 float64
	_ = v1027
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1064 int32
	_ = v1064
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	v5 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(32)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+78)))
	if v34 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+76)))
	v40 = int32(_a_F_compute_distinct_stats_0)
	v45 = base.B2i32(v37 < int32(0))
	v46 = base.B2i32(v37&v40 == v40)
	goto L3
L2:
	;
	v45 = v5
	v46 = v5
	goto L3
L3:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v48 = int32(10)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = v49 << (uint(int32(1)) % 32)
	if v51 <= v48 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v54 = v48
	goto L6
L5:
	;
	v54 = v51
	goto L6
L6:
	;
	v57 = F_palloc(m, v54<<(uint(int32(3))%32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	F_fmgr_info(m, v59, v31+int32(4))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if l2 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	m.G0 = v31 + int32(32)
	return
L11:
	;
	v72 = v5
	v78 = v5
	v80 = v5
	v81 = v5
	v83 = v5
	v88 = float64(0)
	goto L12
L12:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L14
	}
L13:
	;
	if int32(0) < v432 {
		goto L74
	} else {
		goto L75
	}
L14:
	;
	v99 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v78, v31+int32(3))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+3)))
	if v101 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v446 = v78 + int32(1)
	if v446 != l2 {
		v72 = v423
		v78 = v446
		v80 = v431
		v81 = v432
		v83 = v434
		v88 = v439
		goto L12
	} else {
		goto L71
	}
L17:
	;
	v423 = v72
	v431 = v80 + int32(1)
	v432 = v81
	v434 = v83
	v439 = v88
	goto L16
L18:
	;
	goto L19
L19:
	;
	v107 = v81 + int32(1)
	if v46 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v156 = int32(0)
	if v156 < v72 {
		goto L43
	} else {
		goto L44
	}
L21:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v108 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	if v45 == int32(0) {
		v154 = v99
		v155 = v88
		goto L20
	} else {
		goto L38
	}
L24:
	;
	v136 = base.F64_add(v88, base.F64_convert_i32_u(v134))
	v137 = F_toast_raw_datum_size(m, v99)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L7
	} else {
		goto L33
	}
L25:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if base.Ui32((v112-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v134 = int32(6)
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v126 = int32(1)
	if v108&v126 != 0 {
		v134 = int32(base.Ui32(v108) >> (uint(v126) % 32))
		goto L24
	} else {
		goto L32
	}
L28:
	;
	v119 = int32(18)
	if v112&int32(255) == v119 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v125 = v119
	goto L31
L30:
	;
	v125 = int32(2)
	goto L31
L31:
	;
	v134 = v125
	goto L24
L32:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v134 = int32(base.Ui32(v130) >> (uint(int32(2)) % 32))
	goto L24
L33:
	;
	if base.Ui32(int32(1025)) <= base.Ui32(v137) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v423 = v72
	v431 = v80
	v432 = v107
	v434 = v83 + int32(1)
	v439 = v136
	goto L16
L35:
	;
	goto L36
L36:
	;
	v143 = F_pg_detoast_datum(m, v99)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v154 = v143
	v155 = v136
	goto L20
L38:
	;
	v147 = F_strlen(m, v99)
	mBase = m.M
	v154 = v99
	v155 = base.F64_add(v88, base.F64_convert_i32_u(v147+int32(1)))
	goto L20
L39:
	;
	if v210 < v234 {
		goto L68
	} else {
		goto L69
	}
L40:
	;
	if v210 == v233+v72-int32(2) {
		goto L39
	} else {
		goto L64
	}
L41:
	;
	v303 = int32(3)
	v305 = v57 + v236<<(uint(v303)%32)
	v308 = v57 + v234<<(uint(v303)%32)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v308-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = v311
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v308-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v305)+4)) = v315
	v319 = v234 - int32(2)
	v321 = v236
	goto L40
L42:
	;
	v248 = v192 + int32(4)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v249 + int32(1)
	if v163 == int32(0) {
		v423 = v72
		v431 = v80
		v432 = v107
		v434 = v83
		v439 = v155
		goto L16
	} else {
		goto L59
	}
L43:
	;
	v163 = v156
	v164 = v72
	goto L46
L44:
	;
	v210 = v72
	goto L45
L45:
	;
	v233 = base.B2i32(v72 < v54)
	v234 = v72 + v233
	v236 = v234 - int32(1)
	if v236 <= v210 {
		goto L39
	} else {
		goto L57
	}
L46:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v192 = v57 + v163<<(uint(int32(3))%32)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v194 = F_FunctionCall2Coll(m, v31+int32(4), v189, v154, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L7
	} else {
		goto L48
	}
L47:
	;
	v210 = v201
	goto L45
L48:
	;
	if v194 != 0 {
		goto L42
	} else {
		goto L49
	}
L49:
	;
	if v163 < v164 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	if v197 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v201 = v164
	goto L52
L52:
	;
	v203 = v163 + int32(1)
	if v203 != v72 {
		v163 = v203
		v164 = v201
		goto L46
	} else {
		goto L56
	}
L53:
	;
	v200 = v163
	goto L55
L54:
	;
	v200 = v164
	goto L55
L55:
	;
	v201 = v200
	goto L52
L56:
	;
	goto L47
L57:
	;
	if (v72-v233+v210)&int32(1) == int32(0) {
		goto L41
	} else {
		goto L58
	}
L58:
	;
	v319 = v236
	v321 = v234
	goto L40
L59:
	;
	v259 = v163
	goto L60
L60:
	;
	v285 = v57 + v259<<(uint(int32(3))%32)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	v288 = v285 - int32(4)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if v286 <= v289 {
		v423 = v72
		v431 = v80
		v432 = v107
		v434 = v83
		v439 = v155
		goto L16
	} else {
		goto L62
	}
L61:
	;
	v423 = v72
	v431 = v80
	v432 = v107
	v434 = v83
	v439 = v155
	goto L16
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+4)) = v289
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v294 = v285 - int32(8)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v286
	v299 = int32(1)
	if v299 < v259 {
		v259 = v259 - v299
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v327 = v319
	v330 = v321
	goto L65
L65:
	;
	v351 = int32(3)
	v353 = v57 + v327<<(uint(v351)%32)
	v356 = v57 + v330<<(uint(v351)%32)
	v357 = int32(16)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v356-v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v359
	v361 = int32(12)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v356-v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+4)) = v363
	v366 = v327 - int32(1)
	v369 = v57 + v366<<(uint(v351)%32)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v353-v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v372
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v353-v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v376
	v379 = v327 - int32(2)
	if v210 < v379 {
		v327 = v379
		v330 = v366
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L39
L67:
	;
	goto L66
L68:
	;
	v412 = v57 + v210<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v412)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v412))) = v154
	goto L70
L69:
	;
	goto L70
L70:
	;
	v423 = v234
	v431 = v80
	v432 = v107
	v434 = v83
	v439 = v155
	goto L16
L71:
	;
	goto L13
L72:
	;
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v623
	v625 = base.F64_promote_f32(v623)
	if base.F64_gt(v625, base.F64_mul(l3, float64(0.1))) != 0 {
		goto L111
	} else {
		goto L112
	}
L73:
	;
	if v54 <= v423 {
		goto L97
	} else {
		goto L98
	}
L74:
	;
	v450 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v450)
	v453 = base.F64_convert_i32_s(l2)
	v455 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v431), v453))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v455
	if v45 != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L76
L76:
	;
	if v431 <= int32(0) {
		goto L10
	} else {
		goto L93
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v467
	if int32(0) < v423 {
		goto L84
	} else {
		goto L85
	}
L78:
	;
	v458 = base.F64_div(v439, base.F64_convert_i32_u(v432))
	if base.F64_lt(base.F64_abs(v458), float64(2.147483648e+09)) != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v465 = int32(*(*int16)(unsafe.Add(mBase, uint32(v464)+76)))
	v467 = v465
	goto L77
L81:
	;
	v462 = base.I32_trunc_f64_s(v458)
	v467 = v462
	goto L77
L82:
	;
	goto L83
L83:
	;
	v467 = int32(-2147483648)
	goto L77
L84:
	;
	v471 = int32(0)
	v477 = v471
	v478 = v471
	goto L88
L85:
	;
	goto L86
L86:
	;
	v623 = base.F32_neg(base.F32_sub(float32(1), v455))
	goto L72
L87:
	;
	if v477 != 0 {
		v557 = v477
		v558 = v478
		goto L73
	} else {
		goto L92
	}
L88:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v57+v477<<(uint(int32(3))%32))+4))
	if v504 == int32(1) {
		goto L87
	} else {
		goto L90
	}
L89:
	;
	v557 = v423
	v558 = v507
	goto L73
L90:
	;
	v507 = v478 + v504
	v509 = v477 + int32(1)
	if v509 != v423 {
		v477 = v509
		v478 = v507
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	goto L86
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v546 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v546)
	v548 = int32(0)
	if v45 == v548 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v552 = int32(*(*int16)(unsafe.Add(mBase, uint32(v551)+76)))
	v553 = v552
	goto L96
L95:
	;
	v553 = v548
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v553
	goto L10
L97:
	;
	v562 = v432 - v558
	v563 = v562 + v557
	v564 = float64(0)
	v568 = base.F64_mul(l3, base.F64_sub(float64(1), base.F64_promote_f32(v455)))
	if base.F64_gt(v568, v564) == int32(0) {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	if v434 != 0 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	if v557 != v423 {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v623 = base.F32_convert_i32_s(v423)
	goto L72
L101:
	;
	if base.F64_lt(v584, v585) != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v584 = v564
	v585 = base.F64_convert_i32_s(v563)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v575 = base.F64_convert_i32_s(l2 - v431)
	v576 = base.F64_convert_i32_s(v563)
	v578 = base.F64_convert_i32_s(v562)
	v584 = base.F64_div(base.F64_mul(v575, v576), base.F64_add(base.F64_sub(v575, v578), base.F64_div(base.F64_mul(v575, v578), v568)))
	v585 = v576
	goto L101
L105:
	;
	v588 = v585
	goto L107
L106:
	;
	v588 = v584
	goto L107
L107:
	;
	if base.F64_gt(v588, v568) != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v590 = v568
	goto L110
L109:
	;
	v590 = v588
	goto L110
L110:
	;
	v623 = base.F32_demote_f64(base.F64_floor(base.F64_add(v590, float64(0.5))))
	goto L72
L111:
	;
	v631 = base.F32_demote_f64(base.F64_div(base.F64_neg(v625), l3))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v631
	v633 = v631
	goto L113
L112:
	;
	v633 = v623
	goto L113
L113:
	;
	if v54 <= v423 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v1064 <= int32(0) {
		goto L10
	} else {
		goto L172
	}
L115:
	;
	if v49 < v423 {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	if v434 != 0 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	if base.F32_gt(v633, float32(0)) == int32(0) {
		goto L115
	} else {
		goto L118
	}
L118:
	;
	if v423 <= v49 {
		v1064 = v423
		goto L114
	} else {
		goto L119
	}
L119:
	;
	goto L115
L120:
	;
	v641 = v49
	goto L122
L121:
	;
	v641 = v423
	goto L122
L122:
	;
	if v641 <= int32(0) {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	v645 = v641 & int32(3)
	v649 = F_palloc(m, v641<<(uint(int32(2))%32))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L7
	} else {
		goto L124
	}
L124:
	;
	v651 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v641) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v661 = v651
	v665 = int32(0)
	goto L128
L126:
	;
	v732 = v651
	goto L127
L127:
	;
	if v645 != 0 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v685 = int32(2)
	v688 = int32(3)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v57+v661<<(uint(v688)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v649+v661<<(uint(v685)%32)))) = v691
	v694 = v661 | int32(1)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v57+v694<<(uint(v688)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v649+v694<<(uint(v685)%32)))) = v701
	v704 = v661 | v685
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v57+v704<<(uint(v688)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v649+v704<<(uint(v685)%32)))) = v711
	v714 = v661 | v688
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v57+v714<<(uint(v688)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v649+v714<<(uint(v685)%32)))) = v721
	v723 = int32(4)
	v724 = v661 + v723
	v726 = v665 + v723
	if v726 != v641&int32(2147483644) {
		v661 = v724
		v665 = v726
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v732 = v724
	goto L127
L130:
	;
	goto L129
L131:
	;
	v760 = v732
	v762 = int32(0)
	goto L134
L132:
	;
	goto L133
L133:
	;
	v825 = *(*float32)(unsafe.Add(mBase, uint32(l0)+48))
	v826 = base.F64_promote_f32(v825)
	v827 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v829 = float64(0)
	v835 = int32(0)
	v841 = base.F64_convert_i32_s(l2)
	if base.F64_eq(l3, v841) != 0 {
		v1041 = v641
		goto L138
	} else {
		goto L139
	}
L134:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v57+v760<<(uint(int32(3))%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v649+v760<<(uint(int32(2))%32)))) = v790
	v792 = int32(1)
	v795 = v762 + v792
	if v795 != v645 {
		v760 = v760 + v792
		v762 = v795
		goto L134
	} else {
		goto L136
	}
L135:
	;
	goto L133
L136:
	;
	goto L135
L137:
	;
	v1064 = v1041
	goto L114
L138:
	;
	goto L137
L139:
	;
	if base.F64_le(l3, float64(1)) != 0 {
		v1041 = v641
		goto L138
	} else {
		goto L140
	}
L140:
	;
	if base.Ui32(v641) < base.Ui32(int32(2)) {
		v956 = v829
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if base.F64_lt(v826, float64(0)) != 0 {
		goto L154
	} else {
		goto L155
	}
L142:
	;
	v852 = v641 - int32(1)
	v853 = int32(3)
	v854 = v852 & v853
	if base.Ui32(v641-int32(2)) < base.Ui32(v853) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if v854 == int32(0) {
		v956 = v907
		goto L141
	} else {
		goto L150
	}
L144:
	;
	v905 = int32(0)
	v907 = v829
	goto L143
L145:
	;
	goto L146
L146:
	;
	v867 = int32(0)
	v869 = v829
	v878 = v835
	goto L147
L147:
	;
	v883 = v649 + v867<<(uint(int32(2))%32)
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v883)))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v883)+4))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v883)+8))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v883)+12))
	v895 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v869, base.F64_convert_i32_s(v884)), base.F64_convert_i32_s(v887)), base.F64_convert_i32_s(v890)), base.F64_convert_i32_s(v893))
	v896 = int32(4)
	v897 = v867 + v896
	v899 = v878 + v896
	if v899 != v852&int32(-4) {
		v867 = v897
		v869 = v895
		v878 = v899
		goto L147
	} else {
		goto L149
	}
L148:
	;
	v905 = v897
	v907 = v895
	goto L143
L149:
	;
	goto L148
L150:
	;
	v925 = v905
	v927 = v907
	v935 = v835
	goto L151
L151:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v649+v925<<(uint(int32(2))%32))))
	v944 = base.F64_add(v927, base.F64_convert_i32_s(v942))
	v945 = int32(1)
	v948 = v935 + v945
	if v948 != v854 {
		v925 = v925 + v945
		v927 = v944
		v935 = v948
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v956 = v944
	goto L141
L153:
	;
	goto L152
L154:
	;
	v968 = base.F64_mul(l3, base.F64_neg(v826))
	goto L156
L155:
	;
	v968 = v826
	goto L156
L156:
	;
	v977 = v641
	v982 = v956
	goto L157
L157:
	;
	v994 = float64(1)
	v995 = float64(0)
	v999 = base.F64_sub(base.F64_sub(v994, base.F64_div(v982, v841)), base.F64_promote_f32(v827))
	if base.F64_lt(v999, v995) != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v1041 = int32(0)
	goto L138
L159:
	;
	v1002 = v995
	goto L161
L160:
	;
	v1002 = v999
	goto L161
L161:
	;
	if base.F64_gt(v1002, float64(1)) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1005 = v994
	goto L164
L163:
	;
	v1005 = v1002
	goto L164
L164:
	;
	v1007 = v977 - int32(1)
	v1009 = base.F64_sub(v968, base.F64_convert_i32_u(v1007))
	if base.F64_gt(v1009, float64(1)) != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1013 = base.F64_div(v1005, v1009)
	goto L167
L166:
	;
	v1013 = v1005
	goto L167
L167:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v649+v1007<<(uint(int32(2))%32))))
	v1019 = base.F64_convert_i32_s(v1018)
	v1021 = base.F64_div(base.F64_mul(l3, v1019), v841)
	v1027 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l3, v841), base.F64_mul(base.F64_mul(v1021, v841), base.F64_sub(l3, v1021))), base.F64_mul(base.F64_mul(l3, l3), base.F64_add(l3, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v1013, v841), base.F64_add(v1027, v1027)), float64(0.5)), v1019) != 0 {
		v1041 = v977
		goto L138
	} else {
		goto L168
	}
L168:
	;
	if v1007 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v649-int32(8)+v977<<(uint(int32(2))%32))))
	v977 = v1007
	v982 = base.F64_sub(v982, base.F64_convert_i32_s(v1036))
	goto L157
L170:
	;
	goto L171
L171:
	;
	goto L158
L172:
	;
	v1089 = int32(_a_F_compute_distinct_stats_1)
	v1090 = *(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0]))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0])) = v1092
	v1095 = v1064 << (uint(int32(2)) % 32)
	v1096 = F_palloc(m, v1095)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L7
	} else {
		goto L173
	}
L173:
	;
	v1098 = F_palloc(m, v1095)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L7
	} else {
		goto L174
	}
L174:
	;
	v1104 = int32(0)
	goto L175
L175:
	;
	v1129 = v1104 << (uint(int32(2)) % 32)
	v1133 = v57 + v1104<<(uint(int32(3))%32)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1133)))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1135)+78)))
	v1137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1135)+76)))
	v1138 = F_datumCopy(m, v1134, v1136, v1137)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L7
	} else {
		goto L177
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0])) = v1090
	v1152 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v1152)
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1154
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v1098
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1157
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v1096
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v1064
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v1064
	goto L10
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1096+v1129))) = v1138
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+4))
	*(*float32)(unsafe.Add(mBase, uint32(v1129+v1098))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v1142), v453))
	v1148 = v1104 + int32(1)
	if v1148 != v1064 {
		v1104 = v1148
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
}
func F_compute_scalar_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v26 float64
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 float64
	_ = v112
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 float64
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 float64
	_ = v180
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 float64
	_ = v199
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v244 float64
	_ = v244
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var __phi288 int32
	_ = __phi288
	var v291 int32
	_ = v291
	var __phi291 int32
	_ = __phi291
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v369 int32
	_ = v369
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v438 int32
	_ = v438
	var v439 float64
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 float64
	_ = v446
	var v448 float32
	_ = v448
	var v451 float64
	_ = v451
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v469 int32
	_ = v469
	var v470 float64
	_ = v470
	var v474 float64
	_ = v474
	var v481 float64
	_ = v481
	var v482 float64
	_ = v482
	var v486 float64
	_ = v486
	var v492 float64
	_ = v492
	var v493 float64
	_ = v493
	var v496 float64
	_ = v496
	var v498 float64
	_ = v498
	var v508 float32
	_ = v508
	var v510 float64
	_ = v510
	var v516 float32
	_ = v516
	var v518 float32
	_ = v518
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v728 float32
	_ = v728
	var v729 float64
	_ = v729
	var v730 float32
	_ = v730
	var v732 float64
	_ = v732
	var v738 int32
	_ = v738
	var v744 float64
	_ = v744
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v770 int32
	_ = v770
	var v772 float64
	_ = v772
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v798 float64
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v810 float64
	_ = v810
	var v828 int32
	_ = v828
	var v830 float64
	_ = v830
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v847 float64
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v859 float64
	_ = v859
	var v871 float64
	_ = v871
	var v880 int32
	_ = v880
	var v885 float64
	_ = v885
	var v897 float64
	_ = v897
	var v898 float64
	_ = v898
	var v902 float64
	_ = v902
	var v905 float64
	_ = v905
	var v908 float64
	_ = v908
	var v910 int32
	_ = v910
	var v912 float64
	_ = v912
	var v916 float64
	_ = v916
	var v921 int32
	_ = v921
	var v922 float64
	_ = v922
	var v924 float64
	_ = v924
	var v930 float64
	_ = v930
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v962 int32
	_ = v962
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1080 int32
	_ = v1080
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1220 int32
	_ = v1220
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1235 int32
	_ = v1235
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1339 int32
	_ = v1339
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1474 int32
	_ = v1474
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1506 float64
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1511 float64
	_ = v1511
	var v1513 float64
	_ = v1513
	var v1515 float64
	_ = v1515
	var v1533 int32
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1549 int32
	_ = v1549
	var v1554 float32
	_ = v1554
	var v1557 float64
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	v5 = int32(0)
	v26 = float64(0)
	v33 = m.G0
	v35 = v33 - int32(48)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+78)))
	if v38 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+76)))
	v44 = int32(_a_F_compute_scalar_stats_0)
	v49 = base.B2i32(v41&v44 == v44)
	v50 = base.B2i32(v41 < int32(0))
	goto L3
L2:
	;
	v49 = v5
	v50 = v5
	goto L3
L3:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v55 = F_palloc(m, l2<<(uint(int32(3))%32))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v59 = F_palloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v63 = F_palloc(m, v51<<(uint(int32(3))%32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v65 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+44)) = v65
	v67 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+36)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v35)+28)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v35)+20)) = v67
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v76
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)) = uint8(v65)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	F_PrepareSortSupportFromOrderingOp(m, v80, v35+int32(12))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if l2 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.G0 = v35 + int32(48)
	return
L10:
	;
	v91 = v5
	v96 = v5
	v100 = v5
	v106 = v5
	v107 = v5
	v112 = v26
	goto L11
L11:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	if int32(0) < v194 {
		goto L39
	} else {
		goto L40
	}
L13:
	;
	v124 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v91, v35+int32(4))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)))
	if v126 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v201 = v91 + int32(1)
	if v201 != l2 {
		v91 = v201
		v96 = v194
		v100 = v196
		v106 = v197
		v107 = v198
		v112 = v199
		goto L11
	} else {
		goto L38
	}
L16:
	;
	v194 = v96
	v196 = v100 + int32(1)
	v197 = v106
	v198 = v107
	v199 = v112
	goto L15
L17:
	;
	goto L18
L18:
	;
	v132 = v106 + int32(1)
	if v49 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v183 = v55 + v96<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v59+v96<<(uint(int32(2))%32)))) = v96
	v194 = v96 + int32(1)
	v196 = v100
	v197 = v132
	v198 = v107
	v199 = v180
	goto L15
L20:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v133 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	if v50 == int32(0) {
		v177 = v124
		v180 = v112
		goto L19
	} else {
		goto L37
	}
L23:
	;
	v161 = base.F64_add(v112, base.F64_convert_i32_u(v159))
	v162 = F_toast_raw_datum_size(m, v124)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L32
	}
L24:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)))
	if base.Ui32((v137-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v159 = int32(6)
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v151 = int32(1)
	if v133&v151 != 0 {
		v159 = int32(base.Ui32(v133) >> (uint(v151) % 32))
		goto L23
	} else {
		goto L31
	}
L27:
	;
	v144 = int32(18)
	if v137&int32(255) == v144 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v150 = v144
	goto L30
L29:
	;
	v150 = int32(2)
	goto L30
L30:
	;
	v159 = v150
	goto L23
L31:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v159 = int32(base.Ui32(v155) >> (uint(int32(2)) % 32))
	goto L23
L32:
	;
	if base.Ui32(int32(1025)) <= base.Ui32(v162) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v194 = v96
	v196 = v100
	v197 = v132
	v198 = v107 + int32(1)
	v199 = v161
	goto L15
L34:
	;
	goto L35
L35:
	;
	v168 = F_pg_detoast_datum(m, v124)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v177 = v168
	v180 = v161
	goto L19
L37:
	;
	v172 = F_strlen(m, v124)
	mBase = m.M
	v177 = v124
	v180 = base.F64_add(v112, base.F64_convert_i32_u(v172+int32(1)))
	goto L19
L38:
	;
	goto L12
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v35 + int32(12)
	F_qsort_interruptible(m, v55, v194, int32(8), int32(509), v35+int32(4))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if int32(0) < v197 {
		goto L230
	} else {
		goto L231
	}
L42:
	;
	v219 = int32(0)
	v232 = v5
	v233 = v5
	v234 = v5
	v239 = v5
	v244 = v26
	goto L43
L43:
	;
	v251 = v232 + int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v55+v219<<(uint(int32(3))%32))+4))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v59+v256<<(uint(int32(2))%32))))
	if v256 != v262 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v443 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v443)
	v446 = base.F64_convert_i32_s(l2)
	v448 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v196), v446))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v448
	if v50 != 0 {
		goto L64
	} else {
		goto L65
	}
L45:
	;
	v421 = v233
	v422 = v234
	v427 = v239
	v438 = v251
	goto L47
L46:
	;
	if v251 < int32(2) {
		v388 = v233
		v394 = v239
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v439 = base.F64_add(base.F64_mul(base.F64_convert_i32_u(v219), base.F64_convert_i32_s(v256)), v244)
	v441 = v219 + int32(1)
	if v441 != v194 {
		v219 = v441
		v232 = v438
		v233 = v421
		v234 = v422
		v239 = v427
		v244 = v439
		goto L43
	} else {
		goto L62
	}
L48:
	;
	v421 = v388
	v422 = v234 + int32(1)
	v427 = v394
	v438 = int32(0)
	goto L47
L49:
	;
	v269 = v239 + int32(1)
	v270 = base.B2i32(v233 < v51)
	if v270 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v63-int32(8)+v233<<(uint(int32(3))%32))))
	if v251 <= v276 {
		v388 = v233
		v394 = v269
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v278 = v270 + v233
	v280 = v278 - int32(1)
	if v280 <= int32(0) {
		v340 = v280
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v369 = v63 + v340<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v219 - v232
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v251
	v388 = v278
	v394 = v269
	goto L48
L55:
	;
	__phi288 = v280
	__phi291 = v278
	v288 = __phi288
	v291 = __phi291
	goto L56
L56:
	;
	v317 = v63 + v291<<(uint(int32(3))%32)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v317-int32(16))))
	if v251 <= v320 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v340 = int32(0)
	goto L54
L58:
	;
	v340 = v288
	goto L54
L59:
	;
	goto L60
L60:
	;
	v324 = v63 + v288<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = v320
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v317-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v324)+4)) = v328
	v330 = int32(1)
	if v330 < v288 {
		__phi288 = v288 - v330
		__phi291 = v288
		v288 = __phi288
		v291 = __phi291
		goto L56
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	goto L44
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v460
	if v427 == int32(0) {
		v508 = base.F32_neg(base.F32_sub(float32(1), v448))
		goto L70
	} else {
		goto L71
	}
L64:
	;
	v451 = base.F64_div(v199, base.F64_convert_i32_s(v197))
	if base.F64_lt(base.F64_abs(v451), float64(2.147483648e+09)) != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v458 = int32(*(*int16)(unsafe.Add(mBase, uint32(v457)+76)))
	v460 = v458
	goto L63
L67:
	;
	v455 = base.I32_trunc_f64_s(v451)
	v460 = v455
	goto L63
L68:
	;
	goto L69
L69:
	;
	v460 = int32(-2147483648)
	goto L63
L70:
	;
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v508
	v510 = base.F64_promote_f32(v508)
	if base.F64_gt(v510, base.F64_mul(l3, float64(0.1))) != 0 {
		goto L85
	} else {
		goto L86
	}
L71:
	;
	if v198 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v469 = v422 + v198
	v470 = float64(0)
	v474 = base.F64_mul(l3, base.F64_sub(float64(1), base.F64_promote_f32(v448)))
	if base.F64_gt(v474, v470) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	if v422 != v427 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v508 = base.F32_convert_i32_s(v427)
	goto L70
L75:
	;
	if base.F64_lt(v492, v493) != 0 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v492 = v470
	v493 = base.F64_convert_i32_s(v469)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v481 = base.F64_convert_i32_s(l2 - v196)
	v482 = base.F64_convert_i32_s(v469)
	v486 = base.F64_convert_i32_s(v198 - v427 + v422)
	v492 = base.F64_div(base.F64_mul(v481, v482), base.F64_add(base.F64_sub(v481, v486), base.F64_div(base.F64_mul(v481, v486), v474)))
	v493 = v482
	goto L75
L79:
	;
	v496 = v493
	goto L81
L80:
	;
	v496 = v492
	goto L81
L81:
	;
	if base.F64_gt(v496, v474) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v498 = v474
	goto L84
L83:
	;
	v498 = v496
	goto L84
L84:
	;
	v508 = base.F32_demote_f64(base.F64_floor(base.F64_add(v498, float64(0.5))))
	goto L70
L85:
	;
	v516 = base.F32_demote_f64(base.F64_div(base.F64_neg(v510), l3))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v516
	v518 = v516
	goto L87
L86:
	;
	v518 = v508
	goto L87
L87:
	;
	if v421 != v422 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v1114 = v422 - v1080
	if v51 < v1114 {
		goto L154
	} else {
		goto L155
	}
L89:
	;
	v993 = int32(0)
	if v962 <= v993 {
		v1080 = v962
		v1111 = v993
		goto L88
	} else {
		goto L147
	}
L90:
	;
	if v51 < v421 {
		goto L95
	} else {
		goto L96
	}
L91:
	;
	if v198 != 0 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	if base.F32_gt(v518, float32(0)) == int32(0) {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	if v422 <= v51 {
		v962 = v422
		goto L89
	} else {
		goto L94
	}
L94:
	;
	goto L90
L95:
	;
	v528 = v51
	goto L97
L96:
	;
	v528 = v421
	goto L97
L97:
	;
	if v528 <= int32(0) {
		v1080 = v528
		v1111 = int32(0)
		goto L88
	} else {
		goto L98
	}
L98:
	;
	v532 = v528 & int32(3)
	v536 = F_palloc(m, v528<<(uint(int32(2))%32))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v538 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v528) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v548 = v538
	v552 = int32(0)
	goto L103
L101:
	;
	v623 = v538
	goto L102
L102:
	;
	if v532 != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v576 = int32(2)
	v579 = int32(3)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v63+v548<<(uint(v579)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v536+v548<<(uint(v576)%32)))) = v582
	v585 = v548 | int32(1)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v63+v585<<(uint(v579)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v536+v585<<(uint(v576)%32)))) = v592
	v595 = v548 | v576
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v63+v595<<(uint(v579)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v536+v595<<(uint(v576)%32)))) = v602
	v605 = v548 | v579
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v63+v605<<(uint(v579)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v536+v605<<(uint(v576)%32)))) = v612
	v614 = int32(4)
	v615 = v548 + v614
	v617 = v552 + v614
	if v617 != v528&int32(2147483644) {
		v548 = v615
		v552 = v617
		goto L103
	} else {
		goto L105
	}
L104:
	;
	v623 = v615
	goto L102
L105:
	;
	goto L104
L106:
	;
	v655 = v623
	v658 = int32(0)
	goto L109
L107:
	;
	goto L108
L108:
	;
	v728 = *(*float32)(unsafe.Add(mBase, uint32(l0)+48))
	v729 = base.F64_promote_f32(v728)
	v730 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v732 = float64(0)
	v738 = int32(0)
	v744 = base.F64_convert_i32_s(l2)
	if base.F64_eq(l3, v744) != 0 {
		v944 = v528
		goto L113
	} else {
		goto L114
	}
L109:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v63+v655<<(uint(int32(3))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v536+v655<<(uint(int32(2))%32)))) = v689
	v691 = int32(1)
	v694 = v658 + v691
	if v694 != v532 {
		v655 = v655 + v691
		v658 = v694
		goto L109
	} else {
		goto L111
	}
L110:
	;
	goto L108
L111:
	;
	goto L110
L112:
	;
	v962 = v944
	goto L89
L113:
	;
	goto L112
L114:
	;
	if base.F64_le(l3, float64(1)) != 0 {
		v944 = v528
		goto L113
	} else {
		goto L115
	}
L115:
	;
	if base.Ui32(v528) < base.Ui32(int32(2)) {
		v859 = v732
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if base.F64_lt(v729, float64(0)) != 0 {
		goto L129
	} else {
		goto L130
	}
L117:
	;
	v755 = v528 - int32(1)
	v756 = int32(3)
	v757 = v755 & v756
	if base.Ui32(v528-int32(2)) < base.Ui32(v756) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v757 == int32(0) {
		v859 = v810
		goto L116
	} else {
		goto L125
	}
L119:
	;
	v808 = int32(0)
	v810 = v732
	goto L118
L120:
	;
	goto L121
L121:
	;
	v770 = int32(0)
	v772 = v732
	v781 = v738
	goto L122
L122:
	;
	v786 = v536 + v770<<(uint(int32(2))%32)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v786)))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v786)+4))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v786)+8))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v786)+12))
	v798 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v772, base.F64_convert_i32_s(v787)), base.F64_convert_i32_s(v790)), base.F64_convert_i32_s(v793)), base.F64_convert_i32_s(v796))
	v799 = int32(4)
	v800 = v770 + v799
	v802 = v781 + v799
	if v802 != v755&int32(-4) {
		v770 = v800
		v772 = v798
		v781 = v802
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v808 = v800
	v810 = v798
	goto L118
L124:
	;
	goto L123
L125:
	;
	v828 = v808
	v830 = v810
	v838 = v738
	goto L126
L126:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v536+v828<<(uint(int32(2))%32))))
	v847 = base.F64_add(v830, base.F64_convert_i32_s(v845))
	v848 = int32(1)
	v851 = v838 + v848
	if v851 != v757 {
		v828 = v828 + v848
		v830 = v847
		v838 = v851
		goto L126
	} else {
		goto L128
	}
L127:
	;
	v859 = v847
	goto L116
L128:
	;
	goto L127
L129:
	;
	v871 = base.F64_mul(l3, base.F64_neg(v729))
	goto L131
L130:
	;
	v871 = v729
	goto L131
L131:
	;
	v880 = v528
	v885 = v859
	goto L132
L132:
	;
	v897 = float64(1)
	v898 = float64(0)
	v902 = base.F64_sub(base.F64_sub(v897, base.F64_div(v885, v744)), base.F64_promote_f32(v730))
	if base.F64_lt(v902, v898) != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v944 = int32(0)
	goto L113
L134:
	;
	v905 = v898
	goto L136
L135:
	;
	v905 = v902
	goto L136
L136:
	;
	if base.F64_gt(v905, float64(1)) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v908 = v897
	goto L139
L138:
	;
	v908 = v905
	goto L139
L139:
	;
	v910 = v880 - int32(1)
	v912 = base.F64_sub(v871, base.F64_convert_i32_u(v910))
	if base.F64_gt(v912, float64(1)) != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v916 = base.F64_div(v908, v912)
	goto L142
L141:
	;
	v916 = v908
	goto L142
L142:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v536+v910<<(uint(int32(2))%32))))
	v922 = base.F64_convert_i32_s(v921)
	v924 = base.F64_div(base.F64_mul(l3, v922), v744)
	v930 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l3, v744), base.F64_mul(base.F64_mul(v924, v744), base.F64_sub(l3, v924))), base.F64_mul(base.F64_mul(l3, l3), base.F64_add(l3, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v916, v744), base.F64_add(v930, v930)), float64(0.5)), v922) != 0 {
		v944 = v880
		goto L113
	} else {
		goto L143
	}
L143:
	;
	if v910 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v536-int32(8)+v880<<(uint(int32(2))%32))))
	v880 = v910
	v885 = base.F64_sub(v885, base.F64_convert_i32_s(v939))
	goto L132
L145:
	;
	goto L146
L146:
	;
	goto L133
L147:
	;
	v997 = int32(_a_F_compute_scalar_stats_1)
	v998 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1000
	v1003 = v962 << (uint(int32(2)) % 32)
	v1004 = F_palloc(m, v1003)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v1006 = F_palloc(m, v1003)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v1012 = int32(0)
	goto L150
L150:
	;
	v1041 = v1012 << (uint(int32(2)) % 32)
	v1043 = int32(3)
	v1045 = v63 + v1012<<(uint(v1043)%32)
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+4))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v55+v1046<<(uint(v1043)%32))))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1051)+78)))
	v1053 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1051)+76)))
	v1054 = F_datumCopy(m, v1050, v1052, v1053)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L4
	} else {
		goto L152
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v998
	v1068 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v1068)
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1070
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v1006
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v1004
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v962
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v962
	v1080 = v962
	v1111 = v1068
	goto L88
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1004+v1041))) = v1054
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1045)))
	*(*float32)(unsafe.Add(mBase, uint32(v1041+v1006))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v1058), v446))
	v1064 = v1012 + int32(1)
	if v1064 != v962 {
		v1012 = v1064
		goto L150
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	v1116 = v51 + int32(1)
	goto L156
L155:
	;
	v1116 = v1114
	goto L156
L156:
	;
	if int32(2) <= v1116 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v1119 = int32(0)
	F_qsort_interruptible(m, v63, v1080, int32(8), int32(510), v1119)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L4
	} else {
		goto L160
	}
L158:
	;
	v1474 = v1111
	goto L159
L159:
	;
	if v194 == int32(1) {
		goto L9
	} else {
		goto L228
	}
L160:
	;
	if v1111 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v1125 = int32(0)
	v1131 = v1125
	v1134 = v1125
	v1135 = v1119
	goto L164
L162:
	;
	v1339 = v194
	goto L163
L163:
	;
	v1364 = int32(_a_F_compute_scalar_stats_1)
	v1365 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1367
	v1369 = int32(1)
	v1370 = v1339 - v1369
	v1372 = v1116 - v1369
	v1373 = base.I32_div_s(v1370, v1372)
	if v1116 <= v1369 {
		goto L217
	} else {
		goto L218
	}
L164:
	;
	if v1080 <= v1135 {
		v1169 = v194
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v1339 = v1327
	goto L163
L166:
	;
	if v1330 < v194 {
		v1131 = v1330
		v1134 = v1327
		v1135 = v1328
		goto L164
	} else {
		goto L216
	}
L167:
	;
	v1171 = int32(3)
	v1173 = v55 + v1134<<(uint(v1171)%32)
	v1176 = v55 + v1131<<(uint(v1171)%32)
	v1177 = v1169 - v1131
	v1179 = v1177 << (uint(v1171) % 32)
	if v1173 == v1176 {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	v1162 = v63 + v1135<<(uint(int32(3))%32)
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1162)+4))
	if v1131 < v1163 {
		v1169 = v1163
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1162)))
	v1327 = v1134
	v1328 = v1135 + int32(1)
	v1330 = v1167 + v1163
	goto L166
L170:
	;
	v1327 = v1177 + v1134
	v1328 = v1135
	v1330 = v1169
	goto L166
L171:
	;
	goto L170
L172:
	;
	v1183 = v1173 + v1179
	if base.Ui32(v1176-v1183) <= base.Ui32(int32(0)-v1179<<(uint(int32(1))%32)) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v1190 = F___memcpy(m, v1173, v1176, v1179)
	mBase = m.M
	goto L170
L174:
	;
	goto L175
L175:
	;
	v1193 = (v1173 ^ v1176) & int32(3)
	if base.Ui32(v1173) < base.Ui32(v1176) {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	if v1295 == int32(0) {
		goto L171
	} else {
		goto L212
	}
L177:
	;
	if base.Ui32(v1273) <= base.Ui32(int32(3)) {
		v1294 = v1272
		v1295 = v1273
		v1296 = v1274
		goto L176
	} else {
		goto L208
	}
L178:
	;
	if v1193 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	goto L180
L180:
	;
	if v1193 != 0 {
		v1255 = v1179
		goto L191
	} else {
		goto L192
	}
L181:
	;
	v1294 = v1176
	v1295 = v1179
	v1296 = v1173
	goto L176
L182:
	;
	goto L183
L183:
	;
	if v1173&int32(3) == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1272 = v1176
	v1273 = v1179
	v1274 = v1173
	goto L177
L185:
	;
	goto L186
L186:
	;
	v1200 = v1176
	v1201 = v1179
	v1202 = v1173
	goto L187
L187:
	;
	if v1201 == int32(0) {
		goto L171
	} else {
		goto L189
	}
L188:
	;
	v1272 = v1209
	v1273 = v1211
	v1274 = v1213
	goto L177
L189:
	;
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1202))) = uint8(v1206)
	v1208 = int32(1)
	v1209 = v1200 + v1208
	v1211 = v1201 - v1208
	v1213 = v1202 + v1208
	if v1213&int32(3) != 0 {
		v1200 = v1209
		v1201 = v1211
		v1202 = v1213
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	if v1255 == int32(0) {
		goto L171
	} else {
		goto L204
	}
L192:
	;
	if v1183&int32(3) != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1220 = v1179
	goto L196
L194:
	;
	v1235 = v1179
	goto L195
L195:
	;
	if base.Ui32(v1235) <= base.Ui32(int32(3)) {
		v1255 = v1235
		goto L191
	} else {
		goto L200
	}
L196:
	;
	if v1220 == int32(0) {
		goto L171
	} else {
		goto L198
	}
L197:
	;
	v1235 = v1226
	goto L195
L198:
	;
	v1226 = v1220 - int32(1)
	v1227 = v1173 + v1226
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176+v1226))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1227))) = uint8(v1229)
	if v1227&int32(3) != 0 {
		v1220 = v1226
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v1242 = v1235
	goto L201
L201:
	;
	v1246 = v1242 - int32(4)
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1176+v1246)))
	*(*int32)(unsafe.Add(mBase, uint32(v1173+v1246))) = v1249
	if base.Ui32(int32(3)) < base.Ui32(v1246) {
		v1242 = v1246
		goto L201
	} else {
		goto L203
	}
L202:
	;
	v1255 = v1246
	goto L191
L203:
	;
	goto L202
L204:
	;
	v1262 = v1255
	goto L205
L205:
	;
	v1266 = v1262 - int32(1)
	v1269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1176+v1266))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1173+v1266))) = uint8(v1269)
	if v1266 != 0 {
		v1262 = v1266
		goto L205
	} else {
		goto L207
	}
L206:
	;
	goto L171
L207:
	;
	goto L206
L208:
	;
	v1279 = v1272
	v1280 = v1273
	v1281 = v1274
	goto L209
L209:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1279)))
	*(*int32)(unsafe.Add(mBase, uint32(v1281))) = v1283
	v1285 = int32(4)
	v1286 = v1279 + v1285
	v1288 = v1281 + v1285
	v1290 = v1280 - v1285
	if base.Ui32(int32(3)) < base.Ui32(v1290) {
		v1279 = v1286
		v1280 = v1290
		v1281 = v1288
		goto L209
	} else {
		goto L211
	}
L210:
	;
	v1294 = v1286
	v1295 = v1290
	v1296 = v1288
	goto L176
L211:
	;
	goto L210
L212:
	;
	v1301 = v1294
	v1302 = v1295
	v1303 = v1296
	goto L213
L213:
	;
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1301))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1303))) = uint8(v1305)
	v1307 = int32(1)
	v1312 = v1302 - v1307
	if v1312 != 0 {
		v1301 = v1301 + v1307
		v1302 = v1312
		v1303 = v1303 + v1307
		goto L213
	} else {
		goto L215
	}
L214:
	;
	goto L171
L215:
	;
	goto L214
L216:
	;
	goto L165
L217:
	;
	v1379 = v1369
	goto L219
L218:
	;
	v1379 = v1116
	goto L219
L219:
	;
	v1382 = F_palloc(m, v1116<<(uint(int32(2))%32))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L4
	} else {
		goto L220
	}
L220:
	;
	v1384 = int32(0)
	v1388 = v1384
	v1391 = v1384
	v1393 = v1384
	goto L221
L221:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v55+v1393<<(uint(int32(3))%32))))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1426)+78)))
	v1428 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1426)+76)))
	v1429 = F_datumCopy(m, v1425, v1427, v1428)
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L4
	} else {
		goto L223
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1365
	v1444 = int32(1)
	v1447 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v1111<<(uint(v1444)%32))+52)) = uint16(v1447)
	v1451 = l0 + v1111<<(uint(v1447)%32)
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1451-int32(-64)))) = v1454
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1451)+164)) = v1382
	*(*int32)(unsafe.Add(mBase, uint32(v1451)+84)) = v1456
	*(*int32)(unsafe.Add(mBase, uint32(v1451)+144)) = v1116
	v1474 = v1111 + v1444
	goto L159
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1382+v1391<<(uint(int32(2))%32)))) = v1429
	v1432 = v1388 + (v1370 - v1373*v1372)
	v1433 = base.B2i32(v1372 <= v1432)
	if v1372 <= v1432 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1437 = v1372
	goto L226
L225:
	;
	v1437 = int32(0)
	goto L226
L226:
	;
	v1440 = v1391 + int32(1)
	if v1440 != v1379 {
		v1388 = v1432 - v1437
		v1391 = v1440
		v1393 = v1433 + (v1393 + v1373)
		goto L221
	} else {
		goto L227
	}
L227:
	;
	goto L222
L228:
	;
	v1496 = int32(_a_F_compute_scalar_stats_1)
	v1497 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1499
	v1502 = F_palloc(m, int32(4))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L4
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1497
	v1506 = base.F64_convert_i32_u(v194)
	v1508 = int32(1)
	v1511 = base.F64_mul(v1506, base.F64_convert_i32_u(v194-v1508))
	v1513 = base.F64_mul(v1511, float64(0.5))
	v1515 = base.F64_mul(v1513, base.F64_neg(v1513))
	*(*float32)(unsafe.Add(mBase, uint32(v1502))) = base.F32_demote_f64(base.F64_div(base.F64_add(base.F64_mul(v1506, v439), v1515), base.F64_add(base.F64_mul(v1506, base.F64_div(base.F64_mul(v1511, base.F64_convert_i32_s(v194<<(uint(v1508)%32)-v1508)), float64(6))), v1515)))
	v1533 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v1474<<(uint(v1508)%32))+52)) = uint16(v1533)
	v1537 = l0 + v1474<<(uint(int32(2))%32)
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1537-int32(-64)))) = v1540
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+124)) = v1502
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+84)) = v1542
	*(*int32)(unsafe.Add(mBase, uint32(v1537)+104)) = v1508
	goto L9
L230:
	;
	v1549 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1549)
	v1554 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v196), base.F64_convert_i32_s(l2)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v1554
	if v50 != 0 {
		goto L234
	} else {
		goto L235
	}
L231:
	;
	goto L232
L232:
	;
	if v196 <= int32(0) {
		goto L9
	} else {
		goto L240
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1566
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = base.F32_neg(base.F32_sub(float32(1), v1554))
	goto L9
L234:
	;
	v1557 = base.F64_div(v199, base.F64_convert_i32_u(v197))
	if base.F64_lt(base.F64_abs(v1557), float64(2.147483648e+09)) != 0 {
		goto L237
	} else {
		goto L238
	}
L235:
	;
	goto L236
L236:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1564 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1563)+76)))
	v1566 = v1564
	goto L233
L237:
	;
	v1561 = base.I32_trunc_f64_s(v1557)
	v1566 = v1561
	goto L233
L238:
	;
	goto L239
L239:
	;
	v1566 = int32(-2147483648)
	goto L233
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v1576 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1576)
	v1578 = int32(0)
	if v50 == v1578 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1582 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1581)+76)))
	v1583 = v1582
	goto L243
L242:
	;
	v1583 = v1578
	goto L243
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1583
	goto L9
}
func F_connectby_text_serial(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v26 = F_text_to_cstring(m, v22)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v29 = F_pg_detoast_datum_packed(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = F_text_to_cstring(m, v29)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v34 = F_pg_detoast_datum_packed(m, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = F_text_to_cstring(m, v34)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v39 = F_pg_detoast_datum_packed(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = F_text_to_cstring(m, v39)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
									v44 = F_pg_detoast_datum_packed(m, v43)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v46 = F_text_to_cstring(m, v44)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int32(0)
										} else {
											v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v48 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_connectby_text_serial_0), int32(0))
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1076), int32(_a_F_connectby_text_serial_2))
															mBase = m.M
															v146 = m.ExcPending
															if v146 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
												if v51 != int32(383) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v134 = m.ExcPending
													if v134 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_connectby_text_serial_0), int32(0))
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1076), int32(_a_F_connectby_text_serial_2))
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+12)))
													if v54&int32(2) == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v150 = m.ExcPending
														if v150 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(1088))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_connectby_text_serial_3), int32(0))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1081), int32(_a_F_connectby_text_serial_2))
																	mBase = m.M
																	v162 = m.ExcPending
																	if v162 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														v59 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
														if v59 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v150 = m.ExcPending
															if v150 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(1088))
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_connectby_text_serial_3), int32(0))
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1081), int32(_a_F_connectby_text_serial_2))
																		mBase = m.M
																		v162 = m.ExcPending
																		if v162 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															}
														} else {
															v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
															v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															if v63 == int32(7) {
																v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
																v67 = F_pg_detoast_datum_packed(m, v66)
																mBase = m.M
																v68 = m.ExcPending
																if v68 != 0 {
																	return int32(0)
																} else {
																	v69 = F_text_to_cstring(m, v67)
																	mBase = m.M
																	v70 = m.ExcPending
																	if v70 != 0 {
																		return int32(0)
																	} else {
																		v74 = v69
																		v75 = int32(_a_F_connectby_text_serial_4)
																		v76 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																		v78 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
																		v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
																		*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																		v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
																		v82 = F_CreateTupleDescCopy(m, v81)
																		mBase = m.M
																		v83 = m.ExcPending
																		if v83 != 0 {
																			return int32(0)
																		} else {
																			v85 = base.B2i32(v63 == int32(7))
																			F_validateConnectbyTupleDesc(m, v82, v85, int32(1))
																			mBase = m.M
																			v88 = m.ExcPending
																			if v88 != 0 {
																				return int32(0)
																			} else {
																				v89 = F_TupleDescGetAttInMetadata(m, v82)
																				mBase = m.M
																				v90 = m.ExcPending
																				if v90 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = int32(2)
																					v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
																					*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(1)
																					F_SPI_connect_ext(m, int32(0))
																					mBase = m.M
																					v98 = m.ExcPending
																					if v98 != 0 {
																						return int32(0)
																					} else {
																						v99 = int32(_a_F_connectby_text_serial_4)
																						v100 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																						*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																						v109 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[1]))
																						v110 = F_tuplestore_begin_heap(m, int32(base.Ui32(v93&int32(4))>>(uint(int32(2))%32)), int32(0), v109)
																						mBase = m.M
																						v111 = m.ExcPending
																						if v111 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v100
																							F_build_tuplestore_recursively(m, v31, v36, v26, v41, v74, v46, v46, int32(0), v19+int32(12), v62, v85, int32(1), v89, v110)
																							mBase = m.M
																							v119 = m.ExcPending
																							if v119 != 0 {
																								return int32(0)
																							} else {
																								v120 = F_SPI_finish(m)
																								mBase = m.M
																								v121 = m.ExcPending
																								if v121 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = v82
																									*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v110
																									*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v76
																									m.G0 = v19 + int32(16)
																									return int32(0)
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v72 = F_pstrdup(m, int32(_a_F_connectby_text_serial_5))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return int32(0)
																} else {
																	v74 = v72
																	v75 = int32(_a_F_connectby_text_serial_4)
																	v76 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																	v78 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
																	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
																	*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																	v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
																	v82 = F_CreateTupleDescCopy(m, v81)
																	mBase = m.M
																	v83 = m.ExcPending
																	if v83 != 0 {
																		return int32(0)
																	} else {
																		v85 = base.B2i32(v63 == int32(7))
																		F_validateConnectbyTupleDesc(m, v82, v85, int32(1))
																		mBase = m.M
																		v88 = m.ExcPending
																		if v88 != 0 {
																			return int32(0)
																		} else {
																			v89 = F_TupleDescGetAttInMetadata(m, v82)
																			mBase = m.M
																			v90 = m.ExcPending
																			if v90 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = int32(2)
																				v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(1)
																				F_SPI_connect_ext(m, int32(0))
																				mBase = m.M
																				v98 = m.ExcPending
																				if v98 != 0 {
																					return int32(0)
																				} else {
																					v99 = int32(_a_F_connectby_text_serial_4)
																					v100 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																					*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																					v109 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[1]))
																					v110 = F_tuplestore_begin_heap(m, int32(base.Ui32(v93&int32(4))>>(uint(int32(2))%32)), int32(0), v109)
																					mBase = m.M
																					v111 = m.ExcPending
																					if v111 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v100
																						F_build_tuplestore_recursively(m, v31, v36, v26, v41, v74, v46, v46, int32(0), v19+int32(12), v62, v85, int32(1), v89, v110)
																						mBase = m.M
																						v119 = m.ExcPending
																						if v119 != 0 {
																							return int32(0)
																						} else {
																							v120 = F_SPI_finish(m)
																							mBase = m.M
																							v121 = m.ExcPending
																							if v121 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = v82
																								*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v110
																								*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v76
																								m.G0 = v19 + int32(16)
																								return int32(0)
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_contains_multiexpr_param(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 == int32(8) {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			return base.B2i32(v10 == int32(3))
		} else {
			v15 = F_expression_tree_walker_impl(m, l0, int32(1056), l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_contsel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_Float8GetDatum(m, float64(0.001))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_convert_case(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v79 int32
	_ = v79
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var __phi277 int32
	_ = __phi277
	var v285 int32
	_ = v285
	var __phi285 int32
	_ = __phi285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v313 int32
	_ = v313
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v544 int32
	_ = v544
	var v554 int32
	_ = v554
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v593 int32
	_ = v593
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v796 int32
	_ = v796
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v916 int32
	_ = v916
	var v923 int32
	_ = v923
	var v930 int32
	_ = v930
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v983 int32
	_ = v983
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1046 int32
	_ = v1046
	var v1053 int32
	_ = v1053
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1079 int32
	_ = v1079
	var v1091 int32
	_ = v1091
	var v1104 int32
	_ = v1104
	var v1118 int32
	_ = v1118
	v9 = int32(0)
	if l4 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = m.T0[l6].(func(*base.Module, int32) int32)(m, l7)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v28 = v9
	goto L3
L3:
	;
	if l3 == int32(0) {
		v1104 = v9
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v28 = v24
	goto L3
L6:
	;
	if base.Ui32(v1104) < base.Ui32(l1) {
		goto L299
	} else {
		goto L300
	}
L7:
	;
	if l5 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = int32(1)
	goto L10
L9:
	;
	v33 = int32(2)
	goto L10
L10:
	;
	v43 = v9
	v47 = v9
	v49 = l4
	v50 = v28
	goto L11
L11:
	;
	v55 = l2 + v47
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v56 == int32(0) {
		v1104 = v43
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v1104 = v1079
	goto L6
L13:
	;
	if base.I32_extend8_s(v56) < int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if l4 != int32(1) {
		v140 = v49
		v141 = v50
		goto L33
	} else {
		goto L34
	}
L15:
	;
	if v56&int32(224) == int32(192) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v117 = v56
	goto L17
L17:
	;
	v118 = int32(1)
	if base.Ui32(v117) < base.Ui32(int32(128)) {
		v131 = v117
		v132 = v118
		v133 = v118
		goto L14
	} else {
		goto L28
	}
L18:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+v55))))
	v117 = v113&int32(63) | v110
	goto L17
L19:
	;
	v110 = v56 << (uint(int32(6)) % 32) & int32(1984)
	v111 = int32(1)
	goto L18
L20:
	;
	goto L21
L21:
	;
	if v56&int32(240) == int32(224) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v110 = v56<<(uint(int32(12))%32)&int32(_a_F_convert_case_0) | v79&int32(63)<<(uint(int32(6))%32)
	v111 = int32(2)
	goto L18
L23:
	;
	goto L24
L24:
	;
	if v56&int32(248) != int32(240) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v131 = int32(-1)
	v132 = int32(0)
	v133 = int32(4)
	goto L14
L26:
	;
	goto L27
L27:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v98 = int32(63)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+2)))
	v110 = v56<<(uint(int32(18))%32)&int32(_a_F_convert_case_1) | v97&v98<<(uint(int32(12))%32) | v103&v98<<(uint(int32(6))%32)
	v111 = int32(3)
	goto L18
L28:
	;
	v122 = int32(0)
	if base.Ui32(v117) < base.Ui32(int32(2048)) {
		v131 = v117
		v132 = v122
		v133 = int32(2)
		goto L14
	} else {
		goto L29
	}
L29:
	;
	if base.Ui32(v117) < base.Ui32(int32(_a_F_convert_case_2)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v130 = int32(3)
	goto L32
L31:
	;
	v130 = int32(4)
	goto L32
L32:
	;
	v131 = v117
	v132 = v122
	v133 = v130
	goto L14
L33:
	;
	if v132 != 0 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	if v47 != v50 {
		v140 = int32(0)
		v141 = v50
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v138 = m.T0[l6].(func(*base.Module, int32) int32)(m, l7)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v140 = v33
	v141 = v138
	goto L33
L37:
	;
	v1091 = v47 + v133
	if l3 < int32(0) {
		v43 = v1079
		v47 = v1091
		v49 = v140
		v50 = v141
		goto L11
	} else {
		goto L297
	}
L38:
	;
	v968 = v43
	v970 = int32(0)
	goto L273
L39:
	;
	v874 = l0 + v43
	if base.Ui32(v850) <= base.Ui32(int32(2047)) {
		goto L267
	} else {
		goto L268
	}
L40:
	;
	v870 = v43 + int32(1)
	if base.Ui32(l1) < base.Ui32(v870) {
		goto L264
	} else {
		goto L265
	}
L41:
	;
	v864 = v43 + v133
	if base.Ui32(l1) < base.Ui32(v864) {
		goto L257
	} else {
		goto L258
	}
L42:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v849)))
	if base.Ui32(v850) < base.Ui32(int32(128)) {
		goto L40
	} else {
		goto L249
	}
L43:
	;
	v142 = int32(2)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140<<(uint(v142)%32))+uint32(_c_F_convert_case[0])))
	v849 = v146 + v131<<(uint(v142)%32) + int32(4)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v152 = int32(0)
	if base.Ui32(v131) < base.Ui32(int32(1416)) {
		v239 = v131
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if v246 == int32(0) {
		goto L41
	} else {
		goto L102
	}
L47:
	;
	goto L46
L48:
	;
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239<<(uint(int32(1))%32))+uint32(_c_F_convert_case[1]))))
	v246 = v244
	goto L47
L49:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_3)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_4)) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_5)) {
		v246 = v152
		goto L47
	} else {
		goto L77
	}
L53:
	;
	if base.Ui32(v131-int32(_a_F_convert_case_6)) <= base.Ui32(int32(95)) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_7)) {
		v246 = v152
		goto L47
	} else {
		goto L64
	}
L56:
	;
	v239 = v131 - int32(2840)
	goto L48
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_8)) {
		v246 = v152
		goto L47
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_9)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v239 = v131 - int32(3512)
	goto L48
L61:
	;
	goto L62
L62:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_10)) {
		v246 = v152
		goto L47
	} else {
		goto L63
	}
L63:
	;
	v239 = v131 - int32(_a_F_convert_case_11)
	goto L48
L64:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_12)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_13)) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_14)) {
		v246 = v152
		goto L47
	} else {
		goto L72
	}
L68:
	;
	v239 = v131 - int32(_a_F_convert_case_15)
	goto L48
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_16)) {
		v246 = v152
		goto L47
	} else {
		goto L71
	}
L71:
	;
	v239 = v131 - int32(_a_F_convert_case_17)
	goto L48
L72:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_18)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v239 = v131 - int32(_a_F_convert_case_19)
	goto L48
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_20)) {
		v246 = v152
		goto L47
	} else {
		goto L76
	}
L76:
	;
	v239 = v131 - int32(_a_F_convert_case_21)
	goto L48
L77:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_22)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_23)) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_24)) {
		v246 = v152
		goto L47
	} else {
		goto L93
	}
L81:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_25)) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_26)) {
		v246 = v152
		goto L47
	} else {
		goto L88
	}
L84:
	;
	v239 = v131 - int32(_a_F_convert_case_27)
	goto L48
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_28)) {
		v246 = v152
		goto L47
	} else {
		goto L87
	}
L87:
	;
	v239 = v131 - int32(_a_F_convert_case_29)
	goto L48
L88:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_30)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v239 = v131 - int32(_a_F_convert_case_31)
	goto L48
L90:
	;
	goto L91
L91:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_32)) {
		v246 = v152
		goto L47
	} else {
		goto L92
	}
L92:
	;
	v239 = v131 - int32(_a_F_convert_case_33)
	goto L48
L93:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_34)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	if base.Ui32(v131) <= base.Ui32(int32(_a_F_convert_case_35)) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	if base.Ui32(int32(67)) < base.Ui32(v131-int32(_a_F_convert_case_36)) {
		v246 = v152
		goto L47
	} else {
		goto L101
	}
L97:
	;
	v239 = v131 - int32(_a_F_convert_case_37)
	goto L48
L98:
	;
	goto L99
L99:
	;
	if base.Ui32(v131) < base.Ui32(int32(_a_F_convert_case_38)) {
		v246 = v152
		goto L47
	} else {
		goto L100
	}
L100:
	;
	v239 = v131 - int32(_a_F_convert_case_39)
	goto L48
L101:
	;
	v239 = v131 - int32(_a_F_convert_case_40)
	goto L48
L102:
	;
	if l5 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v820 = int32(2)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v140<<(uint(v820)%32))+uint32(_c_F_convert_case[0])))
	v849 = v824 + v246<<(uint(v820)%32)
	goto L42
L104:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_convert_case[2]))))
	if v253 == int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v257 = v253 * int32(52)
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257)+uint32(_c_F_convert_case[3]))))
	switch v260 {
	case 0:
		goto L38
	case 1:
		goto L106
	default:
		goto L103
	}
L106:
	;
	if v47 == int32(0) {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v264 = v47 - int32(1)
	if v264 < int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	if l3 == v47 {
		goto L38
	} else {
		goto L178
	}
L109:
	;
	__phi277 = v264
	__phi285 = v47
	v277 = __phi277
	v285 = __phi285
	goto L110
L110:
	;
	v288 = l2 + v277
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	v290 = base.I32_extend8_s(v289)
	if int32(-64) <= v290 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	if base.Ui32(int32(127)) < base.Ui32(v351) {
		goto L144
	} else {
		goto L145
	}
L112:
	;
	goto L111
L113:
	;
	if int32(0) <= v290 {
		v351 = v289
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	if int32(0) < v277 {
		__phi277 = v277 - int32(1)
		__phi285 = v277
		v277 = __phi277
		v285 = __phi285
		goto L110
	} else {
		goto L141
	}
L116:
	;
	if base.Ui32(int32(127)) < base.Ui32(v351) {
		goto L127
	} else {
		goto L128
	}
L117:
	;
	if v289&int32(224) == int32(192) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v288))))
	v351 = v346&int32(63) | v343
	goto L116
L119:
	;
	v343 = v289 << (uint(int32(6)) % 32) & int32(1984)
	v344 = int32(1)
	goto L118
L120:
	;
	goto L121
L121:
	;
	if v289&int32(240) == int32(224) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v285))))
	v343 = v289<<(uint(int32(12))%32)&int32(_a_F_convert_case_0) | v313&int32(63)<<(uint(int32(6))%32)
	v344 = int32(2)
	goto L118
L123:
	;
	goto L124
L124:
	;
	if v289&int32(248) != int32(240) {
		v351 = int32(-1)
		goto L116
	} else {
		goto L125
	}
L125:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v285))))
	v331 = int32(63)
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+2)))
	v343 = v289<<(uint(int32(18))%32)&int32(_a_F_convert_case_1) | v330&v331<<(uint(int32(12))%32) | v336&v331<<(uint(int32(6))%32)
	v344 = int32(3)
	goto L118
L126:
	;
	if v400 == int32(0) {
		goto L112
	} else {
		goto L140
	}
L127:
	;
	v360 = int32(0)
	v361 = int32(505)
	goto L130
L128:
	;
	goto L129
L129:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351<<(uint(int32(1))%32))+uint32(_c_F_convert_case[4]))))
	v400 = int32(base.Ui32(v390&int32(16)) >> (uint(int32(4)) % 32))
	goto L126
L130:
	;
	v366 = base.I32_div_s(v360+v361, int32(2))
	v368 = v366 << (uint(int32(3)) % 32)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_convert_case[5])))
	if base.Ui32(v371) < base.Ui32(v351) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v400 = int32(0)
	goto L126
L132:
	;
	if v382 <= v383 {
		v360 = v382
		v361 = v383
		goto L130
	} else {
		goto L139
	}
L133:
	;
	v382 = v366 + int32(1)
	v383 = v361
	goto L132
L134:
	;
	goto L135
L135:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v368)+uint32(_c_F_convert_case[6])))
	if base.Ui32(v377) <= base.Ui32(v351) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v400 = int32(1)
	goto L126
L137:
	;
	goto L138
L138:
	;
	v382 = v360
	v383 = v366 - int32(1)
	goto L132
L139:
	;
	goto L131
L140:
	;
	goto L115
L141:
	;
	goto L108
L142:
	;
	if v518 == int32(0) {
		goto L103
	} else {
		goto L177
	}
L143:
	;
	v459 = int32(689)
	v460 = int32(0)
	goto L157
L144:
	;
	v417 = int32(3367)
	v418 = int32(0)
	goto L147
L145:
	;
	goto L146
L146:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351<<(uint(int32(1))%32))+uint32(_c_F_convert_case[4]))))
	v518 = int32(base.Ui32(v449&int32(8)) >> (uint(int32(3)) % 32))
	goto L142
L147:
	;
	v423 = base.I32_div_s(v417+v418, int32(2))
	v425 = v423 * int32(12)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v425)+uint32(_c_F_convert_case[7])))
	if base.Ui32(v428) < base.Ui32(v351) {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425)+uint32(_c_F_convert_case[8]))))
	if v441 != int32(3) {
		goto L143
	} else {
		goto L156
	}
L149:
	;
	goto L148
L150:
	;
	if v439 <= v438 {
		v417 = v438
		v418 = v439
		goto L147
	} else {
		goto L155
	}
L151:
	;
	v438 = v417
	v439 = v423 + int32(1)
	goto L150
L152:
	;
	goto L153
L153:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v425)+uint32(_c_F_convert_case[9])))
	if base.Ui32(v434) <= base.Ui32(v351) {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	v438 = v423 - int32(1)
	v439 = v418
	goto L150
L155:
	;
	goto L143
L156:
	;
	v518 = int32(1)
	goto L142
L157:
	;
	v465 = base.I32_div_s(v459+v460, int32(2))
	v467 = v465 << (uint(int32(3)) % 32)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_convert_case[10])))
	if base.Ui32(v470) < base.Ui32(v351) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v487 = int32(655)
	v488 = int32(0)
	goto L167
L159:
	;
	if v482 <= v481 {
		v459 = v481
		v460 = v482
		goto L157
	} else {
		goto L166
	}
L160:
	;
	v481 = v459
	v482 = v465 + int32(1)
	goto L159
L161:
	;
	goto L162
L162:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_convert_case[11])))
	if base.Ui32(v476) <= base.Ui32(v351) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v518 = int32(1)
	goto L142
L164:
	;
	goto L165
L165:
	;
	v481 = v465 - int32(1)
	v482 = v460
	goto L159
L166:
	;
	goto L158
L167:
	;
	v493 = base.I32_div_s(v487+v488, int32(2))
	v495 = v493 << (uint(int32(3)) % 32)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v495)+uint32(_c_F_convert_case[12])))
	if base.Ui32(v498) < base.Ui32(v351) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v518 = int32(0)
	goto L142
L169:
	;
	if v510 <= v509 {
		v487 = v509
		v488 = v510
		goto L167
	} else {
		goto L176
	}
L170:
	;
	v509 = v487
	v510 = v493 + int32(1)
	goto L169
L171:
	;
	goto L172
L172:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v495)+uint32(_c_F_convert_case[13])))
	if base.Ui32(v504) <= base.Ui32(v351) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v518 = int32(1)
	goto L142
L174:
	;
	goto L175
L175:
	;
	v509 = v493 - int32(1)
	v510 = v488
	goto L169
L176:
	;
	goto L168
L177:
	;
	goto L108
L178:
	;
	v544 = v47 + int32(1)
	if base.Ui32(l3) <= base.Ui32(v544) {
		goto L38
	} else {
		goto L179
	}
L179:
	;
	v554 = v544
	goto L180
L180:
	;
	v567 = l2 + v554
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567))))
	if v568 == int32(0) {
		goto L38
	} else {
		goto L182
	}
L181:
	;
	if base.Ui32(int32(127)) < base.Ui32(v630) {
		goto L215
	} else {
		goto L216
	}
L182:
	;
	v571 = base.I32_extend8_s(v568)
	if int32(-64) <= v571 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	goto L181
L184:
	;
	if int32(0) <= v571 {
		v630 = v568
		goto L187
	} else {
		goto L188
	}
L185:
	;
	goto L186
L186:
	;
	v685 = v554 + int32(1)
	if v685 != l3 {
		v554 = v685
		goto L180
	} else {
		goto L212
	}
L187:
	;
	if base.Ui32(int32(127)) < base.Ui32(v630) {
		goto L198
	} else {
		goto L199
	}
L188:
	;
	if v568&int32(224) == int32(192) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623+v567))))
	v630 = v625&int32(63) | v622
	goto L187
L190:
	;
	v622 = v568 << (uint(int32(6)) % 32) & int32(1984)
	v623 = int32(1)
	goto L189
L191:
	;
	goto L192
L192:
	;
	if v568&int32(240) == int32(224) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+1)))
	v622 = v568<<(uint(int32(12))%32)&int32(_a_F_convert_case_0) | v593&int32(63)<<(uint(int32(6))%32)
	v623 = int32(2)
	goto L189
L194:
	;
	goto L195
L195:
	;
	if v568&int32(248) != int32(240) {
		v630 = int32(-1)
		goto L187
	} else {
		goto L196
	}
L196:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+1)))
	v610 = int32(63)
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+2)))
	v622 = v568<<(uint(int32(18))%32)&int32(_a_F_convert_case_1) | v609&v610<<(uint(int32(12))%32) | v615&v610<<(uint(int32(6))%32)
	v623 = int32(3)
	goto L189
L197:
	;
	if v679 == int32(0) {
		goto L183
	} else {
		goto L211
	}
L198:
	;
	v639 = int32(0)
	v640 = int32(505)
	goto L201
L199:
	;
	goto L200
L200:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630<<(uint(int32(1))%32))+uint32(_c_F_convert_case[4]))))
	v679 = int32(base.Ui32(v669&int32(16)) >> (uint(int32(4)) % 32))
	goto L197
L201:
	;
	v645 = base.I32_div_s(v639+v640, int32(2))
	v647 = v645 << (uint(int32(3)) % 32)
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v647)+uint32(_c_F_convert_case[5])))
	if base.Ui32(v650) < base.Ui32(v630) {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	v679 = int32(0)
	goto L197
L203:
	;
	if v661 <= v662 {
		v639 = v661
		v640 = v662
		goto L201
	} else {
		goto L210
	}
L204:
	;
	v661 = v645 + int32(1)
	v662 = v640
	goto L203
L205:
	;
	goto L206
L206:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v647)+uint32(_c_F_convert_case[6])))
	if base.Ui32(v656) <= base.Ui32(v630) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v679 = int32(1)
	goto L197
L208:
	;
	goto L209
L209:
	;
	v661 = v639
	v662 = v645 - int32(1)
	goto L203
L210:
	;
	goto L202
L211:
	;
	goto L186
L212:
	;
	goto L38
L213:
	;
	if v796 == int32(0) {
		goto L38
	} else {
		goto L248
	}
L214:
	;
	v737 = int32(689)
	v738 = int32(0)
	goto L228
L215:
	;
	v695 = int32(3367)
	v696 = int32(0)
	goto L218
L216:
	;
	goto L217
L217:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630<<(uint(int32(1))%32))+uint32(_c_F_convert_case[4]))))
	v796 = int32(base.Ui32(v727&int32(8)) >> (uint(int32(3)) % 32))
	goto L213
L218:
	;
	v701 = base.I32_div_s(v695+v696, int32(2))
	v703 = v701 * int32(12)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v703)+uint32(_c_F_convert_case[7])))
	if base.Ui32(v706) < base.Ui32(v630) {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+uint32(_c_F_convert_case[8]))))
	if v719 != int32(3) {
		goto L214
	} else {
		goto L227
	}
L220:
	;
	goto L219
L221:
	;
	if v717 <= v716 {
		v695 = v716
		v696 = v717
		goto L218
	} else {
		goto L226
	}
L222:
	;
	v716 = v695
	v717 = v701 + int32(1)
	goto L221
L223:
	;
	goto L224
L224:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v703)+uint32(_c_F_convert_case[9])))
	if base.Ui32(v712) <= base.Ui32(v630) {
		goto L220
	} else {
		goto L225
	}
L225:
	;
	v716 = v701 - int32(1)
	v717 = v696
	goto L221
L226:
	;
	goto L214
L227:
	;
	v796 = int32(1)
	goto L213
L228:
	;
	v743 = base.I32_div_s(v737+v738, int32(2))
	v745 = v743 << (uint(int32(3)) % 32)
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v745)+uint32(_c_F_convert_case[10])))
	if base.Ui32(v748) < base.Ui32(v630) {
		goto L231
	} else {
		goto L232
	}
L229:
	;
	v765 = int32(655)
	v766 = int32(0)
	goto L238
L230:
	;
	if v760 <= v759 {
		v737 = v759
		v738 = v760
		goto L228
	} else {
		goto L237
	}
L231:
	;
	v759 = v737
	v760 = v743 + int32(1)
	goto L230
L232:
	;
	goto L233
L233:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v745)+uint32(_c_F_convert_case[11])))
	if base.Ui32(v754) <= base.Ui32(v630) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v796 = int32(1)
	goto L213
L235:
	;
	goto L236
L236:
	;
	v759 = v743 - int32(1)
	v760 = v738
	goto L230
L237:
	;
	goto L229
L238:
	;
	v771 = base.I32_div_s(v765+v766, int32(2))
	v773 = v771 << (uint(int32(3)) % 32)
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v773)+uint32(_c_F_convert_case[12])))
	if base.Ui32(v776) < base.Ui32(v630) {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	v796 = int32(0)
	goto L213
L240:
	;
	if v788 <= v787 {
		v765 = v787
		v766 = v788
		goto L238
	} else {
		goto L247
	}
L241:
	;
	v787 = v765
	v788 = v771 + int32(1)
	goto L240
L242:
	;
	goto L243
L243:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v773)+uint32(_c_F_convert_case[13])))
	if base.Ui32(v782) <= base.Ui32(v630) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v796 = int32(1)
	goto L213
L245:
	;
	goto L246
L246:
	;
	v787 = v771 - int32(1)
	v788 = v766
	goto L240
L247:
	;
	goto L239
L248:
	;
	goto L103
L249:
	;
	if base.Ui32(v850) < base.Ui32(int32(_a_F_convert_case_2)) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v858 = int32(3)
	goto L252
L251:
	;
	v858 = int32(4)
	goto L252
L252:
	;
	if base.Ui32(v850) < base.Ui32(int32(2048)) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v861 = int32(2)
	goto L255
L254:
	;
	v861 = v858
	goto L255
L255:
	;
	v862 = v861 + v43
	if base.Ui32(v862) <= base.Ui32(l1) {
		goto L39
	} else {
		goto L256
	}
L256:
	;
	v1079 = v862
	goto L37
L257:
	;
	v1079 = v864
	goto L37
L258:
	;
	goto L259
L259:
	;
	if v133 != 0 {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	v1079 = v864
	goto L37
L261:
	;
	v867 = F__emscripten_memcpy_bulkmem(m, l0+v43, v55, v133)
	mBase = m.M
	goto L263
L262:
	;
	goto L263
L263:
	;
	goto L260
L264:
	;
	v1079 = v870
	goto L37
L265:
	;
	goto L266
L266:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v43))) = uint8(v850)
	v1079 = v870
	goto L37
L267:
	;
	v880 = v850&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+1)) = uint8(v880)
	v885 = int32(base.Ui32(v850)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v874))) = uint8(v885)
	v1079 = v862
	goto L37
L268:
	;
	goto L269
L269:
	;
	if base.Ui32(v850) <= base.Ui32(int32(_a_F_convert_case_41)) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v889 = int32(63)
	v891 = int32(128)
	v892 = v850&v889 | v891
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+2)) = uint8(v892)
	v897 = int32(base.Ui32(v850)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v874))) = uint8(v897)
	v904 = int32(base.Ui32(v850)>>(uint(int32(6))%32))&v889 | v891
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+1)) = uint8(v904)
	v1079 = v862
	goto L37
L271:
	;
	goto L272
L272:
	;
	v906 = int32(63)
	v908 = int32(128)
	v909 = v850&v906 | v908
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+3)) = uint8(v909)
	v916 = int32(base.Ui32(v850)>>(uint(int32(6))%32))&v906 | v908
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+2)) = uint8(v916)
	v923 = int32(base.Ui32(v850)>>(uint(int32(12))%32))&v906 | v908
	*(*uint8)(unsafe.Add(mBase, uint32(v874)+1)) = uint8(v923)
	v930 = int32(base.Ui32(v850)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v874))) = uint8(v930)
	v1079 = v862
	goto L37
L273:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v257+v140*int32(12)+int32(_a_F_convert_case_42)+v970<<(uint(int32(2))%32))))
	if v983 == int32(0) {
		v1079 = v968
		goto L37
	} else {
		goto L275
	}
L274:
	;
	v1079 = v1065
	goto L37
L275:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v983) {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	v1067 = v970 + int32(1)
	if v1067 != int32(3) {
		v968 = v1065
		v970 = v1067
		goto L273
	} else {
		goto L296
	}
L277:
	;
	v1065 = v997
	goto L276
L278:
	;
	if base.Ui32(v983) <= base.Ui32(int32(_a_F_convert_case_41)) {
		goto L293
	} else {
		goto L294
	}
L279:
	;
	if base.Ui32(v983) < base.Ui32(int32(_a_F_convert_case_2)) {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	goto L281
L281:
	;
	v1013 = v968 + int32(1)
	if base.Ui32(l1) < base.Ui32(v1013) {
		goto L290
	} else {
		goto L291
	}
L282:
	;
	v993 = int32(3)
	goto L284
L283:
	;
	v993 = int32(4)
	goto L284
L284:
	;
	if base.Ui32(v983) < base.Ui32(int32(2048)) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v996 = int32(2)
	goto L287
L286:
	;
	v996 = v993
	goto L287
L287:
	;
	v997 = v996 + v968
	if base.Ui32(l1) < base.Ui32(v997) {
		goto L277
	} else {
		goto L288
	}
L288:
	;
	v999 = l0 + v968
	if base.Ui32(int32(2047)) < base.Ui32(v983) {
		goto L278
	} else {
		goto L289
	}
L289:
	;
	v1005 = v983&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+1)) = uint8(v1005)
	v1010 = int32(base.Ui32(v983)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v999))) = uint8(v1010)
	goto L277
L290:
	;
	v1065 = v1013
	goto L276
L291:
	;
	goto L292
L292:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v968))) = uint8(v983)
	v1065 = v1013
	goto L276
L293:
	;
	v1019 = int32(63)
	v1021 = int32(128)
	v1022 = v983&v1019 | v1021
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+2)) = uint8(v1022)
	v1027 = int32(base.Ui32(v983)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v999))) = uint8(v1027)
	v1034 = int32(base.Ui32(v983)>>(uint(int32(6))%32))&v1019 | v1021
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+1)) = uint8(v1034)
	goto L277
L294:
	;
	goto L295
L295:
	;
	v1036 = int32(63)
	v1038 = int32(128)
	v1039 = v983&v1036 | v1038
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+3)) = uint8(v1039)
	v1046 = int32(base.Ui32(v983)>>(uint(int32(6))%32))&v1036 | v1038
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+2)) = uint8(v1046)
	v1053 = int32(base.Ui32(v983)>>(uint(int32(12))%32))&v1036 | v1038
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+1)) = uint8(v1053)
	v1060 = int32(base.Ui32(v983)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v999))) = uint8(v1060)
	goto L277
L296:
	;
	goto L274
L297:
	;
	if base.Ui32(v1091) < base.Ui32(l3) {
		v43 = v1079
		v47 = v1091
		v49 = v140
		v50 = v141
		goto L11
	} else {
		goto L298
	}
L298:
	;
	goto L12
L299:
	;
	v1118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v1104))) = uint8(v1118)
	goto L301
L300:
	;
	goto L301
L301:
	;
	return v1104
}
func F_copy_dest_receive(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+172))
	F_MemoryContextReset(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(_a_F_copy_dest_receive_0)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0]))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+172))
		*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0])) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
		if v20 < v19 {
			F_slot_getsomeattrs_int(m, l0, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
				m.T0[v25].(func(*base.Module, int32, int32))(m, v7, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0])) = v14
					v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
					v32 = v30 + int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v32
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[1]))
					if v37 == int32(0) {
					} else {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_copy_dest_receive[2])))
						if v41 != int32(1) {
						} else {
							v44 = int32(_a_F_copy_dest_receive_1)
							v46 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
							v47 = int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v46 + v47
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
							*(*int32)(unsafe.Add(mBase, uint32(v37))) = v50 + v47
							*(*int64)(unsafe.Add(mBase, uint32(v37+int32(16))+232)) = v32
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
							*(*int32)(unsafe.Add(mBase, uint32(v37))) = v58 + v47
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
							*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v64 - v47
						}
					}
					return int32(1)
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
			m.T0[v25].(func(*base.Module, int32, int32))(m, v7, l0)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0])) = v14
				v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v32 = v30 + int64(1)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v32
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[1]))
				if v37 == int32(0) {
				} else {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_copy_dest_receive[2])))
					if v41 != int32(1) {
					} else {
						v44 = int32(_a_F_copy_dest_receive_1)
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
						v47 = int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v46 + v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						*(*int32)(unsafe.Add(mBase, uint32(v37))) = v50 + v47
						*(*int64)(unsafe.Add(mBase, uint32(v37+int32(16))+232)) = v32
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						*(*int32)(unsafe.Add(mBase, uint32(v37))) = v58 + v47
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v64 - v47
					}
				}
				return int32(1)
			}
		}
	}
}
func F_copy_intArrayType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = F_ArrayGetNItems(m, v5, l0+int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if int32(0) < v8 {
			v17 = v8<<(uint(int32(2))%32) + int32(24)
			v18 = F_palloc0(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v8
				*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(23)
				*(*int64)(unsafe.Add(mBase, uint32(v18)+4)) = int64(1)
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v17 << (uint(int32(2)) % 32)
				v34 = v18
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
				v43 = v34
				v44 = (v36<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v46 != 0 {
					v54 = v46
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v54 = (v47<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				v57 = v8 << (uint(int32(2)) % 32)
				if v57 != 0 {
					v58 = F__emscripten_memcpy_bulkmem(m, v43+v44, v54+l0, v57)
					mBase = m.M
				} else {
				}
				return v43
			}
		} else {
			v31 = F_construct_empty_array(m, int32(23))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				if v33 != 0 {
					v43 = v31
					v44 = v33
				} else {
					v34 = v31
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
					v43 = v34
					v44 = (v36<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v46 != 0 {
					v54 = v46
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v54 = (v47<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				v57 = v8 << (uint(int32(2)) % 32)
				if v57 != 0 {
					v58 = F__emscripten_memcpy_bulkmem(m, v43+v44, v54+l0, v57)
					mBase = m.M
				} else {
				}
				return v43
			}
		}
	}
}
func F_copy_pathtarget(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v6 = F_palloc0(m, int32(40))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(277)
		v13 = l0 + int32(8)
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v16
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v18
		v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v20
		v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v22
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v25 = F_list_copy(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v25
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			if v28 != 0 {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v29 != 0 {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					v34 = v30 << (uint(int32(2)) % 32)
				} else {
					v34 = int32(0)
				}
				v35 = F_palloc(m, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v35
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v34 != 0 {
						v39 = F__emscripten_memcpy_bulkmem(m, v35, v38, v34)
						mBase = m.M
					} else {
					}
					return v6
				}
			} else {
				return v6
			}
		}
	}
}
func F_copysignl(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = l1
	v10 = int64(48)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = l2&int64(281474976710655) | base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(l2&int64(9223090561878065152))>>(uint(v10)%64)))|base.I32_wrap_i64(int64(base.Ui64(l3)>>(uint(v10)%64)))&int32(_a_F_copysignl_0))<<(uint(v10)%64)
	return
}
func F_cost_resultscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 float64
	_ = v15
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v62 float64
	_ = v62
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v81 float64
	_ = v81
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	v7 = float64(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v86 = *(*float64)(unsafe.Add(mBase, uint32(l2)+120))
	v88 = *(*float64)(unsafe.Add(mBase, _c_F_cost_resultscan[0]))
	v89 = float64(0)
	v90 = base.F64_add(v85, v89)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v90, base.F64_add(base.F64_mul(v86, base.F64_add(v81, v88)), v89))
	m.G0 = v13 + int32(32)
	return
L2:
	;
	v15 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l1
	if v17 == int32(0) {
		v62 = v7
		v66 = float64(0)
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v71 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v71
	v73 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v74 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v81 = v73
	v85 = v74
	goto L1
L5:
	;
	v67 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v69 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v81 = base.F64_add(v62, v67)
	v85 = base.F64_add(v66, v69)
	goto L1
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v27 <= int32(0) {
		v62 = v7
		v66 = float64(0)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v34 = int32(0)
	goto L8
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v34<<(uint(int32(2))%32))))
	v48 = F_cost_qual_eval_walker(m, v45, v13+int32(8))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v13)+24))
	v55 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
	v62 = v54
	v66 = v55
	goto L5
L10:
	;
	return
L11:
	;
	v51 = v34 + int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v51 < v52 {
		v34 = v51
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
}
func F_cost_seqscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v82 float64
	_ = v82
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v107 float64
	_ = v107
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v116 float64
	_ = v116
	var v118 float64
	_ = v118
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v122 int32
	_ = v122
	var v125 float64
	_ = v125
	var v127 int32
	_ = v127
	var v133 float64
	_ = v133
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v150 float64
	_ = v150
	var v154 float64
	_ = v154
	var v157 float64
	_ = v157
	var v162 int32
	_ = v162
	var v165 float64
	_ = v165
	v8 = float64(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = l3 + int32(8)
	goto L3
L2:
	;
	v23 = l2 + int32(16)
	goto L3
L3:
	;
	v24 = *(*float64)(unsafe.Add(mBase, uint32(v23)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	F_get_tablespace_page_costs(m, v26, int32(0), v17)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
	v31 = *(*float64)(unsafe.Add(mBase, uint32(v17)))
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v112 = *(*float64)(unsafe.Add(mBase, uint32(v111)+24))
	v113 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v116 = *(*float64)(unsafe.Add(mBase, _c_F_cost_seqscan[0]))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(l2)+120))
	v120 = base.F64_add(base.F64_mul(v112, v113), base.F64_mul(base.F64_add(v110, v116), v118))
	v121 = *(*float64)(unsafe.Add(mBase, uint32(v111)+16))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(0) < v122 {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v33 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l1
	if v32 == int32(0) {
		v82 = v8
		v89 = float64(0)
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v107 = v94
	v110 = v95
	goto L6
L10:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v92 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v107 = base.F64_add(v89, v90)
	v110 = base.F64_add(v82, v92)
	goto L6
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v42 <= int32(0) {
		v82 = v8
		v89 = float64(0)
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v49 = int32(0)
	goto L13
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v49<<(uint(int32(2))%32))))
	v67 = F_cost_qual_eval_walker(m, v64, v17+int32(8))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
	v82 = v73
	v89 = v74
	goto L10
L15:
	;
	v70 = v49 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v70 < v71 {
		v49 = v70
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v125 = base.F64_convert_i32_u(v122)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_seqscan[1])))
	if v127 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v157 = v120
	goto L19
L19:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_seqscan[2])))
	v165 = base.F64_add(base.F64_add(v107, float64(0)), v121)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v162 ^ int32(1)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_mul(v31, base.F64_convert_i32_u(v30)), base.F64_add(v165, v157))
	m.G0 = v17 + int32(32)
	return
L20:
	;
	v133 = base.F64_add(base.F64_mul(v125, float64(-0.3)), float64(1))
	if base.F64_gt(v133, float64(0)) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v139 = v125
	goto L22
L22:
	;
	v141 = float64(1e+100)
	v142 = base.F64_div(v113, v139)
	if base.F64_gt(v142, v141) != 0 {
		v154 = v141
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v137 = v133
	goto L25
L24:
	;
	v137 = math.Float64frombits(uint64(0x8000000000000000))
	goto L25
L25:
	;
	v139 = base.F64_add(v137, v125)
	goto L22
L26:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v154
	v157 = base.F64_div(v120, v139)
	goto L19
L27:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v142)&int64(9223372036854775807)) {
		v154 = v141
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v150 = float64(1)
	if base.F64_le(v142, v150) != 0 {
		v154 = v150
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v154 = base.F64_nearest(v142)
	goto L26
}
func F_count_nulls(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v214 int32
	_ = v214
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == v4 {
		v24 = v4
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
		if v16 == int32(0) {
			v24 = v4
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v19 != int32(15) {
				v24 = v4
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+13)))
				v24 = v22
			}
		}
	}
	if v24&int32(1) == int32(0) {
		v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v29 <= int32(0) {
			v192 = v4
			v195 = v29
		} else {
			v33 = v29 & int32(3)
			v35 = l0 + int32(24)
			v36 = int32(0)
			if base.Ui32(int32(4)) <= base.Ui32(v29) {
				v41 = v36
				v44 = v4
				v49 = v4
				for {
					v54 = v35 + v41<<(uint(int32(3))%32)
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
					v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+8)))
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+16)))
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+24)))
					v62 = v44 + v55 + v57 + v59 + v61
					v63 = int32(4)
					v64 = v41 + v63
					v66 = v49 + v63
					if v66 != v29&int32(_a_F_count_nulls_0) {
						v41 = v64
						v44 = v62
						v49 = v66
						continue
					} else {
						break
					}
					break
				}
				v68 = v64
				v71 = v62
			} else {
				v68 = v36
				v71 = v4
			}
			if v33 == int32(0) {
				v192 = v71
				v195 = v29
			} else {
				v81 = v68
				v84 = v71
				v88 = v4
				for {
					v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v81<<(uint(int32(3))%32)))))
					v96 = v84 + v95
					v97 = int32(1)
					v100 = v88 + v97
					if v100 != v33 {
						v81 = v81 + v97
						v84 = v96
						v88 = v100
						continue
					} else {
						break
					}
					break
				}
				v192 = v96
				v195 = v29
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v195
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v192
		v214 = int32(1)
		return v214
	} else {
		v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v103 != 0 {
			v214 = int32(0)
			return v214
		} else {
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v105 = F_pg_detoast_datum(m, v104)
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return int32(0)
			} else {
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
				v111 = v105 + int32(16)
				v112 = F_ArrayGetNItems(m, v109, v111)
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return int32(0)
				} else {
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
					if v114 == int32(0) {
						v192 = v4
						v195 = v112
					} else {
						v117 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
						v120 = v111 + v117<<(uint(int32(3))%32)
						if v120 == int32(0) {
							v192 = v4
							v195 = v112
						} else {
							if v112 <= int32(0) {
								v192 = v4
								v195 = v112
							} else {
								v126 = int32(1)
								if v112 == v126 {
									v174 = v4
									v176 = v120
									v179 = int32(1)
								} else {
									v133 = int32(1)
									v136 = v4
									v137 = v120
									v140 = v4
									for {
										v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
										v149 = int32(1)
										v151 = v133 << (uint(v149) % 32)
										v153 = base.B2i32(v151 == int32(256))
										if v151 == int32(256) {
											v154 = v149
										} else {
											v154 = v151
										}
										v155 = v153 + v137
										v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
										v160 = v136 + base.B2i32(v133&v144 == int32(0)) + base.B2i32(v154&v156 == int32(0))
										v161 = int32(1)
										v163 = v154 << (uint(v161) % 32)
										v165 = base.B2i32(v163 == int32(256))
										if v163 == int32(256) {
											v166 = v161
										} else {
											v166 = v163
										}
										v167 = v155 + v165
										v169 = v140 + int32(2)
										if v169 != v112&int32(2147483646) {
											v133 = v166
											v136 = v160
											v137 = v167
											v140 = v169
											continue
										} else {
											break
										}
										break
									}
									v174 = v160
									v176 = v167
									v179 = v166
								}
								if v112&v126 == int32(0) {
									v192 = v174
									v195 = v112
								} else {
									v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
									v192 = v174 + base.B2i32(v179&v184 == int32(0))
									v195 = v112
								}
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v195
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v192
					v214 = int32(1)
					return v214
				}
			}
		}
	}
}
