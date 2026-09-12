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
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 float64
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v315 int32
	_ = v315
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v468 int32
	_ = v468
	var v479 int32
	_ = v479
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v495 float64
	_ = v495
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 float64
	_ = v509
	var v511 float32
	_ = v511
	var v514 float64
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 float64
	_ = v620
	var v624 float64
	_ = v624
	var v631 float64
	_ = v631
	var v632 float64
	_ = v632
	var v634 float64
	_ = v634
	var v640 float64
	_ = v640
	var v641 float64
	_ = v641
	var v644 float64
	_ = v644
	var v646 float64
	_ = v646
	var v679 float32
	_ = v679
	var v681 float64
	_ = v681
	var v687 float32
	_ = v687
	var v689 float32
	_ = v689
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v881 float32
	_ = v881
	var v882 float64
	_ = v882
	var v883 float32
	_ = v883
	var v885 float64
	_ = v885
	var v891 int32
	_ = v891
	var v897 float64
	_ = v897
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v923 int32
	_ = v923
	var v925 float64
	_ = v925
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v951 float64
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v961 int32
	_ = v961
	var v963 float64
	_ = v963
	var v981 int32
	_ = v981
	var v983 float64
	_ = v983
	var v991 int32
	_ = v991
	var v998 int32
	_ = v998
	var v1000 float64
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1012 float64
	_ = v1012
	var v1024 float64
	_ = v1024
	var v1033 int32
	_ = v1033
	var v1038 float64
	_ = v1038
	var v1050 float64
	_ = v1050
	var v1051 float64
	_ = v1051
	var v1055 float64
	_ = v1055
	var v1058 float64
	_ = v1058
	var v1061 float64
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1065 float64
	_ = v1065
	var v1069 float64
	_ = v1069
	var v1074 int32
	_ = v1074
	var v1075 float64
	_ = v1075
	var v1077 float64
	_ = v1077
	var v1083 float64
	_ = v1083
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1120 int32
	_ = v1120
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
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
	if int32(0) < v488 {
		goto L91
	} else {
		goto L92
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
	v502 = v78 + int32(1)
	if v502 != l2 {
		v72 = v479
		v78 = v502
		v80 = v487
		v81 = v488
		v83 = v490
		v88 = v495
		goto L12
	} else {
		goto L88
	}
L17:
	;
	v479 = v72
	v487 = v80 + int32(1)
	v488 = v81
	v490 = v83
	v495 = v88
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
	v212 = int32(0)
	if v212 < v72 {
		goto L60
	} else {
		goto L61
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
		v210 = v99
		v211 = v88
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
	v479 = v72
	v487 = v80
	v488 = v107
	v490 = v83 + int32(1)
	v495 = v136
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
	v210 = v143
	v211 = v136
	goto L20
L38:
	;
	if v99&int32(3) == int32(0) {
		v170 = v99
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v210 = v99
	v211 = base.F64_add(v88, base.F64_convert_i32_u(v203+int32(1)))
	goto L20
L40:
	;
	v203 = v195 - v99
	goto L39
L41:
	;
	v174 = v170
	goto L50
L42:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v154 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v203 = int32(0)
	goto L39
L44:
	;
	goto L45
L45:
	;
	v159 = v99
	goto L46
L46:
	;
	v163 = v159 + int32(1)
	if v163&int32(3) == int32(0) {
		v170 = v163
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v195 = v163
	goto L40
L48:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v168 != 0 {
		v159 = v163
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v183 = int32(-2139062144)
	if (int32(16843008)-v180|v180)&v183 == v183 {
		v174 = v174 + int32(4)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v189 = v174
	goto L53
L52:
	;
	goto L51
L53:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v193 != 0 {
		v189 = v189 + int32(1)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v195 = v189
	goto L40
L55:
	;
	goto L54
L56:
	;
	if v266 < v290 {
		goto L85
	} else {
		goto L86
	}
L57:
	;
	if v266 == v289+v72-int32(2) {
		goto L56
	} else {
		goto L81
	}
L58:
	;
	v359 = int32(3)
	v361 = v57 + v292<<(uint(v359)%32)
	v364 = v57 + v290<<(uint(v359)%32)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v364-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = v367
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v364-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+4)) = v371
	v375 = v290 - int32(2)
	v377 = v292
	goto L57
L59:
	;
	v304 = v248 + int32(4)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	*(*int32)(unsafe.Add(mBase, uint32(v304))) = v305 + int32(1)
	if v219 == int32(0) {
		v479 = v72
		v487 = v80
		v488 = v107
		v490 = v83
		v495 = v211
		goto L16
	} else {
		goto L76
	}
L60:
	;
	v219 = v212
	v220 = v72
	goto L63
L61:
	;
	v266 = v72
	goto L62
L62:
	;
	v289 = base.B2i32(v72 < v54)
	v290 = v72 + v289
	v292 = v290 - int32(1)
	if v292 <= v266 {
		goto L56
	} else {
		goto L74
	}
L63:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v248 = v57 + v219<<(uint(int32(3))%32)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	v250 = F_FunctionCall2Coll(m, v31+int32(4), v245, v210, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L7
	} else {
		goto L65
	}
L64:
	;
	v266 = v257
	goto L62
L65:
	;
	if v250 != 0 {
		goto L59
	} else {
		goto L66
	}
L66:
	;
	if v219 < v220 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if v253 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v257 = v220
	goto L69
L69:
	;
	v259 = v219 + int32(1)
	if v259 != v72 {
		v219 = v259
		v220 = v257
		goto L63
	} else {
		goto L73
	}
L70:
	;
	v256 = v219
	goto L72
L71:
	;
	v256 = v220
	goto L72
L72:
	;
	v257 = v256
	goto L69
L73:
	;
	goto L64
L74:
	;
	if (v72-v289+v266)&int32(1) == int32(0) {
		goto L58
	} else {
		goto L75
	}
L75:
	;
	v375 = v292
	v377 = v290
	goto L57
L76:
	;
	v315 = v219
	goto L77
L77:
	;
	v341 = v57 + v315<<(uint(int32(3))%32)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	v344 = v341 - int32(4)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	if v342 <= v345 {
		v479 = v72
		v487 = v80
		v488 = v107
		v490 = v83
		v495 = v211
		goto L16
	} else {
		goto L79
	}
L78:
	;
	v479 = v72
	v487 = v80
	v488 = v107
	v490 = v83
	v495 = v211
	goto L16
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341)+4)) = v345
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v350 = v341 - int32(8)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v350))) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v342
	v355 = int32(1)
	if v355 < v315 {
		v315 = v315 - v355
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v383 = v375
	v386 = v377
	goto L82
L82:
	;
	v407 = int32(3)
	v409 = v57 + v383<<(uint(v407)%32)
	v412 = v57 + v386<<(uint(v407)%32)
	v413 = int32(16)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v412-v413)))
	*(*int32)(unsafe.Add(mBase, uint32(v409))) = v415
	v417 = int32(12)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v412-v417)))
	*(*int32)(unsafe.Add(mBase, uint32(v409)+4)) = v419
	v422 = v383 - int32(1)
	v425 = v57 + v422<<(uint(v407)%32)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v409-v413)))
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = v428
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v409-v417)))
	*(*int32)(unsafe.Add(mBase, uint32(v425)+4)) = v432
	v435 = v383 - int32(2)
	if v266 < v435 {
		v383 = v435
		v386 = v422
		goto L82
	} else {
		goto L84
	}
