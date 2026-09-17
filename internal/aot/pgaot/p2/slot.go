package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecBuildSlotValueDescription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v210 int32
	_ = v210
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v20 = F_check_enable_rls(m, l0, v5, int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(48)
	return v210
L2:
	;
	return int32(0)
L3:
	;
	if v20 == int32(2) {
		v210 = v5
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = v16 + int32(32)
	F_initStringInfo(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	F_appendStringInfoChar(m, v27, int32(40))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_ExecBuildSlotValueDescription[0]))
	v36 = F_pg_class_aclcheck(m, l0, v34, int64(2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v39 = v16 + int32(16)
	F_initStringInfo(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v48 < v47 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	F_appendStringInfoChar(m, v39, int32(40))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	F_slot_getsomeattrs_int(m, l1, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v52 = int32(0)
	v53 = base.B2i32(v36 == v52)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v52 < v54 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v63 = v54
	v64 = v53
	v65 = int32(0)
	v68 = v5
	v70 = v5
	goto L20
L18:
	;
	v174 = v53
	goto L19
L19:
	;
	if v174 == int32(0) {
		v210 = v5
		goto L1
	} else {
		goto L56
	}
L20:
	;
	v76 = l2 + v63<<(uint(int32(4))%32) + v65*int32(100)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+111)))
	if v77 != 0 {
		v160 = v64
		v162 = v68
		v163 = v70
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v174 = v160
	goto L19
L22:
	;
	v165 = v65 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v165 < v166 {
		v63 = v166
		v64 = v160
		v65 = v165
		v68 = v162
		v70 = v163
		goto L20
	} else {
		goto L55
	}
L23:
	;
	v79 = v76 + int32(20)
	if v36 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79)+74)))
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_ExecBuildSlotValueDescription[0]))
	v84 = F_pg_attribute_aclcheck(m, l0, v80, v82, int64(2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L27
	}
L25:
	;
	v108 = v64
	v110 = v70
	goto L26
L26:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+90)))
	if v112 == int32(118) {
		v134 = int32(_a_F_ExecBuildSlotValueDescription_0)
		goto L38
	} else {
		goto L39
	}
L27:
	;
	v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79)+74)))
	v89 = F_bms_is_member(m, v86+int32(7), l3)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if v84 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v92 = v89
	goto L31
L30:
	;
	v92 = int32(1)
	goto L31
L31:
	;
	if v92 == int32(0) {
		v160 = v64
		v162 = v68
		v163 = v70
		goto L22
	} else {
		goto L32
	}