L83:
	;
	goto L56
L84:
	;
	goto L83
L85:
	;
	v468 = v57 + v266<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v468)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v468))) = v210
	goto L87
L86:
	;
	goto L87
L87:
	;
	v479 = v290
	v487 = v80
	v488 = v107
	v490 = v83
	v495 = v211
	goto L16
L88:
	;
	goto L13
L89:
	;
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v679
	v681 = base.F64_promote_f32(v679)
	if base.F64_gt(v681, base.F64_mul(l3, float64(0.1))) != 0 {
		goto L128
	} else {
		goto L129
	}
L90:
	;
	if v54 <= v479 {
		goto L114
	} else {
		goto L115
	}
L91:
	;
	v506 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v506)
	v509 = base.F64_convert_i32_s(l2)
	v511 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v487), v509))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v511
	if v45 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	goto L93
L93:
	;
	if v487 <= int32(0) {
		goto L10
	} else {
		goto L110
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v523
	if int32(0) < v479 {
		goto L101
	} else {
		goto L102
	}
L95:
	;
	v514 = base.F64_div(v495, base.F64_convert_i32_u(v488))
	if base.F64_lt(base.F64_abs(v514), float64(2.147483648e+09)) != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v521 = int32(*(*int16)(unsafe.Add(mBase, uint32(v520)+76)))
	v523 = v521
	goto L94
L98:
	;
	v518 = base.I32_trunc_f64_s(v514)
	v523 = v518
	goto L94
L99:
	;
	goto L100
L100:
	;
	v523 = int32(-2147483648)
	goto L94
L101:
	;
	v527 = int32(0)
	v533 = v527
	v534 = v527
	goto L105
L102:
	;
	goto L103
L103:
	;
	v679 = base.F32_neg(base.F32_sub(float32(1), v511))
	goto L89
L104:
	;
	if v533 != 0 {
		v613 = v533
		v614 = v534
		goto L90
	} else {
		goto L109
	}
L105:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v57+v533<<(uint(int32(3))%32))+4))
	if v560 == int32(1) {
		goto L104
	} else {
		goto L107
	}
L106:
	;
	v613 = v479
	v614 = v563
	goto L90
L107:
	;
	v563 = v534 + v560
	v565 = v533 + int32(1)
	if v565 != v479 {
		v533 = v565
		v534 = v563
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	goto L103
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v602 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v602)
	v604 = int32(0)
	if v45 == v604 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v608 = int32(*(*int16)(unsafe.Add(mBase, uint32(v607)+76)))
	v609 = v608
	goto L113
L112:
	;
	v609 = v604
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v609
	goto L10
L114:
	;
	v618 = v488 - v614
	v619 = v618 + v613
	v620 = float64(0)
	v624 = base.F64_mul(l3, base.F64_sub(float64(1), base.F64_promote_f32(v511)))
	if base.F64_gt(v624, v620) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	if v490 != 0 {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	if v613 != v479 {
		goto L114
	} else {
		goto L117
	}
L117:
	;
	v679 = base.F32_convert_i32_s(v479)
	goto L89
L118:
	;
	if base.F64_lt(v640, v641) != 0 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v640 = v620
	v641 = base.F64_convert_i32_s(v619)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v631 = base.F64_convert_i32_s(l2 - v487)
	v632 = base.F64_convert_i32_s(v619)
	v634 = base.F64_convert_i32_s(v618)
	v640 = base.F64_div(base.F64_mul(v631, v632), base.F64_add(base.F64_sub(v631, v634), base.F64_div(base.F64_mul(v631, v634), v624)))
	v641 = v632
	goto L118
L122:
	;
	v644 = v641
	goto L124
L123:
	;
	v644 = v640
	goto L124
L124:
	;
	if base.F64_gt(v644, v624) != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v646 = v624
	goto L127
L126:
	;
	v646 = v644
	goto L127
L127:
	;
	v679 = base.F32_demote_f64(base.F64_floor(base.F64_add(v646, float64(0.5))))
	goto L89
L128:
	;
	v687 = base.F32_demote_f64(base.F64_div(base.F64_neg(v681), l3))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v687
	v689 = v687
	goto L130
L129:
	;
	v689 = v679
	goto L130
L130:
	;
	if v54 <= v479 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v1120 <= int32(0) {
		goto L10
	} else {
		goto L189
	}
L132:
	;
	if v49 < v479 {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	if v490 != 0 {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	if base.F32_gt(v689, float32(0)) == int32(0) {
		goto L132
	} else {
		goto L135
	}
L135:
	;
	if v479 <= v49 {
		v1120 = v479
		goto L131
	} else {
		goto L136
	}
L136:
	;
	goto L132
L137:
	;
	v697 = v49
	goto L139
L138:
	;
	v697 = v479
	goto L139
L139:
	;
	if v697 <= int32(0) {
		goto L10
	} else {
		goto L140
	}
L140:
	;
	v701 = v697 & int32(3)
	v705 = F_palloc(m, v697<<(uint(int32(2))%32))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L7
	} else {
		goto L141
	}
L141:
	;
	v707 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v697) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v717 = v707
	v721 = int32(0)
	goto L145
L143:
	;
	v788 = v707
	goto L144
L144:
	;
	if v701 != 0 {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v741 = int32(2)
	v744 = int32(3)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v57+v717<<(uint(v744)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v705+v717<<(uint(v741)%32)))) = v747
	v750 = v717 | int32(1)
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v57+v750<<(uint(v744)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v705+v750<<(uint(v741)%32)))) = v757
	v760 = v717 | v741
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v57+v760<<(uint(v744)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v705+v760<<(uint(v741)%32)))) = v767
	v770 = v717 | v744
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v57+v770<<(uint(v744)%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v705+v770<<(uint(v741)%32)))) = v777
	v779 = int32(4)
	v780 = v717 + v779
	v782 = v721 + v779
	if v782 != v697&int32(2147483644) {
		v717 = v780
		v721 = v782
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v788 = v780
	goto L144
L147:
	;
	goto L146
L148:
	;
	v816 = v788
	v818 = int32(0)
	goto L151
L149:
	;
	goto L150
L150:
	;
	v881 = *(*float32)(unsafe.Add(mBase, uint32(l0)+48))
	v882 = base.F64_promote_f32(v881)
	v883 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v885 = float64(0)
	v891 = int32(0)
	v897 = base.F64_convert_i32_s(l2)
	if base.F64_eq(l3, v897) != 0 {
		v1097 = v697
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v57+v816<<(uint(int32(3))%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v705+v816<<(uint(int32(2))%32)))) = v846
	v848 = int32(1)
	v851 = v818 + v848
	if v851 != v701 {
		v816 = v816 + v848
		v818 = v851
		goto L151
	} else {
		goto L153
	}
L152:
	;
	goto L150
L153:
	;
	goto L152
L154:
	;
	v1120 = v1097
	goto L131
L155:
	;
	goto L154
L156:
	;
	if base.F64_le(l3, float64(1)) != 0 {
		v1097 = v697
		goto L155
	} else {
		goto L157
	}
L157:
	;
	if base.Ui32(v697) < base.Ui32(int32(2)) {
		v1012 = v885
		goto L158
	} else {
		goto L159
	}
L158:
	;
	if base.F64_lt(v882, float64(0)) != 0 {
		goto L171
	} else {
		goto L172
	}
L159:
	;
	v908 = v697 - int32(1)
	v909 = int32(3)
	v910 = v908 & v909
	if base.Ui32(v697-int32(2)) < base.Ui32(v909) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if v910 == int32(0) {
		v1012 = v963
		goto L158
	} else {
		goto L167
	}
L161:
	;
	v961 = int32(0)
	v963 = v885
	goto L160
L162:
	;
	goto L163
L163:
	;
	v923 = int32(0)
	v925 = v885
	v934 = v891
	goto L164
L164:
	;
	v939 = v705 + v923<<(uint(int32(2))%32)
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v939)))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v939)+4))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v939)+8))
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v939)+12))
	v951 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v925, base.F64_convert_i32_s(v940)), base.F64_convert_i32_s(v943)), base.F64_convert_i32_s(v946)), base.F64_convert_i32_s(v949))
	v952 = int32(4)
	v953 = v923 + v952
	v955 = v934 + v952
	if v955 != v908&int32(-4) {
		v923 = v953
		v925 = v951
		v934 = v955
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v961 = v953
	v963 = v951
	goto L160
L166:
	;
	goto L165
L167:
	;
	v981 = v961
	v983 = v963
	v991 = v891
	goto L168
L168:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v705+v981<<(uint(int32(2))%32))))
	v1000 = base.F64_add(v983, base.F64_convert_i32_s(v998))
	v1001 = int32(1)
	v1004 = v991 + v1001
	if v1004 != v910 {
		v981 = v981 + v1001
		v983 = v1000
		v991 = v1004
		goto L168
	} else {
		goto L170
	}
L169:
	;
	v1012 = v1000
	goto L158
L170:
	;
	goto L169
L171:
	;
	v1024 = base.F64_mul(l3, base.F64_neg(v882))
	goto L173
L172:
	;
	v1024 = v882
	goto L173
L173:
	;
	v1033 = v697
	v1038 = v1012
	goto L174
L174:
	;
	v1050 = float64(1)
	v1051 = float64(0)
	v1055 = base.F64_sub(base.F64_sub(v1050, base.F64_div(v1038, v897)), base.F64_promote_f32(v883))
	if base.F64_lt(v1055, v1051) != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v1097 = int32(0)
	goto L155
L176:
	;
	v1058 = v1051
	goto L178
L177:
	;
	v1058 = v1055
	goto L178
L178:
	;
	if base.F64_gt(v1058, float64(1)) != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v1061 = v1050
	goto L181
L180:
	;
	v1061 = v1058
	goto L181
L181:
	;
	v1063 = v1033 - int32(1)
	v1065 = base.F64_sub(v1024, base.F64_convert_i32_u(v1063))
	if base.F64_gt(v1065, float64(1)) != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v1069 = base.F64_div(v1061, v1065)
	goto L184
L183:
	;
	v1069 = v1061
	goto L184
L184:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v705+v1063<<(uint(int32(2))%32))))
	v1075 = base.F64_convert_i32_s(v1074)
	v1077 = base.F64_div(base.F64_mul(l3, v1075), v897)
	v1083 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l3, v897), base.F64_mul(base.F64_mul(v1077, v897), base.F64_sub(l3, v1077))), base.F64_mul(base.F64_mul(l3, l3), base.F64_add(l3, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v1069, v897), base.F64_add(v1083, v1083)), float64(0.5)), v1075) != 0 {
		v1097 = v1033
		goto L155
	} else {
		goto L185
	}