L32:
	;
	if v70 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_appendStringInfoString(m, v16+int32(16), int32(_a_F_ExecBuildSlotValueDescription_1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_appendStringInfoString(m, v16+int32(16), v76+int32(24))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	v106 = int32(1)
	v108 = v106
	v110 = v106
	goto L26
L38:
	;
	if v68 != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v65))))
	if v118 != 0 {
		v134 = int32(_a_F_ExecBuildSlotValueDescription_2)
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v79)+68))
	F_getTypeOutputInfo(m, v119, v16+int32(12), v16+int32(11))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v65<<(uint(int32(2))%32))))
	v132 = F_OidOutputFunctionCall(m, v126, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v134 = v132
	goto L38
L43:
	;
	F_appendStringInfoString(m, v16+int32(32), int32(_a_F_ExecBuildSlotValueDescription_1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v140 = F_strlen(m, v134)
	mBase = m.M
	if v140 <= int32(64) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L45
L47:
	;
	v160 = v108
	v162 = int32(1)
	v163 = v110
	goto L22
L48:
	;
	F_appendBinaryStringInfo(m, v16+int32(32), v134, v140)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v148 = v16 + int32(32)
	v150 = F_pg_mbcliplen(m, v134, v140, int32(64))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L52
	}
L51:
	;
	goto L47
L52:
	;
	F_appendBinaryStringInfo(m, v148, v134, v150)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	F_appendStringInfoString(m, v148, int32(_a_F_ExecBuildSlotValueDescription_3))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	goto L47
L55:
	;
	goto L21
L56:
	;
	F_appendStringInfoChar(m, v16+int32(32), int32(41))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	if v36 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v189 = v16 + int32(16)
	F_appendStringInfoString(m, v189, int32(_a_F_ExecBuildSlotValueDescription_4))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v210 = v198
	goto L1
L61:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	F_appendBinaryStringInfo(m, v189, v193, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v210 = v197
	goto L1
}
func F_ExecComputeSlotInfo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v3)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v13 != 0 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(base.B2i32(v14 != int32(0)))
		v156 = v13
		v157 = v14
		v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if base.B2i32(v156 == int32(0))|base.B2i32(v160 != int32(1)) != 0 {
			v166 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
			*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
			v182 = int32(1)
		} else {
			if v157 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v157
				v171 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v171)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
				if v157 != int32(_a_F_ExecComputeSlotInfo_0) {
					v182 = int32(1)
				} else {
					v182 = int32(0)
				}
			} else {
				v166 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
				v182 = int32(1)
			}
		}
	} else {
		if v10 == int32(0) {
			v166 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
			*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
			v182 = int32(1)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			switch v20 - int32(2) {
			case 0:
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+102)))
				if v24 != int32(1) {
					if v23 == int32(0) {
						v166 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
						v182 = int32(1)
					} else {
						v40 = v8 + int32(15)
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+103)))
						if v42 == int32(1) {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
							if v45 != 0 {
								if v40 == int32(0) {
									v78 = v45
									v82 = v78
								} else {
									v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+99)))
									*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v48)
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
									v82 = v50
								}
							} else {
								if v40 == int32(0) {
								} else {
									v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+99)))
									v67 = v53
									*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v67)
								}
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
								if v71 == int32(0) {
									v82 = int32(_a_F_ExecComputeSlotInfo_0)
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
									v78 = v75
									v82 = v78
								}
							}
						} else {
							if v40 == int32(0) {
							} else {
								v56 = int32(0)
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
								if v57 == v56 {
									v67 = v56
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+4)))
									v67 = int32(base.Ui32(v60)>>(uint(int32(4))%32)) & int32(1)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v67)
							}
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
							if v71 == int32(0) {
								v82 = int32(_a_F_ExecComputeSlotInfo_0)
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
								v78 = v75
								v82 = v78
							}
						}
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
						v156 = v83
						v157 = v82
						v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						if base.B2i32(v156 == int32(0))|base.B2i32(v160 != int32(1)) != 0 {
							v166 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
							v182 = int32(1)
						} else {
							if v157 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v157
								v171 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v171)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
								if v157 != int32(_a_F_ExecComputeSlotInfo_0) {
									v182 = int32(1)
								} else {
									v182 = int32(0)
								}
							} else {
								v166 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v182 = int32(1)
							}
						}
					}
				} else {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+98)))
					if v27 != int32(1) {
						v166 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
						v182 = int32(1)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
						if v30 == int32(0) {
							if v23 == int32(0) {
								v166 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v182 = int32(1)
							} else {
								v40 = v8 + int32(15)
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+103)))
								if v42 == int32(1) {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
									if v45 != 0 {
										if v40 == int32(0) {
											v78 = v45
											v82 = v78
										} else {
											v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+99)))
											*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v48)
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
											v82 = v50
										}
									} else {
										if v40 == int32(0) {
										} else {
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+99)))
											v67 = v53
											*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v67)
										}
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
										if v71 == int32(0) {
											v82 = int32(_a_F_ExecComputeSlotInfo_0)
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
											v78 = v75
											v82 = v78
										}
									}
								} else {
									if v40 == int32(0) {
									} else {
										v56 = int32(0)
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
										if v57 == v56 {
											v67 = v56
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+4)))
											v67 = int32(base.Ui32(v60)>>(uint(int32(4))%32)) & int32(1)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v67)
									}
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
									if v71 == int32(0) {
										v82 = int32(_a_F_ExecComputeSlotInfo_0)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
										v78 = v75
										v82 = v78
									}
								}
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
								v156 = v83
								v157 = v82
								v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
								if base.B2i32(v156 == int32(0))|base.B2i32(v160 != int32(1)) != 0 {
									v166 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
									*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
									v182 = int32(1)
								} else {
									if v157 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v157
										v171 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v171)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
										if v157 != int32(_a_F_ExecComputeSlotInfo_0) {
											v182 = int32(1)
										} else {
											v182 = int32(0)
										}
									} else {
										v166 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
										*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
										v182 = int32(1)
									}
								}
							}
						} else {
							v33 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v33)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
							v156 = v35
							v157 = v30
							v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							if base.B2i32(v156 == int32(0))|base.B2i32(v160 != int32(1)) != 0 {
								v166 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v182 = int32(1)
							} else {
								if v157 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v157
									v171 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v171)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
									if v157 != int32(_a_F_ExecComputeSlotInfo_0) {
										v182 = int32(1)
									} else {
										v182 = int32(0)
									}
								} else {
									v166 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
									*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
									v182 = int32(1)
								}
							}
						}
					}
				}
			case 1:
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
				v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+101)))
				if v85 != int32(1) {
					if v84 == int32(0) {
						v166 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
						v182 = int32(1)
					} else {
						v101 = v8 + int32(15)
						v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+103)))
						if v103 == int32(1) {
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v84)+92))
							if v106 != 0 {
								if v101 == int32(0) {
									v139 = v106
									v143 = v139
								} else {
									v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+99)))
									*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v109)
									v111 = *(*int32)(unsafe.Add(mBase, uint32(v84)+92))
									v143 = v111
								}
							} else {
								if v101 == int32(0) {
								} else {
									v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+99)))
									v128 = v114
									*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v128)
								}
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v84)+60))
								if v132 == int32(0) {
									v143 = int32(_a_F_ExecComputeSlotInfo_0)
								} else {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
									v139 = v136
									v143 = v139
								}
							}
						} else {
							if v101 == int32(0) {
							} else {
								v117 = int32(0)
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v84)+60))
								if v118 == v117 {
									v128 = v117
								} else {
									v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
									v128 = int32(base.Ui32(v121)>>(uint(int32(4))%32)) & int32(1)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v128)
							}
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v84)+60))
							if v132 == int32(0) {
								v143 = int32(_a_F_ExecComputeSlotInfo_0)
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
								v139 = v136
								v143 = v139
							}
						}
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v84)+56))
						v156 = v144
						v157 = v143
						v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						if base.B2i32(v156 == int32(0))|base.B2i32(v160 != int32(1)) != 0 {
							v166 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
							v182 = int32(1)
						} else {
							if v157 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v157
								v171 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v171)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
								if v157 != int32(_a_F_ExecComputeSlotInfo_0) {
									v182 = int32(1)
								} else {
									v182 = int32(0)
								}
							} else {
								v166 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v182 = int32(1)
							}
						}
					}
				} else {
					v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+97)))
					if v88 != int32(1) {
						v166 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
						v182 = int32(1)
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
						if v91 == int32(0) {
							if v84 == int32(0) {
								v166 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v182 = int32(1)
							} else {
								v101 = v8 + int32(15)
								v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+103)))
								if v103 == int32(1) {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v84)+92))
									if v106 != 0 {
										if v101 == int32(0) {
											v139 = v106
											v143 = v139
										} else {
											v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+99)))
											*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v109)
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v84)+92))
											v143 = v111
										}
									} else {
										if v101 == int32(0) {
										} else {
											v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+99)))
											v128 = v114
											*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v128)
										}
										v132 = *(*int32)(unsafe.Add(mBase, uint32(v84)+60))
										if v132 == int32(0) {
											v143 = int32(_a_F_ExecComputeSlotInfo_0)
										} else {
											v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
											v139 = v136
											v143 = v139
										}
									}
								} else {
									if v101 == int32(0) {
									} else {
										v117 = int32(0)
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v84)+60))
										if v118 == v117 {
											v128 = v117
										} else {
											v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
											v128 = int32(base.Ui32(v121)>>(uint(int32(4))%32)) & int32(1)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v128)
									}
									v132 = *(*int32)(unsafe.Add(mBase, uint32(v84)+60))
									if v132 == int32(0) {
										v143 = int32(_a_F_ExecComputeSlotInfo_0)
									} else {
										v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
										v139 = v136
										v143 = v139
									}
								}
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v84)+56))
								v156 = v144
								v157 = v143
								v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
								if base.B2i32(v156 == int32(0))|base.B2i32(v160 != int32(1)) != 0 {
									v166 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
									*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
									v182 = int32(1)
								} else {
									if v157 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v157
										v171 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v171)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
										if v157 != int32(_a_F_ExecComputeSlotInfo_0) {
											v182 = int32(1)
										} else {
											v182 = int32(0)
										}
									} else {
										v166 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
										*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
										v182 = int32(1)
									}
								}
							}
						} else {
							v94 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v94)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v84)+56))
							v156 = v96
							v157 = v91
							v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							if base.B2i32(v156 == int32(0))|base.B2i32(v160 != int32(1)) != 0 {
								v166 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v182 = int32(1)
							} else {
								if v157 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v157
									v171 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v171)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
									if v157 != int32(_a_F_ExecComputeSlotInfo_0) {
										v182 = int32(1)
									} else {
										v182 = int32(0)
									}
								} else {
									v166 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
									*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
									v182 = int32(1)
								}
							}
						}
					}
				}
			default:
				if base.Ui32(int32(2)) < base.Ui32(v20-int32(4)) {
					v166 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
					*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
					v182 = int32(1)
				} else {
					v149 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
					v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+100)))
					if v151 != int32(1) {
						v156 = v150
						v157 = v149
					} else {
						v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+96)))
						*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v154)
						v156 = v150
						v157 = v149
					}
					v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					if base.B2i32(v156 == int32(0))|base.B2i32(v160 != int32(1)) != 0 {
						v166 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
						v182 = int32(1)
					} else {
						if v157 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v157
							v171 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v171)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v156
							if v157 != int32(_a_F_ExecComputeSlotInfo_0) {
								v182 = int32(1)
							} else {
								v182 = int32(0)
							}
						} else {
							v166 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v166)
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
							v182 = int32(1)
						}
					}
				}
			}
		}
	}
	m.G0 = v8 + int32(16)
	return v182
}
func F_ExecCreateScanSlotFromOuterPlan(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+56))
	F_ExecInitScanTupleSlot(m, l0, l1, v5, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_ExecSetSlotDescriptor(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	m.T0[v5].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v16 != 0 {
				F_pfree(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v19 != 0 {
						F_pfree(m, v19)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							if int32(0) <= v23 {
								F_IncrTupleDescRefCount(m, l1)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						if int32(0) <= v23 {
							F_IncrTupleDescRefCount(m, l1)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v37 = F_MemoryContextAlloc(m, v35, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
									return
								}
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v19 != 0 {
					F_pfree(m, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						if int32(0) <= v23 {
							F_IncrTupleDescRefCount(m, l1)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v37 = F_MemoryContextAlloc(m, v35, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
									return
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if int32(0) <= v23 {
						F_IncrTupleDescRefCount(m, l1)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v37 = F_MemoryContextAlloc(m, v35, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
									return
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v37 = F_MemoryContextAlloc(m, v35, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
								return
							}
						}
					}
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v11 < int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v16 != 0 {
					F_pfree(m, v16)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v19 != 0 {
							F_pfree(m, v19)
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								if int32(0) <= v23 {
									F_IncrTupleDescRefCount(m, l1)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
											v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v37 = F_MemoryContextAlloc(m, v35, v36)
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
												return
											}
										}
									}
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							if int32(0) <= v23 {
								F_IncrTupleDescRefCount(m, l1)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v19 != 0 {
						F_pfree(m, v19)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							if int32(0) <= v23 {
								F_IncrTupleDescRefCount(m, l1)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						if int32(0) <= v23 {
							F_IncrTupleDescRefCount(m, l1)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v37 = F_MemoryContextAlloc(m, v35, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
									return
								}
							}
						}
					}
				}
			} else {
				F_DecrTupleDescRefCount(m, v8)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v16 != 0 {
						F_pfree(m, v16)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v19 != 0 {
								F_pfree(m, v19)
								mBase = m.M
								v21 = m.ExcPending
								if v21 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
									v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
									if int32(0) <= v23 {
										F_IncrTupleDescRefCount(m, l1)
										mBase = m.M
										v27 = m.ExcPending
										if v27 != 0 {
											return
										} else {
											v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
											mBase = m.M
											v33 = m.ExcPending
											if v33 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
												v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v37 = F_MemoryContextAlloc(m, v35, v36)
												mBase = m.M
												v38 = m.ExcPending
												if v38 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
													return
												}
											}
										}
									} else {
										v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
											v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v37 = F_MemoryContextAlloc(m, v35, v36)
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
												return
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								if int32(0) <= v23 {
									F_IncrTupleDescRefCount(m, l1)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
											v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v37 = F_MemoryContextAlloc(m, v35, v36)
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
												return
											}
										}
									}
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							}
						}
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v19 != 0 {
							F_pfree(m, v19)
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								if int32(0) <= v23 {
									F_IncrTupleDescRefCount(m, l1)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
											v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v37 = F_MemoryContextAlloc(m, v35, v36)
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
												return
											}
										}
									}
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							if int32(0) <= v23 {
								F_IncrTupleDescRefCount(m, l1)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
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
func F_GetSlotInvalidationCauseName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	if base.B2i32(int32(base.Ui32(int32(279))>>(uint(l0)%32))&int32(1) == int32(0))|base.B2i32(base.Ui32(int32(8)) < base.Ui32(l0)) != 0 {
		v18 = int32(_a_F_GetSlotInvalidationCauseName_0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_GetSlotInvalidationCauseName[0])))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		v18 = v17
	}
	return v18
}
func F_ReportSlotInvalidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	v5 = l4
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	F_initStringInfo(m, v12+int32(160))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		F_initStringInfo(m, v12+int32(144))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if l0&(l0-int32(1)) != 0 {
				v92 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					if v92 != 0 {
						if l1 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
							F_errmsg(m, int32(_a_F_ReportSlotInvalidation_0), v12+int32(32))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v107
								F_errdetail_internal(m, int32(_a_F_ReportSlotInvalidation_1), v12+int32(16))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return
								} else {
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
									if v114 != 0 {
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v115
										F_errhint(m, int32(_a_F_ReportSlotInvalidation_1), v12)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
												F_pfree(m, v125)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return
												} else {
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													F_pfree(m, v128)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return
													} else {
														m.G0 = v12 + int32(176)
														return
													}
												}
											}
										}
									} else {
										F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return
										} else {
											v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
											F_pfree(m, v125)
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return
											} else {
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												F_pfree(m, v128)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return
												} else {
													m.G0 = v12 + int32(176)
													return
												}
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
							F_errmsg(m, int32(_a_F_ReportSlotInvalidation_4), v12+int32(48))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v107
								F_errdetail_internal(m, int32(_a_F_ReportSlotInvalidation_1), v12+int32(16))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return
								} else {
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
									if v114 != 0 {
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v115
										F_errhint(m, int32(_a_F_ReportSlotInvalidation_1), v12)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
												F_pfree(m, v125)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return
												} else {
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													F_pfree(m, v128)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return
													} else {
														m.G0 = v12 + int32(176)
														return
													}
												}
											}
										}
									} else {
										F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return
										} else {
											v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
											F_pfree(m, v125)
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return
											} else {
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												F_pfree(m, v128)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return
												} else {
													m.G0 = v12 + int32(176)
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
						F_pfree(m, v125)
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return
						} else {
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
							F_pfree(m, v128)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return
							} else {
								m.G0 = v12 + int32(176)
								return
							}
						}
					}
				}
			} else {
				switch base.I32_ctz(l0) - int32(1) {
				case 0:
					*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = l6
					F_appendStringInfo(m, v12+int32(160), int32(_a_F_ReportSlotInvalidation_5), v12+int32(96))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v92 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							if v92 != 0 {
								if l1 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
									F_errmsg(m, int32(_a_F_ReportSlotInvalidation_0), v12+int32(32))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v107
										F_errdetail_internal(m, int32(_a_F_ReportSlotInvalidation_1), v12+int32(16))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
											if v114 != 0 {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v115
												F_errhint(m, int32(_a_F_ReportSlotInvalidation_1), v12)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v128)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											} else {
												F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v128)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
									F_errmsg(m, int32(_a_F_ReportSlotInvalidation_4), v12+int32(48))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return
									} else {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v107
										F_errdetail_internal(m, int32(_a_F_ReportSlotInvalidation_1), v12+int32(16))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
											if v114 != 0 {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v115
												F_errhint(m, int32(_a_F_ReportSlotInvalidation_1), v12)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v128)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											} else {
												F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v128)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
								F_pfree(m, v125)
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return
								} else {
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
									F_pfree(m, v128)
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return
									} else {
										m.G0 = v12 + int32(176)
										return
									}
								}
							}
						}
					}
				case 1:
					F_appendStringInfoString(m, v12+int32(160), int32(_a_F_ReportSlotInvalidation_6))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						v92 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							if v92 != 0 {
								if l1 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
									F_errmsg(m, int32(_a_F_ReportSlotInvalidation_0), v12+int32(32))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v107
										F_errdetail_internal(m, int32(_a_F_ReportSlotInvalidation_1), v12+int32(16))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
											if v114 != 0 {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v115
												F_errhint(m, int32(_a_F_ReportSlotInvalidation_1), v12)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v128)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											} else {
												F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v128)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
									F_errmsg(m, int32(_a_F_ReportSlotInvalidation_4), v12+int32(48))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return
									} else {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v107
										F_errdetail_internal(m, int32(_a_F_ReportSlotInvalidation_1), v12+int32(16))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
											if v114 != 0 {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v115
												F_errhint(m, int32(_a_F_ReportSlotInvalidation_1), v12)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v128)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											} else {
												F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v125)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v128)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
								F_pfree(m, v125)
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return
								} else {
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
									F_pfree(m, v128)
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return
									} else {
										m.G0 = v12 + int32(176)
										return
									}
								}
							}
						}
					}
				case 2:
					*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v12)+132)) = int32(_a_F_ReportSlotInvalidation_7)
					v71 = *(*int32)(unsafe.Add(mBase, _c_F_ReportSlotInvalidation[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+136)) = v71
					F_appendStringInfo(m, v12+int32(160), int32(_a_F_ReportSlotInvalidation_8), v12+int32(128))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = int32(_a_F_ReportSlotInvalidation_7)
						F_appendStringInfo(m, v12+int32(144), int32(_a_F_ReportSlotInvalidation_9), v12+int32(112))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							v92 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								if v92 != 0 {
									if l1 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
										F_errmsg(m, int32(_a_F_ReportSlotInvalidation_0), v12+int32(32))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v107
											F_errdetail_internal(m, int32(_a_F_ReportSlotInvalidation_1), v12+int32(16))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return
											} else {
												v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
												if v114 != 0 {
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													*(*int32)(unsafe.Add(mBase, uint32(v12))) = v115
													F_errhint(m, int32(_a_F_ReportSlotInvalidation_1), v12)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
															F_pfree(m, v125)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return
															} else {
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
																F_pfree(m, v128)
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
																	return
																} else {
																	m.G0 = v12 + int32(176)
																	return
																}
															}
														}
													}
												} else {
													F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v128)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
										F_errmsg(m, int32(_a_F_ReportSlotInvalidation_4), v12+int32(48))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return
										} else {
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v107
											F_errdetail_internal(m, int32(_a_F_ReportSlotInvalidation_1), v12+int32(16))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return
											} else {
												v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
												if v114 != 0 {
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													*(*int32)(unsafe.Add(mBase, uint32(v12))) = v115
													F_errhint(m, int32(_a_F_ReportSlotInvalidation_1), v12)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
															F_pfree(m, v125)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return
															} else {
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
																F_pfree(m, v128)
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
																	return
																} else {
																	m.G0 = v12 + int32(176)
																	return
																}
															}
														}
													}
												} else {
													F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v128)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											}
										}
									}
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
									F_pfree(m, v125)
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return
									} else {
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
										F_pfree(m, v128)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return
										} else {
											m.G0 = v12 + int32(176)
											return
										}
									}
								}
							}
						}
					}
				default:
					*(*uint32)(unsafe.Add(mBase, uint32(v12)+84)) = uint32(v5)
					v30 = int64(base.Ui64(v5) >> (uint(int64(32)) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v12)+80)) = uint32(v30)
					v32 = l5 - v5
					*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v32
					if v32 == int64(1) {
						v40 = int32(_a_F_ReportSlotInvalidation_10)
					} else {
						v40 = int32(_a_F_ReportSlotInvalidation_11)
					}
					F_appendStringInfo(m, v12+int32(160), v40, v12+int32(80))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = int32(_a_F_ReportSlotInvalidation_12)
						F_appendStringInfo(m, v12+int32(144), int32(_a_F_ReportSlotInvalidation_9), v12-int32(-64))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v92 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								if v92 != 0 {
									if l1 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
										F_errmsg(m, int32(_a_F_ReportSlotInvalidation_0), v12+int32(32))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v107
											F_errdetail_internal(m, int32(_a_F_ReportSlotInvalidation_1), v12+int32(16))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return
											} else {
												v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
												if v114 != 0 {
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													*(*int32)(unsafe.Add(mBase, uint32(v12))) = v115
													F_errhint(m, int32(_a_F_ReportSlotInvalidation_1), v12)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
															F_pfree(m, v125)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return
															} else {
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
																F_pfree(m, v128)
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
																	return
																} else {
																	m.G0 = v12 + int32(176)
																	return
																}
															}
														}
													}
												} else {
													F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v128)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
										F_errmsg(m, int32(_a_F_ReportSlotInvalidation_4), v12+int32(48))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return
										} else {
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v107
											F_errdetail_internal(m, int32(_a_F_ReportSlotInvalidation_1), v12+int32(16))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return
											} else {
												v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
												if v114 != 0 {
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													*(*int32)(unsafe.Add(mBase, uint32(v12))) = v115
													F_errhint(m, int32(_a_F_ReportSlotInvalidation_1), v12)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
															F_pfree(m, v125)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return
															} else {
																v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
																F_pfree(m, v128)
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
																	return
																} else {
																	m.G0 = v12 + int32(176)
																	return
																}
															}
														}
													}
												} else {
													F_errfinish(m, int32(_a_F_ReportSlotInvalidation_2), int32(1705), int32(_a_F_ReportSlotInvalidation_3))
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return
														} else {
															v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v128)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											}
										}
									}
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
									F_pfree(m, v125)
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return
									} else {
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
										F_pfree(m, v128)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return
										} else {
											m.G0 = v12 + int32(176)
											return
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