L185:
	;
	if v1063 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v705-int32(8)+v1033<<(uint(int32(2))%32))))
	v1033 = v1063
	v1038 = base.F64_sub(v1038, base.F64_convert_i32_s(v1092))
	goto L174
L187:
	;
	goto L188
L188:
	;
	goto L175
L189:
	;
	v1145 = int32(_a_F_compute_distinct_stats_1)
	v1146 = *(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0]))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0])) = v1148
	v1151 = v1120 << (uint(int32(2)) % 32)
	v1152 = F_palloc(m, v1151)
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L7
	} else {
		goto L190
	}
L190:
	;
	v1154 = F_palloc(m, v1151)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L7
	} else {
		goto L191
	}
L191:
	;
	v1160 = int32(0)
	goto L192
L192:
	;
	v1185 = v1160 << (uint(int32(2)) % 32)
	v1189 = v57 + v1160<<(uint(int32(3))%32)
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1189)))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1191)+78)))
	v1193 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1191)+76)))
	v1194 = F_datumCopy(m, v1190, v1192, v1193)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L7
	} else {
		goto L194
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0])) = v1146
	v1208 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v1208)
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v1154
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v1152
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v1120
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v1120
	goto L10
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1152+v1185))) = v1194
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1189)+4))
	*(*float32)(unsafe.Add(mBase, uint32(v1185+v1154))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v1198), v509))
	v1204 = v1160 + int32(1)
	if v1204 != v1120 {
		v1160 = v1204
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
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
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v236 float64
	_ = v236
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 float64
	_ = v255
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v300 float64
	_ = v300
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var __phi344 int32
	_ = __phi344
	var v347 int32
	_ = v347
	var __phi347 int32
	_ = __phi347
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v396 int32
	_ = v396
	var v425 int32
	_ = v425
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v494 int32
	_ = v494
	var v495 float64
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 float64
	_ = v502
	var v504 float32
	_ = v504
	var v507 float64
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v525 int32
	_ = v525
	var v526 float64
	_ = v526
	var v530 float64
	_ = v530
	var v537 float64
	_ = v537
	var v538 float64
	_ = v538
	var v542 float64
	_ = v542
	var v548 float64
	_ = v548
	var v549 float64
	_ = v549
	var v552 float64
	_ = v552
	var v554 float64
	_ = v554
	var v564 float32
	_ = v564
	var v566 float64
	_ = v566
	var v572 float32
	_ = v572
	var v574 float32
	_ = v574
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v784 float32
	_ = v784
	var v785 float64
	_ = v785
	var v786 float32
	_ = v786
	var v788 float64
	_ = v788
	var v794 int32
	_ = v794
	var v800 float64
	_ = v800
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v826 int32
	_ = v826
	var v828 float64
	_ = v828
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v854 float64
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v864 int32
	_ = v864
	var v866 float64
	_ = v866
	var v884 int32
	_ = v884
	var v886 float64
	_ = v886
	var v894 int32
	_ = v894
	var v901 int32
	_ = v901
	var v903 float64
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v915 float64
	_ = v915
	var v927 float64
	_ = v927
	var v936 int32
	_ = v936
	var v941 float64
	_ = v941
	var v953 float64
	_ = v953
	var v954 float64
	_ = v954
	var v958 float64
	_ = v958
	var v961 float64
	_ = v961
	var v964 float64
	_ = v964
	var v966 int32
	_ = v966
	var v968 float64
	_ = v968
	var v972 float64
	_ = v972
	var v977 int32
	_ = v977
	var v978 float64
	_ = v978
	var v980 float64
	_ = v980
	var v986 float64
	_ = v986
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1018 int32
	_ = v1018
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1136 int32
	_ = v1136
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1276 int32
	_ = v1276
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1291 int32
	_ = v1291
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1368 int32
	_ = v1368
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1395 int32
	_ = v1395
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1530 int32
	_ = v1530
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 float64
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1567 float64
	_ = v1567
	var v1569 float64
	_ = v1569
	var v1571 float64
	_ = v1571
	var v1589 int32
	_ = v1589
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1605 int32
	_ = v1605
	var v1610 float32
	_ = v1610
	var v1613 float64
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
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
	if int32(0) < v250 {
		goto L56
	} else {
		goto L57
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
	v257 = v91 + int32(1)
	if v257 != l2 {
		v91 = v257
		v96 = v250
		v100 = v252
		v106 = v253
		v107 = v254
		v112 = v255
		goto L11
	} else {
		goto L55
	}
L16:
	;
	v250 = v96
	v252 = v100 + int32(1)
	v253 = v106
	v254 = v107
	v255 = v112
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
	v239 = v55 + v96<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+4)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v59+v96<<(uint(int32(2))%32)))) = v96
	v250 = v96 + int32(1)
	v252 = v100
	v253 = v132
	v254 = v107
	v255 = v236
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
		v233 = v124
		v236 = v112
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
	v250 = v96
	v252 = v100
	v253 = v132
	v254 = v107 + int32(1)
	v255 = v161
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
	v233 = v168
	v236 = v161
	goto L19
L37:
	;
	if v124&int32(3) == int32(0) {
		v195 = v124
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v233 = v124
	v236 = base.F64_add(v112, base.F64_convert_i32_u(v228+int32(1)))
	goto L19
L39:
	;
	v228 = v220 - v124
	goto L38
L40:
	;
	v199 = v195
	goto L49
L41:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v179 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v228 = int32(0)
	goto L38
L43:
	;
	goto L44
L44:
	;
	v184 = v124
	goto L45
L45:
	;
	v188 = v184 + int32(1)
	if v188&int32(3) == int32(0) {
		v195 = v188
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v220 = v188
	goto L39
L47:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v193 != 0 {
		v184 = v188
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v208 = int32(-2139062144)
	if (int32(16843008)-v205|v205)&v208 == v208 {
		v199 = v199 + int32(4)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v214 = v199
	goto L52
L51:
	;
	goto L50
L52:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	if v218 != 0 {
		v214 = v214 + int32(1)
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v220 = v214
	goto L39
L54:
	;
	goto L53
L55:
	;
	goto L12
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v35 + int32(12)
	F_qsort_interruptible(m, v55, v250, int32(8), int32(509), v35+int32(4))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if int32(0) < v253 {
		goto L247
	} else {
		goto L248
	}
L59:
	;
	v275 = int32(0)
	v288 = v5
	v289 = v5
	v290 = v5
	v295 = v5
	v300 = v26
	goto L60
L60:
	;
	v307 = v288 + int32(1)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v55+v275<<(uint(int32(3))%32))+4))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v59+v312<<(uint(int32(2))%32))))
	if v312 != v318 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v499 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v499)
	v502 = base.F64_convert_i32_s(l2)
	v504 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v252), v502))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v504
	if v50 != 0 {
		goto L81
	} else {
		goto L82
	}
L62:
	;
	v477 = v289
	v478 = v290
	v483 = v295
	v494 = v307
	goto L64
L63:
	;
	if v307 < int32(2) {
		v444 = v289
		v450 = v295
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v495 = base.F64_add(base.F64_mul(base.F64_convert_i32_u(v275), base.F64_convert_i32_s(v312)), v300)
	v497 = v275 + int32(1)
	if v497 != v250 {
		v275 = v497
		v288 = v494
		v289 = v477
		v290 = v478
		v295 = v483
		v300 = v495
		goto L60
	} else {
		goto L79
	}
L65:
	;
	v477 = v444
	v478 = v290 + int32(1)
	v483 = v450
	v494 = int32(0)
	goto L64
L66:
	;
	v325 = v295 + int32(1)
	v326 = base.B2i32(v289 < v51)
	if v326 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v63-int32(8)+v289<<(uint(int32(3))%32))))
	if v307 <= v332 {
		v444 = v289
		v450 = v325
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v334 = v326 + v289
	v336 = v334 - int32(1)
	if v336 <= int32(0) {
		v396 = v336
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v425 = v63 + v396<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v425)+4)) = v275 - v288
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = v307
	v444 = v334
	v450 = v325
	goto L65
L72:
	;
	__phi344 = v336
	__phi347 = v334
	v344 = __phi344
	v347 = __phi347
	goto L73
L73:
	;
	v373 = v63 + v347<<(uint(int32(3))%32)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v373-int32(16))))
	if v307 <= v376 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v396 = int32(0)
	goto L71
L75:
	;
	v396 = v344
	goto L71
L76:
	;
	goto L77
L77:
	;
	v380 = v63 + v344<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = v376
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v373-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+4)) = v384
	v386 = int32(1)
	if v386 < v344 {
		__phi344 = v344 - v386
		__phi347 = v344
		v344 = __phi344
		v347 = __phi347
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	goto L61
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v516
	if v483 == int32(0) {
		v564 = base.F32_neg(base.F32_sub(float32(1), v504))
		goto L87
	} else {
		goto L88
	}
L81:
	;
	v507 = base.F64_div(v255, base.F64_convert_i32_s(v253))
	if base.F64_lt(base.F64_abs(v507), float64(2.147483648e+09)) != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v514 = int32(*(*int16)(unsafe.Add(mBase, uint32(v513)+76)))
	v516 = v514
	goto L80
L84:
	;
	v511 = base.I32_trunc_f64_s(v507)
	v516 = v511
	goto L80
L85:
	;
	goto L86
L86:
	;
	v516 = int32(-2147483648)
	goto L80
L87:
	;
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v564
	v566 = base.F64_promote_f32(v564)
	if base.F64_gt(v566, base.F64_mul(l3, float64(0.1))) != 0 {
		goto L102
	} else {
		goto L103
	}
L88:
	;
	if v254 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v525 = v478 + v254
	v526 = float64(0)
	v530 = base.F64_mul(l3, base.F64_sub(float64(1), base.F64_promote_f32(v504)))
	if base.F64_gt(v530, v526) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	if v478 != v483 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v564 = base.F32_convert_i32_s(v483)
	goto L87
L92:
	;
	if base.F64_lt(v548, v549) != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v548 = v526
	v549 = base.F64_convert_i32_s(v525)
	goto L92
L94:
	;
	goto L95
L95:
	;
	v537 = base.F64_convert_i32_s(l2 - v252)
	v538 = base.F64_convert_i32_s(v525)
	v542 = base.F64_convert_i32_s(v254 - v483 + v478)
	v548 = base.F64_div(base.F64_mul(v537, v538), base.F64_add(base.F64_sub(v537, v542), base.F64_div(base.F64_mul(v537, v542), v530)))
	v549 = v538
	goto L92
L96:
	;
	v552 = v549
	goto L98
L97:
	;
	v552 = v548
	goto L98
L98:
	;
	if base.F64_gt(v552, v530) != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v554 = v530
	goto L101
L100:
	;
	v554 = v552
	goto L101
L101:
	;
	v564 = base.F32_demote_f64(base.F64_floor(base.F64_add(v554, float64(0.5))))
	goto L87
L102:
	;
	v572 = base.F32_demote_f64(base.F64_div(base.F64_neg(v566), l3))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v572
	v574 = v572
	goto L104
L103:
	;
	v574 = v564
	goto L104
L104:
	;
	if v477 != v478 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v1170 = v478 - v1136
	if v51 < v1170 {
		goto L171
	} else {
		goto L172
	}
L106:
	;
	v1049 = int32(0)
	if v1018 <= v1049 {
		v1136 = v1018
		v1167 = v1049
		goto L105
	} else {
		goto L164
	}
L107:
	;
	if v51 < v477 {
		goto L112
	} else {
		goto L113
	}
L108:
	;
	if v254 != 0 {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	if base.F32_gt(v574, float32(0)) == int32(0) {
		goto L107
	} else {
		goto L110
	}
L110:
	;
	if v478 <= v51 {
		v1018 = v478
		goto L106
	} else {
		goto L111
	}
L111:
	;
	goto L107
L112:
	;
	v584 = v51
	goto L114
L113:
	;
	v584 = v477
	goto L114
L114:
	;
	if v584 <= int32(0) {
		v1136 = v584
		v1167 = int32(0)
		goto L105
	} else {
		goto L115
	}
L115:
	;
	v588 = v584 & int32(3)
	v592 = F_palloc(m, v584<<(uint(int32(2))%32))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	v594 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v584) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v604 = v594
	v608 = int32(0)
	goto L120
L118:
	;
	v679 = v594
	goto L119
L119:
	;
	if v588 != 0 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v632 = int32(2)
	v635 = int32(3)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v63+v604<<(uint(v635)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v592+v604<<(uint(v632)%32)))) = v638
	v641 = v604 | int32(1)
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v63+v641<<(uint(v635)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v592+v641<<(uint(v632)%32)))) = v648
	v651 = v604 | v632
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v63+v651<<(uint(v635)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v592+v651<<(uint(v632)%32)))) = v658
	v661 = v604 | v635
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v63+v661<<(uint(v635)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v592+v661<<(uint(v632)%32)))) = v668
	v670 = int32(4)
	v671 = v604 + v670
	v673 = v608 + v670
	if v673 != v584&int32(2147483644) {
		v604 = v671
		v608 = v673
		goto L120
	} else {
		goto L122
	}
L121:
	;
	v679 = v671
	goto L119
L122:
	;
	goto L121
L123:
	;
	v711 = v679
	v714 = int32(0)
	goto L126
L124:
	;
	goto L125
L125:
	;
	v784 = *(*float32)(unsafe.Add(mBase, uint32(l0)+48))
	v785 = base.F64_promote_f32(v784)
	v786 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v788 = float64(0)
	v794 = int32(0)
	v800 = base.F64_convert_i32_s(l2)
	if base.F64_eq(l3, v800) != 0 {
		v1000 = v584
		goto L130
	} else {
		goto L131
	}
L126:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v63+v711<<(uint(int32(3))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v592+v711<<(uint(int32(2))%32)))) = v745
	v747 = int32(1)
	v750 = v714 + v747
	if v750 != v588 {
		v711 = v711 + v747
		v714 = v750
		goto L126
	} else {
		goto L128
	}
L127:
	;
	goto L125
L128:
	;
	goto L127
L129:
	;
	v1018 = v1000
	goto L106
L130:
	;
	goto L129
L131:
	;
	if base.F64_le(l3, float64(1)) != 0 {
		v1000 = v584
		goto L130
	} else {
		goto L132
	}
L132:
	;
	if base.Ui32(v584) < base.Ui32(int32(2)) {
		v915 = v788
		goto L133
	} else {
		goto L134
	}
L133:
	;
	if base.F64_lt(v785, float64(0)) != 0 {
		goto L146
	} else {
		goto L147
	}
L134:
	;
	v811 = v584 - int32(1)
	v812 = int32(3)
	v813 = v811 & v812
	if base.Ui32(v584-int32(2)) < base.Ui32(v812) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	if v813 == int32(0) {
		v915 = v866
		goto L133
	} else {
		goto L142
	}
L136:
	;
	v864 = int32(0)
	v866 = v788
	goto L135
L137:
	;
	goto L138
L138:
	;
	v826 = int32(0)
	v828 = v788
	v837 = v794
	goto L139
L139:
	;
	v842 = v592 + v826<<(uint(int32(2))%32)
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v842)))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v842)+4))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v842)+8))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v842)+12))
	v854 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v828, base.F64_convert_i32_s(v843)), base.F64_convert_i32_s(v846)), base.F64_convert_i32_s(v849)), base.F64_convert_i32_s(v852))
	v855 = int32(4)
	v856 = v826 + v855
	v858 = v837 + v855
	if v858 != v811&int32(-4) {
		v826 = v856
		v828 = v854
		v837 = v858
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v864 = v856
	v866 = v854
	goto L135
L141:
	;
	goto L140
L142:
	;
	v884 = v864
	v886 = v866
	v894 = v794
	goto L143
L143:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v592+v884<<(uint(int32(2))%32))))
	v903 = base.F64_add(v886, base.F64_convert_i32_s(v901))
	v904 = int32(1)
	v907 = v894 + v904
	if v907 != v813 {
		v884 = v884 + v904
		v886 = v903
		v894 = v907
		goto L143
	} else {
		goto L145
	}
L144:
	;
	v915 = v903
	goto L133
L145:
	;
	goto L144
L146:
	;
	v927 = base.F64_mul(l3, base.F64_neg(v785))
	goto L148
L147:
	;
	v927 = v785
	goto L148
L148:
	;
	v936 = v584
	v941 = v915
	goto L149
L149:
	;
	v953 = float64(1)
	v954 = float64(0)
	v958 = base.F64_sub(base.F64_sub(v953, base.F64_div(v941, v800)), base.F64_promote_f32(v786))
	if base.F64_lt(v958, v954) != 0 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v1000 = int32(0)
	goto L130
L151:
	;
	v961 = v954
	goto L153
L152:
	;
	v961 = v958
	goto L153
L153:
	;
	if base.F64_gt(v961, float64(1)) != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v964 = v953
	goto L156
L155:
	;
	v964 = v961
	goto L156
L156:
	;
	v966 = v936 - int32(1)
	v968 = base.F64_sub(v927, base.F64_convert_i32_u(v966))
	if base.F64_gt(v968, float64(1)) != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v972 = base.F64_div(v964, v968)
	goto L159
L158:
	;
	v972 = v964
	goto L159
L159:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v592+v966<<(uint(int32(2))%32))))
	v978 = base.F64_convert_i32_s(v977)
	v980 = base.F64_div(base.F64_mul(l3, v978), v800)
	v986 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l3, v800), base.F64_mul(base.F64_mul(v980, v800), base.F64_sub(l3, v980))), base.F64_mul(base.F64_mul(l3, l3), base.F64_add(l3, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v972, v800), base.F64_add(v986, v986)), float64(0.5)), v978) != 0 {
		v1000 = v936
		goto L130
	} else {
		goto L160
	}
L160:
	;
	if v966 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v592-int32(8)+v936<<(uint(int32(2))%32))))
	v936 = v966
	v941 = base.F64_sub(v941, base.F64_convert_i32_s(v995))
	goto L149
L162:
	;
	goto L163
L163:
	;
	goto L150
L164:
	;
	v1053 = int32(_a_F_compute_scalar_stats_1)
	v1054 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1056
	v1059 = v1018 << (uint(int32(2)) % 32)
	v1060 = F_palloc(m, v1059)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	v1062 = F_palloc(m, v1059)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v1068 = int32(0)
	goto L167
L167:
	;
	v1097 = v1068 << (uint(int32(2)) % 32)
	v1099 = int32(3)
	v1101 = v63 + v1068<<(uint(v1099)%32)
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+4))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v55+v1102<<(uint(v1099)%32))))
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107)+78)))
	v1109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1107)+76)))
	v1110 = F_datumCopy(m, v1106, v1108, v1109)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L4
	} else {
		goto L169
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1054
	v1124 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v1124)
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v1062
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1129
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v1060
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v1018
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v1018
	v1136 = v1018
	v1167 = v1124
	goto L105
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1060+v1097))) = v1110
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1101)))
	*(*float32)(unsafe.Add(mBase, uint32(v1097+v1062))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v1114), v502))
	v1120 = v1068 + int32(1)
	if v1120 != v1018 {
		v1068 = v1120
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v1172 = v51 + int32(1)
	goto L173
L172:
	;
	v1172 = v1170
	goto L173
L173:
	;
	if int32(2) <= v1172 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v1175 = int32(0)
	F_qsort_interruptible(m, v63, v1136, int32(8), int32(510), v1175)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L4
	} else {
		goto L177
	}
L175:
	;
	v1530 = v1167
	goto L176
L176:
	;
	if v250 == int32(1) {
		goto L9
	} else {
		goto L245
	}
L177:
	;
	if v1167 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v1181 = int32(0)
	v1187 = v1181
	v1190 = v1181
	v1191 = v1175
	goto L181
L179:
	;
	v1395 = v250
	goto L180
L180:
	;
	v1420 = int32(_a_F_compute_scalar_stats_1)
	v1421 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1423
	v1425 = int32(1)
	v1426 = v1395 - v1425
	v1428 = v1172 - v1425
	v1429 = base.I32_div_s(v1426, v1428)
	if v1172 <= v1425 {
		goto L234
	} else {
		goto L235
	}
L181:
	;
	if v1136 <= v1191 {
		v1225 = v250
		goto L184
	} else {
		goto L185
	}
L182:
	;
	v1395 = v1383
	goto L180
L183:
	;
	if v1386 < v250 {
		v1187 = v1386
		v1190 = v1383
		v1191 = v1384
		goto L181
	} else {
		goto L233
	}
L184:
	;
	v1227 = int32(3)
	v1229 = v55 + v1190<<(uint(v1227)%32)
	v1232 = v55 + v1187<<(uint(v1227)%32)
	v1233 = v1225 - v1187
	v1235 = v1233 << (uint(v1227) % 32)
	if v1229 == v1232 {
		goto L188
	} else {
		goto L189
	}
L185:
	;
	v1218 = v63 + v1191<<(uint(int32(3))%32)
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+4))
	if v1187 < v1219 {
		v1225 = v1219
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1218)))
	v1383 = v1190
	v1384 = v1191 + int32(1)
	v1386 = v1223 + v1219
	goto L183
L187:
	;
	v1383 = v1233 + v1190
	v1384 = v1191
	v1386 = v1225
	goto L183
L188:
	;
	goto L187
L189:
	;
	v1239 = v1229 + v1235
	if base.Ui32(v1232-v1239) <= base.Ui32(int32(0)-v1235<<(uint(int32(1))%32)) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v1246 = F___memcpy(m, v1229, v1232, v1235)
	mBase = m.M
	goto L187
L191:
	;
	goto L192
L192:
	;
	v1249 = (v1229 ^ v1232) & int32(3)
	if base.Ui32(v1229) < base.Ui32(v1232) {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	if v1351 == int32(0) {
		goto L188
	} else {
		goto L229
	}
L194:
	;
	if base.Ui32(v1329) <= base.Ui32(int32(3)) {
		v1350 = v1328
		v1351 = v1329
		v1352 = v1330
		goto L193
	} else {
		goto L225
	}
L195:
	;
	if v1249 != 0 {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L197
L197:
	;
	if v1249 != 0 {
		v1311 = v1235
		goto L208
	} else {
		goto L209
	}
L198:
	;
	v1350 = v1232
	v1351 = v1235
	v1352 = v1229
	goto L193
L199:
	;
	goto L200
L200:
	;
	if v1229&int32(3) == int32(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1328 = v1232
	v1329 = v1235
	v1330 = v1229
	goto L194
L202:
	;
	goto L203
L203:
	;
	v1256 = v1232
	v1257 = v1235
	v1258 = v1229
	goto L204
L204:
	;
	if v1257 == int32(0) {
		goto L188
	} else {
		goto L206
	}
L205:
	;
	v1328 = v1265
	v1329 = v1267
	v1330 = v1269
	goto L194
L206:
	;
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1256))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1258))) = uint8(v1262)
	v1264 = int32(1)
	v1265 = v1256 + v1264
	v1267 = v1257 - v1264
	v1269 = v1258 + v1264
	if v1269&int32(3) != 0 {
		v1256 = v1265
		v1257 = v1267
		v1258 = v1269
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	if v1311 == int32(0) {
		goto L188
	} else {
		goto L221
	}
L209:
	;
	if v1239&int32(3) != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1276 = v1235
	goto L213
L211:
	;
	v1291 = v1235
	goto L212
L212:
	;
	if base.Ui32(v1291) <= base.Ui32(int32(3)) {
		v1311 = v1291
		goto L208
	} else {
		goto L217
	}
L213:
	;
	if v1276 == int32(0) {
		goto L188
	} else {
		goto L215
	}
L214:
	;
	v1291 = v1282
	goto L212
L215:
	;
	v1282 = v1276 - int32(1)
	v1283 = v1229 + v1282
	v1285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232+v1282))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1283))) = uint8(v1285)
	if v1283&int32(3) != 0 {
		v1276 = v1282
		goto L213
	} else {
		goto L216
	}
L216:
	;
	goto L214
L217:
	;
	v1298 = v1291
	goto L218
L218:
	;
	v1302 = v1298 - int32(4)
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1232+v1302)))
	*(*int32)(unsafe.Add(mBase, uint32(v1229+v1302))) = v1305
	if base.Ui32(int32(3)) < base.Ui32(v1302) {
		v1298 = v1302
		goto L218
	} else {
		goto L220
	}
L219:
	;
	v1311 = v1302
	goto L208
L220:
	;
	goto L219
L221:
	;
	v1318 = v1311
	goto L222
L222:
	;
	v1322 = v1318 - int32(1)
	v1325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232+v1322))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1229+v1322))) = uint8(v1325)
	if v1322 != 0 {
		v1318 = v1322
		goto L222
	} else {
		goto L224
	}
L223:
	;
	goto L188
L224:
	;
	goto L223
L225:
	;
	v1335 = v1328
	v1336 = v1329
	v1337 = v1330
	goto L226
L226:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1335)))
	*(*int32)(unsafe.Add(mBase, uint32(v1337))) = v1339
	v1341 = int32(4)
	v1342 = v1335 + v1341
	v1344 = v1337 + v1341
	v1346 = v1336 - v1341
	if base.Ui32(int32(3)) < base.Ui32(v1346) {
		v1335 = v1342
		v1336 = v1346
		v1337 = v1344
		goto L226
	} else {
		goto L228
	}
L227:
	;
	v1350 = v1342
	v1351 = v1346
	v1352 = v1344
	goto L193
L228:
	;
	goto L227
L229:
	;
	v1357 = v1350
	v1358 = v1351
	v1359 = v1352
	goto L230
L230:
	;
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1357))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1359))) = uint8(v1361)
	v1363 = int32(1)
	v1368 = v1358 - v1363
	if v1368 != 0 {
		v1357 = v1357 + v1363
		v1358 = v1368
		v1359 = v1359 + v1363
		goto L230
	} else {
		goto L232
	}
L231:
	;
	goto L188
L232:
	;
	goto L231
L233:
	;
	goto L182
L234:
	;
	v1435 = v1425
	goto L236
L235:
	;
	v1435 = v1172
	goto L236
L236:
	;
	v1438 = F_palloc(m, v1172<<(uint(int32(2))%32))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L4
	} else {
		goto L237
	}
L237:
	;
	v1440 = int32(0)
	v1444 = v1440
	v1447 = v1440
	v1449 = v1440
	goto L238
L238:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v55+v1449<<(uint(int32(3))%32))))
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1482)+78)))
	v1484 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1482)+76)))
	v1485 = F_datumCopy(m, v1481, v1483, v1484)
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L4
	} else {
		goto L240
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1421
	v1500 = int32(1)
	v1503 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v1167<<(uint(v1500)%32))+52)) = uint16(v1503)
	v1507 = l0 + v1167<<(uint(v1503)%32)
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1507-int32(-64)))) = v1510
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+164)) = v1438
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+84)) = v1512
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+144)) = v1172
	v1530 = v1167 + v1500
	goto L176
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1438+v1447<<(uint(int32(2))%32)))) = v1485
	v1488 = v1444 + (v1426 - v1429*v1428)
	v1489 = base.B2i32(v1428 <= v1488)
	if v1428 <= v1488 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1493 = v1428
	goto L243
L242:
	;
	v1493 = int32(0)
	goto L243
L243:
	;
	v1496 = v1447 + int32(1)
	if v1496 != v1435 {
		v1444 = v1488 - v1493
		v1447 = v1496
		v1449 = v1489 + (v1449 + v1429)
		goto L238
	} else {
		goto L244
	}
L244:
	;
	goto L239
L245:
	;
	v1552 = int32(_a_F_compute_scalar_stats_1)
	v1553 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1555
	v1558 = F_palloc(m, int32(4))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L4
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1553
	v1562 = base.F64_convert_i32_u(v250)
	v1564 = int32(1)
	v1567 = base.F64_mul(v1562, base.F64_convert_i32_u(v250-v1564))
	v1569 = base.F64_mul(v1567, float64(0.5))
	v1571 = base.F64_mul(v1569, base.F64_neg(v1569))
	*(*float32)(unsafe.Add(mBase, uint32(v1558))) = base.F32_demote_f64(base.F64_div(base.F64_add(base.F64_mul(v1562, v495), v1571), base.F64_add(base.F64_mul(v1562, base.F64_div(base.F64_mul(v1567, base.F64_convert_i32_s(v250<<(uint(v1564)%32)-v1564)), float64(6))), v1571)))
	v1589 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v1530<<(uint(v1564)%32))+52)) = uint16(v1589)
	v1593 = l0 + v1530<<(uint(int32(2))%32)
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1593-int32(-64)))) = v1596
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1593)+124)) = v1558
	*(*int32)(unsafe.Add(mBase, uint32(v1593)+84)) = v1598
	*(*int32)(unsafe.Add(mBase, uint32(v1593)+104)) = v1564
	goto L9
L247:
	;
	v1605 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1605)
	v1610 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v252), base.F64_convert_i32_s(l2)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v1610
	if v50 != 0 {
		goto L251
	} else {
		goto L252
	}
L248:
	;
	goto L249
L249:
	;
	if v252 <= int32(0) {
		goto L9
	} else {
		goto L257
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1622
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = base.F32_neg(base.F32_sub(float32(1), v1610))
	goto L9
L251:
	;
	v1613 = base.F64_div(v255, base.F64_convert_i32_u(v253))
	if base.F64_lt(base.F64_abs(v1613), float64(2.147483648e+09)) != 0 {
		goto L254
	} else {
		goto L255
	}
L252:
	;
	goto L253
L253:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1620 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1619)+76)))
	v1622 = v1620
	goto L250
L254:
	;
	v1617 = base.I32_trunc_f64_s(v1613)
	v1622 = v1617
	goto L250
L255:
	;
	goto L256
L256:
	;
	v1622 = int32(-2147483648)
	goto L250
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v1632 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1632)
	v1634 = int32(0)
	if v50 == v1634 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1638 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1637)+76)))
	v1639 = v1638
	goto L260
L259:
	;
	v1639 = v1634
	goto L260
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1639
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
