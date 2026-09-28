package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NamedTuplestoreScanNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	F_tuplestore_select_read_pointer(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
		v13 = F_tuplestore_gettupleslot(m, v10, int32(1), int32(0), v3)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return v3
		}
	}
}
func F___netlink_enumerate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v182 int32
	_ = v182
	v3 = l2
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 + int32(-8192)
	m.G0 = v9
	goto L3
L1:
	;
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F___netlink_enumerate[0]))) = uint8(v119)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F___netlink_enumerate[1]))) = l1
	v122 = int32(769)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F___netlink_enumerate[2]))) = uint16(v122)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F___netlink_enumerate[3]))) = uint16(v3)
	v125 = int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F___netlink_enumerate[4]))) = v125
	v128 = F_sendto(m, l0, v9, v125)
	mBase = m.M
	if v128 < v119 {
		v182 = v128
		goto L13
	} else {
		goto L14
	}
L2:
	;
	goto L1
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F___netlink_enumerate[4]))) = uint8(v5)
	v19 = v7 + int32(-8172)
	*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(1)))) = uint8(v5)
	goto L4
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F___netlink_enumerate[5]))) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F___netlink_enumerate[6]))) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(3)))) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(2)))) = uint8(v5)
	goto L5
L5:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+uint32(_c_F___netlink_enumerate[7]))) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(4)))) = uint8(v5)
	goto L6
L6:
	;
	v41 = int32(0)
	v44 = (v41 - v9) & int32(3)
	v45 = v9 + v44
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v41
	v53 = (int32(20) - v44) & int32(-4)
	v54 = v45 + v53
	*(*int32)(unsafe.Add(mBase, uint32(v54-int32(4)))) = v41
	if base.Ui32(v53) < base.Ui32(int32(9)) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v54-int32(8)))) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v54-int32(12)))) = v41
	if base.Ui32(v53) < base.Ui32(int32(25)) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+24)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v54-int32(16)))) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v54-int32(20)))) = v41
	v80 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v54-v80))) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v54-int32(28)))) = v41
	v89 = v45&int32(4) | v80
	v90 = v53 - v89
	if base.Ui32(v90) < base.Ui32(int32(32)) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v95 = base.I64_extend_i32_u(v41) * int64(4294967297)
	v98 = v89 + v45
	v99 = v90
	goto L10
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v98)+24)) = v95
	*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = v95
	*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = v95
	*(*int64)(unsafe.Add(mBase, uint32(v98))) = v95
	v107 = int32(32)
	v110 = v99 - v107
	if base.Ui32(int32(31)) < base.Ui32(v110) {
		v98 = v98 + v107
		v99 = v110
		goto L10
	} else {
		goto L12
	}
L11:
	;
	goto L2
L12:
	;
	goto L11
L13:
	;
	m.G0 = v9 - int32(-8192)
	return v182
L14:
	;
	v131 = F_recvfrom(m, l0, v9)
	mBase = m.M
	if v131 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v182 = int32(-1)
	goto L13
L16:
	;
	v136 = v131
	goto L17
L17:
	;
	if base.Ui32(int32(16)) <= base.Ui32(v136) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v182 = int32(0)
	goto L13
L19:
	;
	goto L18
L20:
	;
	v144 = v9
	goto L23
L21:
	;
	goto L22
L22:
	;
	v169 = F_recvfrom(m, l0, v9)
	mBase = m.M
	if int32(0) < v169 {
		v136 = v169
		goto L17
	} else {
		goto L28
	}
L23:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+4)))
	switch v150 - int32(2) {
	case 0:
		v182 = int32(-1)
		goto L13
	case 1:
		goto L19
	default:
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	v153 = F_netlink_msg_to_ifaddr(m, l3, v144)
	mBase = m.M
	if v153 != 0 {
		v182 = v153
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v159 = v144 + (v154+int32(3))&int32(-4)
	if base.Ui32(int32(15)) < base.Ui32(v136+v9-v159) {
		v144 = v159
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	goto L15
}
func F_namele(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.I64_extend_i32_u(base.B2i32(v61 <= int32(0)))
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	v55 = F_strlen(m, v5)
	mBase = m.M
	v56 = F_strlen(m, v4)
	mBase = m.M
	v57 = F_varstr_cmp(m, v5, v55, v4, v56, v6)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v61 = v46 - v47
	goto L1
L7:
	;
	goto L8
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L13
L10:
	;
	v42 = v4
	v46 = int32(0)
	goto L11
L11:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	goto L5
L12:
	;
	v42 = v37
	v46 = v39
	goto L11
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(v19 != v21)|base.B2i32(v21 == int32(0)) != 0 {
		v37 = v17
		v39 = v19
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v37 = v31
	v39 = int32(0)
	goto L12
L15:
	;
	v27 = v18 - int32(1)
	if v27 == int32(0) {
		v37 = v17
		v39 = v19
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v30 = int32(1)
	v31 = v17 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v32 != 0 {
		v16 = v16 + v30
		v17 = v31
		v18 = v27
		v19 = v32
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	return int64(0)
L19:
	;
	v61 = v57
	goto L1
}
func F_ndistinct_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14253(m, l0, l1, int32(_a_F_ndistinct_array_element_start_0), int32(383), int32(_a_F_ndistinct_array_element_start_1), int32(_a_F_ndistinct_array_element_start_2), int32(376), int32(_a_F_ndistinct_array_element_start_3), int32(366))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_ndistinct_object_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 == int32(2) {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v13 == int32(0) {
			v16 = int32(23)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = F_errsave_start(m, v17)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					v209 = v16
					m.G0 = v8 + int32(112)
					return v209
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v27
						F_errmsg(m, int32(_a_F_ndistinct_object_end_0), v8+int32(80))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = int32(_a_F_ndistinct_object_end_1)
							v39 = F_errdetail(m, int32(_a_F_ndistinct_object_end_2), v8-int32(-64))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v17, int32(_a_F_ndistinct_object_end_3), int32(148), int32(_a_F_ndistinct_object_end_4))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v209 = v16
									m.G0 = v8 + int32(112)
									return v209
								}
							}
						}
					}
				}
			}
		} else {
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
			if v46 == int32(0) {
				v49 = int32(23)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v51 = F_errsave_start(m, v50)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					if v51 == int32(0) {
						v209 = v49
						m.G0 = v8 + int32(112)
						return v209
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v58
							F_errmsg(m, int32(_a_F_ndistinct_object_end_0), v8+int32(48))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_ndistinct_object_end_5)
								v70 = F_errdetail(m, int32(_a_F_ndistinct_object_end_2), v8+int32(32))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, v50, int32(_a_F_ndistinct_object_end_3), int32(158), int32(_a_F_ndistinct_object_end_4))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v209 = v49
										m.G0 = v8 + int32(112)
										return v209
									}
								}
							}
						}
					}
				}
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v77 != 0 {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
					if base.Ui32(int32(-8)) < base.Ui32(v78-int32(9)) {
						v113 = F_palloc(m, int32(16))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v78
							v118 = F_palloc0(m, v78<<(uint(int32(1))%32))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v118
								v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								*(*float64)(unsafe.Add(mBase, uint32(v113))) = base.F64_convert_i32_s(v121)
								v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
								*(*uint16)(unsafe.Add(mBase, uint32(v118))) = uint16(v126)
								if v78 == int32(1) {
								} else {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
									v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
									*(*uint16)(unsafe.Add(mBase, uint32(v130)+2)) = uint16(v133)
									if v78 == int32(2) {
									} else {
										v137 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
										v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
										*(*uint16)(unsafe.Add(mBase, uint32(v137)+4)) = uint16(v140)
										if v78 == int32(3) {
										} else {
											v144 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
											v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+12))
											v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
											*(*uint16)(unsafe.Add(mBase, uint32(v144)+6)) = uint16(v147)
											if v78 == int32(4) {
											} else {
												v151 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
												v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
												v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
												*(*uint16)(unsafe.Add(mBase, uint32(v151)+8)) = uint16(v154)
												if v78 == int32(5) {
												} else {
													v158 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
													v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
													v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+20))
													*(*uint16)(unsafe.Add(mBase, uint32(v158)+10)) = uint16(v161)
													if v78 == int32(6) {
													} else {
														v165 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
														v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
														v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
														*(*uint16)(unsafe.Add(mBase, uint32(v165)+12)) = uint16(v168)
														if v78 == int32(7) {
														} else {
															v172 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
															v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
															v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+28))
															*(*uint16)(unsafe.Add(mBase, uint32(v172)+14)) = uint16(v175)
														}
													}
												}
											}
										}
									}
								}
								v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v196 = F_lappend(m, v195, v113)
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v196
									v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									F_list_free(m, v199)
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return int32(0)
									} else {
										v202 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v202)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
										v209 = v202
										m.G0 = v8 + int32(112)
										return v209
									}
								}
							}
						}
					} else {
						v84 = int32(23)
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v86 = F_errsave_start(m, v85)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							if v86 == int32(0) {
								v209 = v84
								m.G0 = v8 + int32(112)
								return v209
							} else {
								F_errcode(m, int32(33685634))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v93
									F_errmsg(m, int32(_a_F_ndistinct_object_end_0), v8+int32(16))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(34359738370)
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_ndistinct_object_end_1)
										v105 = F_errdetail(m, int32(_a_F_ndistinct_object_end_6), v8)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, v85, int32(_a_F_ndistinct_object_end_3), int32(173), int32(_a_F_ndistinct_object_end_4))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int32(0)
											} else {
												v209 = v84
												m.G0 = v8 + int32(112)
												return v209
											}
										}
									}
								}
							}
						}
					}
				} else {
					v84 = int32(23)
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v86 = F_errsave_start(m, v85)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						if v86 == int32(0) {
							v209 = v84
							m.G0 = v8 + int32(112)
							return v209
						} else {
							F_errcode(m, int32(33685634))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v93
								F_errmsg(m, int32(_a_F_ndistinct_object_end_0), v8+int32(16))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(34359738370)
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_ndistinct_object_end_1)
									v105 = F_errdetail(m, int32(_a_F_ndistinct_object_end_6), v8)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, v85, int32(_a_F_ndistinct_object_end_3), int32(173), int32(_a_F_ndistinct_object_end_4))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											v209 = v84
											m.G0 = v8 + int32(112)
											return v209
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v180 = m.ExcPending
		if v180 != 0 {
			return int32(0)
		} else {
			v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = v181
			*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = int32(_a_F_ndistinct_object_end_7)
			F_errmsg_internal(m, int32(_a_F_ndistinct_object_end_8), v8+int32(96))
			mBase = m.M
			v189 = m.ExcPending
			if v189 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_ndistinct_object_end_3), int32(140), int32(_a_F_ndistinct_object_end_4))
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_ndistinct_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v9 = int32(_a_F_ndistinct_object_field_start_0)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ndistinct_object_field_start[0])))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v7 + int32(96)
	return v169
L2:
	;
	if v33-v34 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	goto L2
L4:
	;
	v18 = l1
	v19 = v9
	goto L5
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L3
	} else {
		goto L7
	}
L6:
	;
	v33 = v23
	v34 = v22
	goto L3
L7:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v38 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v74 = int32(_a_F_ndistinct_object_field_start_1)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ndistinct_object_field_start[1])))
	if base.B2i32(v77 == int32(0))|base.B2i32(v77 != v80) != 0 {
		v98 = v77
		v99 = v80
		goto L23
	} else {
		goto L24
	}
L12:
	;
	v41 = int32(23)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = F_errsave_start(m, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(3)
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
	v169 = int32(0)
	goto L1
L15:
	;
	return int32(0)
L16:
	;
	if v43 == int32(0) {
		v169 = v41
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v52
	F_errmsg(m, int32(_a_F_ndistinct_object_field_start_2), v7+int32(16))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_ndistinct_object_field_start_0)
	v62 = F_errdetail(m, int32(_a_F_ndistinct_object_field_start_3), v7)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	F_errsave_finish(m, v42, int32(_a_F_ndistinct_object_field_start_4), int32(313), int32(_a_F_ndistinct_object_field_start_5))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v169 = v41
	goto L1
L22:
	;
	if v98-v99 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	v83 = l1
	v84 = v74
	goto L25
L25:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v88 == int32(0) {
		v98 = v88
		v99 = v87
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v98 = v88
	v99 = v87
	goto L23
L27:
	;
	v91 = int32(1)
	if v88 == v87 {
		v83 = v83 + v91
		v84 = v84 + v91
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v103 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v139 = int32(23)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v141 = F_errsave_start(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L15
	} else {
		goto L41
	}
L32:
	;
	v106 = int32(23)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v108 = F_errsave_start(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L15
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
	v136 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v136)
	v169 = int32(0)
	goto L1
L35:
	;
	if v108 == int32(0) {
		v169 = v106
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v115
	F_errmsg(m, int32(_a_F_ndistinct_object_field_start_2), v7+int32(48))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_ndistinct_object_field_start_1)
	v127 = F_errdetail(m, int32(_a_F_ndistinct_object_field_start_3), v7+int32(32))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, v107, int32(_a_F_ndistinct_object_field_start_4), int32(329), int32(_a_F_ndistinct_object_field_start_5))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	v169 = v106
	goto L1
L41:
	;
	if v141 == int32(0) {
		v169 = v139
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L15
	} else {
		goto L43
	}
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = v148
	F_errmsg(m, int32(_a_F_ndistinct_object_field_start_2), v7+int32(80))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = int32(_a_F_ndistinct_object_field_start_1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = int32(_a_F_ndistinct_object_field_start_0)
	v162 = F_errdetail(m, int32(_a_F_ndistinct_object_field_start_6), v7-int32(-64))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	F_errsave_finish(m, v140, int32(_a_F_ndistinct_object_field_start_4), int32(342), int32(_a_F_ndistinct_object_field_start_5))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	v169 = v139
	goto L1
}
func F_nepali_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8
	v14 = F_find_among_b(m, l0, int32(_a_F_nepali_UTF_8_stem_0), int32(17), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v236
L2:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
	goto L20
L3:
	;
	return int32(0)
L4:
	;
	if v14 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v20
	switch v14 - int32(1) {
	case 0:
		goto L7
	case 1:
		goto L6
	default:
		goto L2
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = int32(3)
	v30 = int32(0)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v32-v33 < v28 {
		v43 = v30
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v24 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v24 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v236 = v24
	goto L1
L9:
	;
	if v43 != 0 {
		goto L2
	} else {
		goto L13
	}
L10:
	;
	goto L9
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = F_memcmp(m, v36+v32-v28, int32(_a_F_nepali_UTF_8_stem_1), v28)
	mBase = m.M
	if v39 != 0 {
		v43 = v30
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32 - v28
	v43 = int32(1)
	goto L10
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v45 = v27 - v20
	v46 = v44 - v45
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46
	v48 = int32(3)
	v50 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v46-v53 < v48 {
		v63 = v50
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v63 != 0 {
		goto L2
	} else {
		goto L18
	}
L15:
	;
	goto L14
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = F_memcmp(m, v56+v46-v48, int32(_a_F_nepali_UTF_8_stem_2), v48)
	mBase = m.M
	if v59 != 0 {
		v63 = v50
		goto L15
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46 - v48
	v63 = int32(1)
	goto L15
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64 - v45
	v67 = F_slice_del(m, l0)
	mBase = m.M
	if v67 < int32(0) {
		v236 = v67
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L2
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v79-int32(2) <= v82 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v232
	v236 = int32(1)
	goto L1
L22:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v219 = v217 + (v79 - v81)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v219
	v225 = F_find_among_b(m, l0, int32(_a_F_nepali_UTF_8_stem_3), int32(91), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L3
	} else {
		goto L57
	}
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = int32(1)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v79-v88))))
	if base.B2i32(v90&int32(224) != int32(128))|base.B2i32(v88<<(uint(v90)%32)&int32(262) == int32(0)) != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v105 = F_find_among_b(m, l0, int32(_a_F_nepali_UTF_8_stem_4), int32(3), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	if v105 == int32(0) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v109
	switch v105 - int32(1) {
	case 0:
		goto L28
	case 1:
		goto L27
	default:
		goto L22
	}
L27:
	;
	v194 = int32(9)
	v196 = int32(0)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v198-v199 < v194 {
		v209 = v196
		goto L52
	} else {
		goto L53
	}
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v114 = int32(6)
	v116 = int32(0)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v118-v119 < v114 {
		v129 = v116
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v191 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v191 {
		goto L22
	} else {
		goto L50
	}
L30:
	;
	if v129 != 0 {
		goto L29
	} else {
		goto L34
	}
L31:
	;
	goto L30
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v125 = F_memcmp(m, v122+v118-v114, int32(_a_F_nepali_UTF_8_stem_5), v114)
	mBase = m.M
	if v125 != 0 {
		v129 = v116
		goto L31
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v118 - v114
	v129 = int32(1)
	goto L31
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v131 = v113 - v109
	v132 = v130 - v131
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v132
	v134 = int32(6)
	v136 = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v132-v139 < v134 {
		v149 = v136
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v149 != 0 {
		goto L29
	} else {
		goto L39
	}
L36:
	;
	goto L35
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v145 = F_memcmp(m, v142+v132-v134, int32(_a_F_nepali_UTF_8_stem_6), v134)
	mBase = m.M
	if v145 != 0 {
		v149 = v136
		goto L36
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v132 - v134
	v149 = int32(1)
	goto L36
L39:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v151 = v150 - v131
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v151
	v153 = int32(6)
	v155 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v151-v158 < v153 {
		v168 = v155
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v168 != 0 {
		goto L29
	} else {
		goto L44
	}
L41:
	;
	goto L40
L42:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v164 = F_memcmp(m, v161+v151-v153, int32(_a_F_nepali_UTF_8_stem_7), v153)
	mBase = m.M
	if v164 != 0 {
		v168 = v155
		goto L41
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v151 - v153
	v168 = int32(1)
	goto L41
L44:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v170 = v169 - v131
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170
	v172 = int32(6)
	v174 = int32(0)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v170-v177 < v172 {
		v187 = v174
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v187 == int32(0) {
		goto L22
	} else {
		goto L49
	}
L46:
	;
	goto L45
L47:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v183 = F_memcmp(m, v180+v170-v172, int32(_a_F_nepali_UTF_8_stem_8), v172)
	mBase = m.M
	if v183 != 0 {
		v187 = v174
		goto L46
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170 - v172
	v187 = int32(1)
	goto L46
L49:
	;
	goto L29
L50:
	;
	v236 = v191
	goto L1
L51:
	;
	if v209 == int32(0) {
		goto L22
	} else {
		goto L55
	}
L52:
	;
	goto L51
L53:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v205 = F_memcmp(m, v202+v198-v194, int32(_a_F_nepali_UTF_8_stem_9), v194)
	mBase = m.M
	if v205 != 0 {
		v209 = v196
		goto L52
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v198 - v194
	v209 = int32(1)
	goto L52
L55:
	;
	v212 = F_slice_del(m, l0)
	mBase = m.M
	if v212 < int32(0) {
		v236 = v212
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L22
L57:
	;
	if v225 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v227
	v229 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v229 {
		goto L20
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L21
L61:
	;
	v236 = v229
	goto L1
}
func F_netlink_msg_to_ifaddr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v756 int32
	_ = v756
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v817 int32
	_ = v817
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v832 int64
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v908 int32
	_ = v908
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1113 int64
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1130 int32
	_ = v1130
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1191 int32
	_ = v1191
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1206 int64
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1238 int32
	_ = v1238
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v14 == int32(16) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	v92 = int32(1)
	goto L23
L3:
	;
	v17 = int32(164)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v18&int32(-4) == int32(32) {
		v81 = v18
		v82 = v17
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0+v53&int32(63)<<(uint(int32(2))%32))+8))
	if v59 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v28 = l1 + int32(32)
	goto L7
L7:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+2)))
	if v40 == int32(7) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v81 = v49
	v82 = v17
	goto L2
L9:
	;
	v81 = v28
	v82 = v39 + int32(160)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v45 = int32(3)
	v49 = v28 + (v39+v45)&int32(_a_F_netlink_msg_to_ifaddr_0)
	if base.Ui32(v45) < base.Ui32(l1+v18-v49) {
		v28 = v49
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	v64 = v59
	goto L14
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64)+140))
	if v53 == v75 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L1
L16:
	;
	v81 = v64
	v82 = int32(164)
	goto L2
L17:
	;
	goto L18
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	if v78 != 0 {
		v64 = v78
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	if v113 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	goto L20
L22:
	;
	v113 = F_emscripten_builtin_malloc(m, v112)
	mBase = m.M
	if v113 == int32(0) {
		goto L21
	} else {
		goto L28
	}
L23:
	;
	v100 = base.I64_extend_i32_u(v92) * base.I64_extend_i32_u(v82)
	v101 = base.I32_wrap_i64(v100)
	if base.Ui32(v92|v82) < base.Ui32(int32(_a_F_netlink_msg_to_ifaddr_1)) {
		v112 = v101
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if base.I32_wrap_i64(int64(base.Ui64(v100)>>(uint(int64(32))%64))) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v109 = int32(-1)
	goto L27
L26:
	;
	v109 = v101
	goto L27
L27:
	;
	v112 = v109
	goto L22
L28:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113-int32(4)))))
	if v118&int32(3) == int32(0) {
		goto L21
	} else {
		goto L29
	}
L29:
	;
	F___memset(m, v113, int32(0), v112)
	mBase = m.M
	goto L21
L30:
	;
	return int32(-1)
L31:
	;
	goto L32
L32:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v129 == int32(16) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v1290 != 0 {
		goto L314
	} else {
		goto L315
	}
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+140)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v136&int32(-4) != int32(32) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v596
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v598
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v600&int32(-4) != int32(24) {
		goto L152
	} else {
		goto L153
	}
L37:
	;
	v142 = v113 + int32(144)
	v143 = int32(32)
	v144 = v113 + v143
	v148 = v113 + int32(104)
	v152 = v113 + int32(164)
	v157 = l1 + v143
	goto L40
L38:
	;
	goto L39
L39:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v584 == int32(0) {
		goto L33
	} else {
		goto L151
	}
L40:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+2)))
	switch v168 - int32(1) {
	case 0:
		goto L45
	case 1:
		goto L44
	case 2:
		goto L46
	default:
		goto L42
	case 6:
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v562 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157))))
	v563 = int32(3)
	v567 = v157 + (v562+v563)&int32(_a_F_netlink_msg_to_ifaddr_0)
	if base.Ui32(v563) < base.Ui32(l1+v560-v567) {
		v157 = v567
		goto L40
	} else {
		goto L150
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+24)) = v152
	v385 = int32(4)
	v386 = v157 + v385
	v387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157))))
	v389 = v387 - v385
	if base.Ui32(int32(512)) <= base.Ui32(v389) {
		goto L104
	} else {
		goto L105
	}
L44:
	;
	v366 = int32(4)
	v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157))))
	v370 = v368 - v366
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v370) <= base.Ui32(int32(24)) {
		goto L100
	} else {
		goto L101
	}
L45:
	;
	v348 = int32(4)
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157))))
	v352 = v350 - v348
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v352) <= base.Ui32(int32(24)) {
		goto L96
	} else {
		goto L97
	}
L46:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157))))
	v173 = v171 - int32(4)
	if base.Ui32(int32(16)) < base.Ui32(v173) {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v177 = v157 + int32(4)
	if base.Ui32(int32(512)) <= base.Ui32(v173) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v142
	goto L42
L49:
	;
	if v173 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v184 = v142 + v173
	if (v142^v177)&int32(3) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	base.MemoryCopy(m, v142, v177, v173)
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L48
L55:
	;
	if base.Ui32(v316) < base.Ui32(v184) {
		goto L89
	} else {
		goto L90
	}
L56:
	;
	if v142&int32(3) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(v184) < base.Ui32(int32(4)) {
		goto L80
	} else {
		goto L81
	}
L59:
	;
	v220 = v184 & int32(-4)
	if base.Ui32(v184) < base.Ui32(int32(64)) {
		v270 = v214
		v271 = v215
		goto L70
	} else {
		goto L71
	}
L60:
	;
	v214 = v177
	v215 = v142
	goto L59
L61:
	;
	goto L62
L62:
	;
	if v173 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v214 = v177
	v215 = v142
	goto L59
L64:
	;
	goto L65
L65:
	;
	v197 = v177
	v198 = v142
	goto L66
L66:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v202)
	v204 = int32(1)
	v205 = v197 + v204
	v207 = v198 + v204
	if v207&int32(3) == int32(0) {
		v214 = v205
		v215 = v207
		goto L59
	} else {
		goto L68
	}
L67:
	;
	v214 = v205
	v215 = v207
	goto L59
L68:
	;
	if base.Ui32(v207) < base.Ui32(v184) {
		v197 = v205
		v198 = v207
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	if base.Ui32(v220) <= base.Ui32(v271) {
		v315 = v270
		v316 = v271
		goto L55
	} else {
		goto L76
	}
L71:
	;
	v224 = v220 + int32(-64)
	if base.Ui32(v224) < base.Ui32(v215) {
		v270 = v214
		v271 = v215
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v227 = v214
	v228 = v215
	goto L73
L73:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+8)) = v236
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+12)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v227)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+16)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v227)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+20)) = v242
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v227)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+24)) = v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v227)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+28)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v227)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+32)) = v248
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v227)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+36)) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v227)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+40)) = v252
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v227)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+44)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v227)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+48)) = v256
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v227)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+52)) = v258
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v227)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+56)) = v260
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v227)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+60)) = v262
	v264 = int32(-64)
	v265 = v227 - v264
	v267 = v228 - v264
	if base.Ui32(v267) <= base.Ui32(v224) {
		v227 = v265
		v228 = v267
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v270 = v265
	v271 = v267
	goto L70
L75:
	;
	goto L74
L76:
	;
	v277 = v270
	v278 = v271
	goto L77
L77:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	*(*int32)(unsafe.Add(mBase, uint32(v278))) = v282
	v284 = int32(4)
	v285 = v277 + v284
	v287 = v278 + v284
	if base.Ui32(v287) < base.Ui32(v220) {
		v277 = v285
		v278 = v287
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v315 = v285
	v316 = v287
	goto L55
L79:
	;
	goto L78
L80:
	;
	v315 = v177
	v316 = v142
	goto L55
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(v173) < base.Ui32(int32(4)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v315 = v177
	v316 = v142
	goto L55
L84:
	;
	goto L85
L85:
	;
	v296 = v177
	v297 = v142
	goto L86
L86:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	*(*uint8)(unsafe.Add(mBase, uint32(v297))) = uint8(v301)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v297)+1)) = uint8(v303)
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v297)+2)) = uint8(v305)
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v297)+3)) = uint8(v307)
	v309 = int32(4)
	v310 = v296 + v309
	v312 = v297 + v309
	if base.Ui32(v312) <= base.Ui32(v184-int32(4)) {
		v296 = v310
		v297 = v312
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v315 = v310
	v316 = v312
	goto L55
L88:
	;
	goto L87
L89:
	;
	v322 = v315
	v323 = v316
	goto L92
L90:
	;
	goto L91
L91:
	;
	goto L48
L92:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v327)
	v329 = int32(1)
	v332 = v323 + v329
	if v332 != v184 {
		v322 = v322 + v329
		v323 = v332
		goto L92
	} else {
		goto L94
	}
L93:
	;
	goto L91
L94:
	;
	goto L93
L95:
	;
	goto L42
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+11)) = uint8(v352)
	*(*uint16)(unsafe.Add(mBase, uint32(v144)+8)) = uint16(v354)
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v353
	v360 = int32(17)
	*(*uint16)(unsafe.Add(mBase, uint32(v144))) = uint16(v360)
	v364 = F___memcpy(m, v113+int32(44), v157+v348, v352)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v113+int32(12)))) = v144
	goto L98
L97:
	;
	goto L98
L98:
	;
	goto L95
L99:
	;
	goto L42
L100:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+11)) = uint8(v370)
	*(*uint16)(unsafe.Add(mBase, uint32(v148)+8)) = uint16(v372)
	*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v371
	v378 = int32(17)
	*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v378)
	v382 = F___memcpy(m, v113+int32(116), v157+v366, v370)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v113+int32(20)))) = v148
	goto L102
L101:
	;
	goto L102
L102:
	;
	goto L99
L103:
	;
	goto L42
L104:
	;
	if v389 != 0 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	v396 = v152 + v389
	if (v152^v386)&int32(3) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	base.MemoryCopy(m, v152, v386, v389)
	goto L109
L108:
	;
	goto L109
L109:
	;
	goto L103
L110:
	;
	if base.Ui32(v528) < base.Ui32(v396) {
		goto L144
	} else {
		goto L145
	}
L111:
	;
	if v152&int32(3) == int32(0) {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	goto L113
L113:
	;
	if base.Ui32(v396) < base.Ui32(int32(4)) {
		goto L135
	} else {
		goto L136
	}
L114:
	;
	v432 = v396 & int32(-4)
	if base.Ui32(v396) < base.Ui32(int32(64)) {
		v482 = v426
		v483 = v427
		goto L125
	} else {
		goto L126
	}
L115:
	;
	v426 = v386
	v427 = v152
	goto L114
L116:
	;
	goto L117
L117:
	;
	if v389 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v426 = v386
	v427 = v152
	goto L114
L119:
	;
	goto L120
L120:
	;
	v409 = v386
	v410 = v152
	goto L121
L121:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	*(*uint8)(unsafe.Add(mBase, uint32(v410))) = uint8(v414)
	v416 = int32(1)
	v417 = v409 + v416
	v419 = v410 + v416
	if v419&int32(3) == int32(0) {
		v426 = v417
		v427 = v419
		goto L114
	} else {
		goto L123
	}
L122:
	;
	v426 = v417
	v427 = v419
	goto L114
L123:
	;
	if base.Ui32(v419) < base.Ui32(v396) {
		v409 = v417
		v410 = v419
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	if base.Ui32(v432) <= base.Ui32(v483) {
		v527 = v482
		v528 = v483
		goto L110
	} else {
		goto L131
	}
L126:
	;
	v436 = v432 + int32(-64)
	if base.Ui32(v436) < base.Ui32(v427) {
		v482 = v426
		v483 = v427
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v439 = v426
	v440 = v427
	goto L128
L128:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v439)))
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v444
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+4)) = v446
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v439)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+8)) = v448
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v439)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+12)) = v450
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v439)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+16)) = v452
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v439)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+20)) = v454
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v439)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+24)) = v456
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v439)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+28)) = v458
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v439)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+32)) = v460
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v439)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+36)) = v462
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v439)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+40)) = v464
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v439)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+44)) = v466
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v439)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+48)) = v468
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v439)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+52)) = v470
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v439)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+56)) = v472
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v439)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+60)) = v474
	v476 = int32(-64)
	v477 = v439 - v476
	v479 = v440 - v476
	if base.Ui32(v479) <= base.Ui32(v436) {
		v439 = v477
		v440 = v479
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v482 = v477
	v483 = v479
	goto L125
L130:
	;
	goto L129
L131:
	;
	v489 = v482
	v490 = v483
	goto L132
L132:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	*(*int32)(unsafe.Add(mBase, uint32(v490))) = v494
	v496 = int32(4)
	v497 = v489 + v496
	v499 = v490 + v496
	if base.Ui32(v499) < base.Ui32(v432) {
		v489 = v497
		v490 = v499
		goto L132
	} else {
		goto L134
	}
L133:
	;
	v527 = v497
	v528 = v499
	goto L110
L134:
	;
	goto L133
L135:
	;
	v527 = v386
	v528 = v152
	goto L110
L136:
	;
	goto L137
L137:
	;
	if base.Ui32(v389) < base.Ui32(int32(4)) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v527 = v386
	v528 = v152
	goto L110
L139:
	;
	goto L140
L140:
	;
	v508 = v386
	v509 = v152
	goto L141
L141:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
	*(*uint8)(unsafe.Add(mBase, uint32(v509))) = uint8(v513)
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v509)+1)) = uint8(v515)
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v509)+2)) = uint8(v517)
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v509)+3)) = uint8(v519)
	v521 = int32(4)
	v522 = v508 + v521
	v524 = v509 + v521
	if base.Ui32(v524) <= base.Ui32(v396-int32(4)) {
		v508 = v522
		v509 = v524
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v527 = v522
	v528 = v524
	goto L110
L143:
	;
	goto L142
L144:
	;
	v534 = v527
	v535 = v528
	goto L147
L145:
	;
	goto L146
L146:
	;
	goto L103
L147:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v539)
	v541 = int32(1)
	v544 = v535 + v541
	if v544 != v396 {
		v534 = v534 + v541
		v535 = v544
		goto L147
	} else {
		goto L149
	}
L148:
	;
	goto L146
L149:
	;
	goto L148
L150:
	;
	goto L41
L151:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v113)+140))
	v592 = l0 + v587&int32(63)<<(uint(int32(2))%32)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v592)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+28)) = v593
	*(*int32)(unsafe.Add(mBase, uint32(v592)+8)) = v113
	goto L33
L152:
	;
	v606 = v113 + int32(20)
	v608 = v113 + int32(32)
	v610 = v113 + int32(104)
	v612 = v113 + int32(12)
	v614 = v113 + int32(144)
	v619 = l1 + int32(24)
	goto L155
L153:
	;
	goto L154
L154:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	if v1100 == int32(0) {
		goto L33
	} else {
		goto L281
	}
L155:
	;
	v630 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v619)+2)))
	switch v630 - int32(1) {
	case 0:
		goto L161
	case 1:
		goto L159
	case 2:
		goto L158
	case 3:
		goto L160
	default:
		goto L157
	}
L156:
	;
	goto L154
L157:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1078 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v619))))
	v1079 = int32(3)
	v1083 = v619 + (v1078+v1079)&int32(_a_F_netlink_msg_to_ifaddr_0)
	if base.Ui32(v1079) < base.Ui32(l1+v1076-v1083) {
		v619 = v1083
		goto L155
	} else {
		goto L280
	}
L158:
	;
	v895 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v619))))
	v897 = v895 - int32(4)
	if base.Ui32(int32(16)) < base.Ui32(v897) {
		goto L157
	} else {
		goto L232
	}
L159:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	if v744 != 0 {
		goto L204
	} else {
		goto L205
	}
L160:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v706 = int32(4)
	v707 = v619 + v706
	v708 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v619))))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v705 != int32(10) {
		goto L194
	} else {
		goto L195
	}
L161:
	;
	v633 = int32(4)
	v634 = v619 + v633
	v635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v619))))
	v637 = v635 - v633
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	if v640 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	if v639 != int32(10) {
		goto L168
	} else {
		goto L169
	}
L163:
	;
	goto L164
L164:
	;
	if v639 != int32(10) {
		goto L181
	} else {
		goto L182
	}
L165:
	;
	goto L157
L166:
	;
	goto L165
L167:
	;
	if base.Ui32(v637) < base.Ui32(v665) {
		goto L166
	} else {
		goto L177
	}
L168:
	;
	if v639 != int32(2) {
		goto L166
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v651 = v113 + int32(112)
	v652 = int32(16)
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	switch v653 - int32(254) {
	case 0:
		goto L174
	case 1:
		goto L173
	default:
		v665 = v652
		v666 = v651
		goto L167
	}
L171:
	;
	v665 = int32(4)
	v666 = v113 + int32(108)
	goto L167
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610)+24)) = v638
	v665 = v652
	v666 = v651
	goto L167
L173:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+1)))
	if v659&int32(15) != int32(2) {
		v665 = v652
		v666 = v651
		goto L167
	} else {
		goto L176
	}
L174:
	;
	v656 = int32(*(*int8)(unsafe.Add(mBase, uint32(v634)+1)))
	if v656 < int32(-64) {
		goto L172
	} else {
		goto L175
	}
L175:
	;
	v665 = v652
	v666 = v651
	goto L167
L176:
	;
	goto L172
L177:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v610))) = uint16(v639)
	v669 = F___memcpy(m, v666, v634, v665)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v606))) = v610
	goto L166
L178:
	;
	goto L157
L179:
	;
	goto L178
L180:
	;
	if base.Ui32(v637) < base.Ui32(v697) {
		goto L179
	} else {
		goto L190
	}
L181:
	;
	if v639 != int32(2) {
		goto L179
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v683 = v113 + int32(40)
	v684 = int32(16)
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	switch v685 - int32(254) {
	case 0:
		goto L187
	case 1:
		goto L186
	default:
		v697 = v684
		v698 = v683
		goto L180
	}
L184:
	;
	v697 = int32(4)
	v698 = v113 + int32(36)
	goto L180
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v608)+24)) = v638
	v697 = v684
	v698 = v683
	goto L180
L186:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634)+1)))
	if v691&int32(15) != int32(2) {
		v697 = v684
		v698 = v683
		goto L180
	} else {
		goto L189
	}
L187:
	;
	v688 = int32(*(*int8)(unsafe.Add(mBase, uint32(v634)+1)))
	if v688 < int32(-64) {
		goto L185
	} else {
		goto L188
	}
L188:
	;
	v697 = v684
	v698 = v683
	goto L180
L189:
	;
	goto L185
L190:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v608))) = uint16(v639)
	v701 = F___memcpy(m, v698, v634, v697)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v612))) = v608
	goto L179
L191:
	;
	goto L157
L192:
	;
	goto L191
L193:
	;
	if base.Ui32(v708-v706) < base.Ui32(v736) {
		goto L192
	} else {
		goto L203
	}
L194:
	;
	if v705 != int32(2) {
		goto L192
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v722 = v113 + int32(112)
	v723 = int32(16)
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707))))
	switch v724 - int32(254) {
	case 0:
		goto L200
	case 1:
		goto L199
	default:
		v736 = v723
		v737 = v722
		goto L193
	}
L197:
	;
	v736 = int32(4)
	v737 = v113 + int32(108)
	goto L193
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610)+24)) = v711
	v736 = v723
	v737 = v722
	goto L193
L199:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707)+1)))
	if v730&int32(15) != int32(2) {
		v736 = v723
		v737 = v722
		goto L193
	} else {
		goto L202
	}
L200:
	;
	v727 = int32(*(*int8)(unsafe.Add(mBase, uint32(v707)+1)))
	if v727 < int32(-64) {
		goto L198
	} else {
		goto L201
	}
L201:
	;
	v736 = v723
	v737 = v722
	goto L193
L202:
	;
	goto L198
L203:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v610))) = uint16(v705)
	v740 = F___memcpy(m, v737, v707, v736)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v606))) = v610
	goto L192
L204:
	;
	v745 = int32(36)
	base.MemoryCopy(m, v610, v608, v745)
	*(*int32)(unsafe.Add(mBase, uint32(v606))) = v610
	v748 = int32(0)
	goto L209
L205:
	;
	goto L206
L206:
	;
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v857 = int32(4)
	v858 = v619 + v857
	v859 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v619))))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v856 != int32(10) {
		goto L222
	} else {
		goto L223
	}
L207:
	;
	goto L206
L208:
	;
	goto L207
L209:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v608))) = uint8(v748)
	v756 = v113 + int32(68)
	*(*uint8)(unsafe.Add(mBase, uint32(v756-int32(1)))) = uint8(v748)
	goto L210
L210:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v608)+2)) = uint8(v748)
	*(*uint8)(unsafe.Add(mBase, uint32(v608)+1)) = uint8(v748)
	*(*uint8)(unsafe.Add(mBase, uint32(v756-int32(3)))) = uint8(v748)
	*(*uint8)(unsafe.Add(mBase, uint32(v756-int32(2)))) = uint8(v748)
	goto L211
L211:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v608)+3)) = uint8(v748)
	*(*uint8)(unsafe.Add(mBase, uint32(v756-int32(4)))) = uint8(v748)
	goto L212
L212:
	;
	v778 = int32(0)
	v781 = (v778 - v608) & int32(3)
	v782 = v608 + v781
	*(*int32)(unsafe.Add(mBase, uint32(v782))) = v778
	v790 = (v745 - v781) & int32(-4)
	v791 = v782 + v790
	*(*int32)(unsafe.Add(mBase, uint32(v791-int32(4)))) = v778
	if base.Ui32(v790) < base.Ui32(int32(9)) {
		goto L208
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v782)+8)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v782)+4)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v791-int32(8)))) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v791-int32(12)))) = v778
	if base.Ui32(v790) < base.Ui32(int32(25)) {
		goto L208
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v782)+24)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v782)+20)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v782)+16)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v782)+12)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v791-int32(16)))) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v791-int32(20)))) = v778
	v817 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v791-v817))) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v791-int32(28)))) = v778
	v826 = v782&int32(4) | v817
	v827 = v790 - v826
	if base.Ui32(v827) < base.Ui32(int32(32)) {
		goto L208
	} else {
		goto L215
	}
L215:
	;
	v832 = base.I64_extend_i32_u(v778) * int64(4294967297)
	v835 = v826 + v782
	v836 = v827
	goto L216
L216:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v835)+24)) = v832
	*(*int64)(unsafe.Add(mBase, uint32(v835)+16)) = v832
	*(*int64)(unsafe.Add(mBase, uint32(v835)+8)) = v832
	*(*int64)(unsafe.Add(mBase, uint32(v835))) = v832
	v844 = int32(32)
	v847 = v836 - v844
	if base.Ui32(int32(31)) < base.Ui32(v847) {
		v835 = v835 + v844
		v836 = v847
		goto L216
	} else {
		goto L218
	}
L217:
	;
	goto L208
L218:
	;
	goto L217
L219:
	;
	goto L157
L220:
	;
	goto L219
L221:
	;
	if base.Ui32(v859-v857) < base.Ui32(v887) {
		goto L220
	} else {
		goto L231
	}
L222:
	;
	if v856 != int32(2) {
		goto L220
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v873 = v113 + int32(40)
	v874 = int32(16)
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v858))))
	switch v875 - int32(254) {
	case 0:
		goto L228
	case 1:
		goto L227
	default:
		v887 = v874
		v888 = v873
		goto L221
	}
L225:
	;
	v887 = int32(4)
	v888 = v113 + int32(36)
	goto L221
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v608)+24)) = v862
	v887 = v874
	v888 = v873
	goto L221
L227:
	;
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v858)+1)))
	if v881&int32(15) != int32(2) {
		v887 = v874
		v888 = v873
		goto L221
	} else {
		goto L230
	}
L228:
	;
	v878 = int32(*(*int8)(unsafe.Add(mBase, uint32(v858)+1)))
	if v878 < int32(-64) {
		goto L226
	} else {
		goto L229
	}
L229:
	;
	v887 = v874
	v888 = v873
	goto L221
L230:
	;
	goto L226
L231:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v608))) = uint16(v856)
	v891 = F___memcpy(m, v888, v858, v887)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v612))) = v608
	goto L220
L232:
	;
	v901 = v619 + int32(4)
	if base.Ui32(int32(512)) <= base.Ui32(v897) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v614
	goto L157
L234:
	;
	if v897 != 0 {
		goto L237
	} else {
		goto L238
	}
L235:
	;
	goto L236
L236:
	;
	v908 = v614 + v897
	if (v614^v901)&int32(3) == int32(0) {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	base.MemoryCopy(m, v614, v901, v897)
	goto L239
L238:
	;
	goto L239
L239:
	;
	goto L233
L240:
	;
	if base.Ui32(v1040) < base.Ui32(v908) {
		goto L274
	} else {
		goto L275
	}
L241:
	;
	if v614&int32(3) == int32(0) {
		goto L245
	} else {
		goto L246
	}
L242:
	;
	goto L243
L243:
	;
	if base.Ui32(v908) < base.Ui32(int32(4)) {
		goto L265
	} else {
		goto L266
	}
L244:
	;
	v944 = v908 & int32(-4)
	if base.Ui32(v908) < base.Ui32(int32(64)) {
		v994 = v938
		v995 = v939
		goto L255
	} else {
		goto L256
	}
L245:
	;
	v938 = v901
	v939 = v614
	goto L244
L246:
	;
	goto L247
L247:
	;
	if v897 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v938 = v901
	v939 = v614
	goto L244
L249:
	;
	goto L250
L250:
	;
	v921 = v901
	v922 = v614
	goto L251
L251:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921))))
	*(*uint8)(unsafe.Add(mBase, uint32(v922))) = uint8(v926)
	v928 = int32(1)
	v929 = v921 + v928
	v931 = v922 + v928
	if v931&int32(3) == int32(0) {
		v938 = v929
		v939 = v931
		goto L244
	} else {
		goto L253
	}
L252:
	;
	v938 = v929
	v939 = v931
	goto L244
L253:
	;
	if base.Ui32(v931) < base.Ui32(v908) {
		v921 = v929
		v922 = v931
		goto L251
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	if base.Ui32(v944) <= base.Ui32(v995) {
		v1039 = v994
		v1040 = v995
		goto L240
	} else {
		goto L261
	}
L256:
	;
	v948 = v944 + int32(-64)
	if base.Ui32(v948) < base.Ui32(v939) {
		v994 = v938
		v995 = v939
		goto L255
	} else {
		goto L257
	}
L257:
	;
	v951 = v938
	v952 = v939
	goto L258
L258:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v951)))
	*(*int32)(unsafe.Add(mBase, uint32(v952))) = v956
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v951)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+4)) = v958
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v951)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+8)) = v960
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v951)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+12)) = v962
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v951)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+16)) = v964
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v951)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+20)) = v966
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v951)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+24)) = v968
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v951)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+28)) = v970
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v951)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+32)) = v972
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v951)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+36)) = v974
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v951)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+40)) = v976
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v951)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+44)) = v978
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v951)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+48)) = v980
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v951)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+52)) = v982
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v951)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+56)) = v984
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v951)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v952)+60)) = v986
	v988 = int32(-64)
	v989 = v951 - v988
	v991 = v952 - v988
	if base.Ui32(v991) <= base.Ui32(v948) {
		v951 = v989
		v952 = v991
		goto L258
	} else {
		goto L260
	}
L259:
	;
	v994 = v989
	v995 = v991
	goto L255
L260:
	;
	goto L259
L261:
	;
	v1001 = v994
	v1002 = v995
	goto L262
L262:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1001)))
	*(*int32)(unsafe.Add(mBase, uint32(v1002))) = v1006
	v1008 = int32(4)
	v1009 = v1001 + v1008
	v1011 = v1002 + v1008
	if base.Ui32(v1011) < base.Ui32(v944) {
		v1001 = v1009
		v1002 = v1011
		goto L262
	} else {
		goto L264
	}
L263:
	;
	v1039 = v1009
	v1040 = v1011
	goto L240
L264:
	;
	goto L263
L265:
	;
	v1039 = v901
	v1040 = v614
	goto L240
L266:
	;
	goto L267
L267:
	;
	if base.Ui32(v897) < base.Ui32(int32(4)) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1039 = v901
	v1040 = v614
	goto L240
L269:
	;
	goto L270
L270:
	;
	v1020 = v901
	v1021 = v614
	goto L271
L271:
	;
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1021))) = uint8(v1025)
	v1027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1021)+1)) = uint8(v1027)
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1021)+2)) = uint8(v1029)
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1021)+3)) = uint8(v1031)
	v1033 = int32(4)
	v1034 = v1020 + v1033
	v1036 = v1021 + v1033
	if base.Ui32(v1036) <= base.Ui32(v908-int32(4)) {
		v1020 = v1034
		v1021 = v1036
		goto L271
	} else {
		goto L273
	}
L272:
	;
	v1039 = v1034
	v1040 = v1036
	goto L240
L273:
	;
	goto L272
L274:
	;
	v1046 = v1039
	v1047 = v1040
	goto L277
L275:
	;
	goto L276
L276:
	;
	goto L233
L277:
	;
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1046))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1047))) = uint8(v1051)
	v1053 = int32(1)
	v1056 = v1047 + v1053
	if v1056 != v908 {
		v1046 = v1046 + v1053
		v1047 = v1056
		goto L277
	} else {
		goto L279
	}
L278:
	;
	goto L276
L279:
	;
	goto L278
L280:
	;
	goto L156
L281:
	;
	v1103 = int32(16)
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v1107 = v113 + int32(68)
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	v1109 = m.G0
	v1111 = v1109 - v1103
	m.G0 = v1111
	v1113 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1111)+8)) = v1113
	*(*int64)(unsafe.Add(mBase, uint32(v1111))) = v1113
	v1117 = int32(255)
	v1118 = int32(128)
	if base.Ui32(v1118) <= base.Ui32(v1108) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1121 = v1118
	goto L284
L283:
	;
	v1121 = v1108
	goto L284
L284:
	;
	v1123 = int32(base.Ui32(v1121) >> (uint(int32(3)) % 32))
	if v1123 == int32(0) {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	if base.Ui32(v1108) <= base.Ui32(int32(127)) {
		goto L297
	} else {
		goto L298
	}
L286:
	;
	goto L285
L287:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1111))) = uint8(v1117)
	v1130 = v1111 + v1123
	*(*uint8)(unsafe.Add(mBase, uint32(v1130-int32(1)))) = uint8(v1117)
	if base.Ui32(v1123) < base.Ui32(int32(3)) {
		goto L286
	} else {
		goto L288
	}
L288:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1111)+2)) = uint8(v1117)
	*(*uint8)(unsafe.Add(mBase, uint32(v1111)+1)) = uint8(v1117)
	*(*uint8)(unsafe.Add(mBase, uint32(v1130-int32(3)))) = uint8(v1117)
	*(*uint8)(unsafe.Add(mBase, uint32(v1130-int32(2)))) = uint8(v1117)
	if base.Ui32(v1123) < base.Ui32(int32(7)) {
		goto L286
	} else {
		goto L289
	}
L289:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1111)+3)) = uint8(v1117)
	*(*uint8)(unsafe.Add(mBase, uint32(v1130-int32(4)))) = uint8(v1117)
	if base.Ui32(v1123) < base.Ui32(int32(9)) {
		goto L286
	} else {
		goto L290
	}
L290:
	;
	v1155 = (int32(0) - v1111) & int32(3)
	v1156 = v1111 + v1155
	v1160 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1156))) = v1160
	v1164 = (v1123 - v1155) & int32(-4)
	v1165 = v1156 + v1164
	*(*int32)(unsafe.Add(mBase, uint32(v1165-int32(4)))) = v1160
	if base.Ui32(v1164) < base.Ui32(int32(9)) {
		goto L286
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1156)+8)) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(v1156)+4)) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(v1165-int32(8)))) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(v1165-int32(12)))) = v1160
	if base.Ui32(v1164) < base.Ui32(int32(25)) {
		goto L286
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1156)+24)) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(v1156)+20)) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(v1156)+16)) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(v1156)+12)) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(v1165-int32(16)))) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(v1165-int32(20)))) = v1160
	v1191 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v1165-v1191))) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(v1165-int32(28)))) = v1160
	v1200 = v1156&int32(4) | v1191
	v1201 = v1164 - v1200
	if base.Ui32(v1201) < base.Ui32(int32(32)) {
		goto L286
	} else {
		goto L293
	}
L293:
	;
	v1206 = base.I64_extend_i32_u(v1160) * int64(4294967297)
	v1209 = v1200 + v1156
	v1210 = v1201
	goto L294
L294:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1209)+24)) = v1206
	*(*int64)(unsafe.Add(mBase, uint32(v1209)+16)) = v1206
	*(*int64)(unsafe.Add(mBase, uint32(v1209)+8)) = v1206
	*(*int64)(unsafe.Add(mBase, uint32(v1209))) = v1206
	v1218 = int32(32)
	v1221 = v1210 - v1218
	if base.Ui32(int32(31)) < base.Ui32(v1221) {
		v1209 = v1209 + v1218
		v1210 = v1221
		goto L294
	} else {
		goto L296
	}
L295:
	;
	goto L286
L296:
	;
	goto L295
L297:
	;
	v1238 = int32(255) << (uint(int32(8)-v1121&int32(7)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1111+v1123))) = uint8(v1238)
	goto L299
L298:
	;
	goto L299
L299:
	;
	if v1105 != int32(10) {
		goto L303
	} else {
		goto L304
	}
L300:
	;
	m.G0 = v1111 + int32(16)
	goto L33
L301:
	;
	goto L300
L302:
	;
	if base.Ui32(int32(16)) < base.Ui32(v1266) {
		goto L301
	} else {
		goto L312
	}
L303:
	;
	if v1105 != int32(2) {
		goto L301
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v1252 = v113 + int32(76)
	v1253 = int32(16)
	v1254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111))))
	switch v1254 - int32(254) {
	case 0:
		goto L309
	case 1:
		goto L308
	default:
		v1266 = v1253
		v1267 = v1252
		goto L302
	}
L306:
	;
	v1266 = int32(4)
	v1267 = v113 + int32(72)
	goto L302
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1107)+24)) = int32(0)
	v1266 = v1253
	v1267 = v1252
	goto L302
L308:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111)+1)))
	if v1260&int32(15) != int32(2) {
		v1266 = v1253
		v1267 = v1252
		goto L302
	} else {
		goto L311
	}
L309:
	;
	v1257 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1111)+1)))
	if v1257 < int32(-64) {
		goto L307
	} else {
		goto L310
	}
L310:
	;
	v1266 = v1253
	v1267 = v1252
	goto L302
L311:
	;
	goto L307
L312:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1107))) = uint16(v1105)
	v1270 = F___memcpy(m, v1267, v1111, v1266)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v113+v1103))) = v1107
	goto L301
L313:
	;
	goto L1
L314:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v1291 == int32(0) {
		goto L317
	} else {
		goto L318
	}
L315:
	;
	goto L316
L316:
	;
	F_emscripten_builtin_free(m, v113)
	mBase = m.M
	goto L313
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v113
	goto L319
L318:
	;
	goto L319
L319:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1295 != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1295))) = v113
	goto L322
L321:
	;
	goto L322
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v113
	goto L313
}
func F_networkjoinsel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 float64
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 float32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 float64
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v130 float64
	_ = v130
	var v138 int32
	_ = v138
	var v139 float32
	_ = v139
	var v142 float32
	_ = v142
	var v145 float32
	_ = v145
	var v148 float32
	_ = v148
	var v150 float64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v177 float64
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v203 float64
	_ = v203
	var v209 int32
	_ = v209
	var v213 float32
	_ = v213
	var v215 float64
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v226 float64
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 float64
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int64
	_ = v256
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v294 float64
	_ = v294
	var v297 float64
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 float32
	_ = v304
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 float64
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v352 int32
	_ = v352
	var v361 float64
	_ = v361
	var v367 int32
	_ = v367
	var v368 float32
	_ = v368
	var v371 float32
	_ = v371
	var v374 float32
	_ = v374
	var v377 float32
	_ = v377
	var v379 float64
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v409 float64
	_ = v409
	var v413 int32
	_ = v413
	var v425 int32
	_ = v425
	var v435 float64
	_ = v435
	var v442 float32
	_ = v442
	var v444 float64
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v472 float64
	_ = v472
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v512 float64
	_ = v512
	var v528 int32
	_ = v528
	var v547 float64
	_ = v547
	var v557 int64
	_ = v557
	var v561 int64
	_ = v561
	var v562 int64
	_ = v562
	var v563 int32
	_ = v563
	var v566 float32
	_ = v566
	var v570 float32
	_ = v570
	var v574 float64
	_ = v574
	var v576 int32
	_ = v576
	var v597 float64
	_ = v597
	var v605 int32
	_ = v605
	var v631 float64
	_ = v631
	var v635 float64
	_ = v635
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v669 float64
	_ = v669
	var v677 float32
	_ = v677
	var v682 int64
	_ = v682
	var v683 float64
	_ = v683
	var v684 int32
	_ = v684
	var v686 float64
	_ = v686
	var v688 int32
	_ = v688
	var v711 float64
	_ = v711
	var v740 float64
	_ = v740
	var v743 float64
	_ = v743
	var v750 float64
	_ = v750
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v779 float64
	_ = v779
	var v787 float32
	_ = v787
	var v792 int64
	_ = v792
	var v793 float64
	_ = v793
	var v794 int32
	_ = v794
	var v796 float64
	_ = v796
	var v798 int32
	_ = v798
	var v821 float64
	_ = v821
	var v850 float64
	_ = v850
	var v860 float64
	_ = v860
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v904 float64
	_ = v904
	var v911 int32
	_ = v911
	var v915 int64
	_ = v915
	var v916 float64
	_ = v916
	var v917 int32
	_ = v917
	var v918 float64
	_ = v918
	var v919 int32
	_ = v919
	var v950 float64
	_ = v950
	var v954 int32
	_ = v954
	var v956 int64
	_ = v956
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v993 float64
	_ = v993
	var v998 float64
	_ = v998
	var v1004 float64
	_ = v1004
	var v1005 float64
	_ = v1005
	var v1012 int32
	_ = v1012
	var v1018 float64
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1054 float64
	_ = v1054
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1069 float64
	_ = v1069
	var v1077 float64
	_ = v1077
	v2 = int32(0)
	v20 = float64(0)
	v27 = m.G0
	v29 = v27 - int32(272)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	switch v35 - int32(931) {
	case 0:
		goto L5
	case 1:
		goto L6
	case 2:
		v59 = int32(-2)
		goto L1
	case 3:
		goto L3
	default:
		goto L4
	}
L1:
	;
	F_get_join_variables(m, v33, v32, v31, v29+int32(56), v29+int32(24), v29+int32(23))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L12
	}
L2:
	;
	v59 = int32(0)
	goto L1
L3:
	;
	v59 = int32(-1)
	goto L1
L4:
	;
	if v35 == int32(3552) {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	v59 = int32(2)
	goto L1
L6:
	;
	v59 = int32(1)
	goto L1
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int64(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v35
	F_errmsg_internal(m, int32(_a_F_networkjoinsel_0), v29)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_networkjoinsel_1), int32(870), int32(_a_F_networkjoinsel_2))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	switch v68 {
	case 0, 1, 2:
		goto L19
	default:
		goto L17
	case 4, 5:
		goto L18
	}
L13:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	if v1061 != 0 {
		goto L123
	} else {
		goto L124
	}
L14:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v300 != 0 {
		goto L47
	} else {
		goto L48
	}
L15:
	;
	v254 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v254
	v256 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+232)) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v29)+224)) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v29)+216)) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v29)+208)) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v29)+128)) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v29)+136)) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v29)+144)) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v29)+152)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v29)+160)) = v254
	v278 = v2
	v282 = v2
	v288 = v2
	v294 = v20
	v297 = v20
	goto L14
L16:
	;
	v244 = F_get_commutator(m, v35)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L8
	} else {
		goto L44
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L41
	}
L18:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+23)))
	if v221 != 0 {
		goto L16
	} else {
		goto L39
	}
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	if v69 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v75 = *(*float32)(unsafe.Add(mBase, uint32(v72+v73)+8))
	v81 = F_get_attstatsslot(m, v29+int32(208), v69, int32(1), int32(0), int32(3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v89 = F_get_attstatsslot(m, v29+int32(128), v85, int32(2), int32(0), int32(1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v91 = int32(1024)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v29)+224))
	if v91 <= v92 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v95 = v91
	goto L25
L24:
	;
	v95 = v92
	goto L25
L25:
	;
	v96 = base.F64_promote_f32(v75)
	if v81 == int32(0) {
		v278 = v95
		v282 = v2
		v288 = v89
		v294 = v20
		v297 = v96
		goto L14
	} else {
		goto L26
	}
L26:
	;
	if v92 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v278 = v95
	v282 = int32(1)
	v288 = v89
	v294 = v20
	v297 = v96
	goto L14
L28:
	;
	goto L29
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v29)+228))
	v104 = v95 & int32(3)
	v105 = int32(0)
	if v92 < int32(4) {
		v157 = v105
		v177 = v20
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v183 = v157
	v188 = v2
	v203 = v177
	goto L36
L31:
	;
	v110 = v105
	v119 = v2
	v130 = v20
	goto L32
L32:
	;
	v138 = v102 + v110<<(uint(int32(2))%32)
	v139 = *(*float32)(unsafe.Add(mBase, uint32(v138)))
	v142 = *(*float32)(unsafe.Add(mBase, uint32(v138)+4))
	v145 = *(*float32)(unsafe.Add(mBase, uint32(v138)+8))
	v148 = *(*float32)(unsafe.Add(mBase, uint32(v138)+12))
	v150 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v130, base.F64_promote_f32(v139)), base.F64_promote_f32(v142)), base.F64_promote_f32(v145)), base.F64_promote_f32(v148))
	v151 = int32(4)
	v152 = v110 + v151
	v154 = v119 + v151
	if v154 != v95&int32(2044) {
		v110 = v152
		v119 = v154
		v130 = v150
		goto L32
	} else {
		goto L34
	}
L33:
	;
	if v104 != 0 {
		v157 = v152
		v177 = v150
		goto L30
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v278 = v95
	v282 = int32(1)
	v288 = v89
	v294 = v150
	v297 = v96
	goto L14
L36:
	;
	v209 = int32(1)
	v213 = *(*float32)(unsafe.Add(mBase, uint32(v102+v183<<(uint(int32(2))%32))))
	v215 = base.F64_add(v203, base.F64_promote_f32(v213))
	v219 = v188 + v209
	if v219 != v104 {
		v183 = v183 + v209
		v188 = v219
		v203 = v215
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v278 = v95
	v282 = v209
	v288 = v89
	v294 = v215
	v297 = v96
	goto L14
L38:
	;
	goto L37
L39:
	;
	v226 = F_networkjoinsel_semi(m, v35, v59, v29+int32(56), v29+int32(24))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	v1054 = v226
	goto L13
L41:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v232
	F_errmsg_internal(m, int32(_a_F_networkjoinsel_3), v29+int32(16))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_networkjoinsel_1), int32(257), int32(_a_F_networkjoinsel_4))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v252 = F_networkjoinsel_semi(m, v244, int32(0)-v59, v29+int32(24), v29+int32(56))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v1054 = v252
	goto L13
L46:
	;
	if v35 == int32(3552) {
		goto L113
	} else {
		goto L114
	}
L47:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+16))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+22)))
	v304 = *(*float32)(unsafe.Add(mBase, uint32(v301+v302)+8))
	v310 = F_get_attstatsslot(m, v29+int32(168), v300, int32(1), int32(0), int32(3))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L8
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v954 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+200)) = v954
	v956 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+192)) = v956
	*(*int64)(unsafe.Add(mBase, uint32(v29)+184)) = v956
	*(*int64)(unsafe.Add(mBase, uint32(v29)+176)) = v956
	*(*int64)(unsafe.Add(mBase, uint32(v29)+168)) = v956
	*(*int64)(unsafe.Add(mBase, uint32(v29)+88)) = v956
	*(*int64)(unsafe.Add(mBase, uint32(v29)+96)) = v956
	*(*int64)(unsafe.Add(mBase, uint32(v29)+104)) = v956
	*(*int64)(unsafe.Add(mBase, uint32(v29)+112)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v954
	v985 = v2
	v989 = v2
	v993 = v20
	v998 = v20
	goto L46
L50:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v318 = F_get_attstatsslot(m, v29+int32(88), v314, int32(2), int32(0), int32(1))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v320 = int32(1024)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v29)+184))
	if v320 <= v321 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v324 = v320
	goto L54
L53:
	;
	v324 = v321
	goto L54
L54:
	;
	v325 = base.F64_promote_f32(v304)
	if v310 == int32(0) {
		v631 = v20
		v635 = float64(0)
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v288&v310 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L56:
	;
	if v282&v318 == int32(0) {
		v740 = v635
		v743 = v631
		goto L55
	} else {
		goto L87
	}
L57:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v29)+188))
	if v321 <= int32(0) {
		v472 = v20
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v282 == int32(0) {
		v740 = v20
		v743 = v472
		goto L55
	} else {
		goto L70
	}
L59:
	;
	v333 = v324 & int32(3)
	v334 = int32(0)
	if int32(4) <= v321 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v339 = v334
	v352 = v2
	v361 = v20
	goto L63
L61:
	;
	v387 = v334
	v409 = v20
	goto L62
L62:
	;
	v413 = v387
	v425 = v2
	v435 = v409
	goto L67
L63:
	;
	v367 = v329 + v339<<(uint(int32(2))%32)
	v368 = *(*float32)(unsafe.Add(mBase, uint32(v367)))
	v371 = *(*float32)(unsafe.Add(mBase, uint32(v367)+4))
	v374 = *(*float32)(unsafe.Add(mBase, uint32(v367)+8))
	v377 = *(*float32)(unsafe.Add(mBase, uint32(v367)+12))
	v379 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v361, base.F64_promote_f32(v368)), base.F64_promote_f32(v371)), base.F64_promote_f32(v374)), base.F64_promote_f32(v377))
	v380 = int32(4)
	v381 = v339 + v380
	v383 = v352 + v380
	if v383 != v324&int32(2044) {
		v339 = v381
		v352 = v383
		v361 = v379
		goto L63
	} else {
		goto L65
	}
L64:
	;
	if v333 == int32(0) {
		v472 = v379
		goto L58
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v387 = v381
	v409 = v379
	goto L62
L67:
	;
	v442 = *(*float32)(unsafe.Add(mBase, uint32(v329+v413<<(uint(int32(2))%32))))
	v444 = base.F64_add(v435, base.F64_promote_f32(v442))
	v445 = int32(1)
	v448 = v425 + v445
	if v448 != v333 {
		v413 = v413 + v445
		v425 = v448
		v435 = v444
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v472 = v444
	goto L58
L69:
	;
	goto L68
L70:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v29)+180))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v29)+228))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v29)+220))
	v481 = F_get_opcode(m, v35)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	F_fmgr_info(m, v481, v29+int32(244))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	if v278 <= int32(0) {
		v631 = v472
		v635 = float64(0)
		goto L56
	} else {
		goto L73
	}
L73:
	;
	v490 = int32(0)
	v496 = v490
	v512 = v20
	goto L74
L74:
	;
	if base.B2i32(v321 <= v490) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v631 = v472
	v635 = base.F64_add(v597, float64(0))
	goto L56
L76:
	;
	v528 = int32(0)
	v547 = v512
	goto L79
L77:
	;
	v597 = v512
	goto L78
L78:
	;
	v605 = v496 + int32(1)
	if v605 != v278 {
		v496 = v605
		v512 = v597
		goto L74
	} else {
		goto L86
	}
L79:
	;
	v557 = *(*int64)(unsafe.Add(mBase, uint32(v480+v496<<(uint(int32(3))%32))))
	v561 = *(*int64)(unsafe.Add(mBase, uint32(v478+v528<<(uint(int32(3))%32))))
	v562 = F_FunctionCall2Coll(m, v29+int32(244), int32(0), v557, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L8
	} else {
		goto L81
	}
L80:
	;
	v597 = v574
	goto L78
L81:
	;
	if v562 != int64(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v566 = *(*float32)(unsafe.Add(mBase, uint32(v479+v496<<(uint(int32(2))%32))))
	v570 = *(*float32)(unsafe.Add(mBase, uint32(v329+v528<<(uint(int32(2))%32))))
	v574 = base.F64_add(v547, base.F64_promote_f32(base.F32_mul(v566, v570)))
	goto L84
L83:
	;
	v574 = v547
	goto L84
L84:
	;
	v576 = v528 + int32(1)
	if v576 != v324 {
		v528 = v576
		v547 = v574
		goto L79
	} else {
		goto L85
	}
L85:
	;
	goto L80
L86:
	;
	goto L75
L87:
	;
	if int32(0) < v278 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v641 = int32(0)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v29)+104))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v29)+100))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v29)+228))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v29)+220))
	v648 = v641
	v669 = v20
	goto L91
L89:
	;
	v711 = v20
	goto L90
L90:
	;
	v740 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), v325), v631), v711), v635)
	v743 = v631
	goto L55
L91:
	;
	v677 = *(*float32)(unsafe.Add(mBase, uint32(v646+v648<<(uint(int32(2))%32))))
	v682 = *(*int64)(unsafe.Add(mBase, uint32(v647+v648<<(uint(int32(3))%32))))
	v683 = F_inet_hist_value_sel(m, v645, v644, v682, v641-v59)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L8
	} else {
		goto L93
	}
L92:
	;
	v711 = v686
	goto L90
L93:
	;
	v686 = base.F64_add(base.F64_mul(base.F64_promote_f32(v677), v683), v669)
	v688 = v648 + int32(1)
	if v688 != v278 {
		v648 = v688
		v669 = v686
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v750 = float64(0)
	if int32(0) < v321 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v850 = v740
	goto L97
L97:
	;
	if v318&v288 != int32(1) {
		v985 = v318
		v989 = v310
		v993 = v850
		v998 = v325
		goto L46
	} else {
		goto L105
	}
L98:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v29)+144))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v29)+140))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v29)+188))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v29)+180))
	v758 = int32(0)
	v779 = v750
	goto L101
L99:
	;
	v821 = v750
	goto L100
L100:
	;
	v850 = base.F64_add(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), v297), v294), v821), v740)
	goto L97
L101:
	;
	v787 = *(*float32)(unsafe.Add(mBase, uint32(v756+v758<<(uint(int32(2))%32))))
	v792 = *(*int64)(unsafe.Add(mBase, uint32(v757+v758<<(uint(int32(3))%32))))
	v793 = F_inet_hist_value_sel(m, v755, v754, v792, v59)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L8
	} else {
		goto L103
	}
L102:
	;
	v821 = v796
	goto L100
L103:
	;
	v796 = base.F64_add(base.F64_mul(base.F64_promote_f32(v787), v793), v779)
	v798 = v758 + int32(1)
	if v798 != v324 {
		v758 = v798
		v779 = v796
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v860 = float64(1)
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v29)+104))
	if int32(3) <= v868 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v29)+100))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v29)+144))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v29)+140))
	v874 = int32(1)
	v884 = v874
	v888 = int32(0)
	v904 = float64(0)
	goto L109
L107:
	;
	v950 = float64(0)
	goto L108
L108:
	;
	v985 = int32(1)
	v989 = v310
	v993 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_sub(v860, v297), v294), base.F64_sub(base.F64_sub(v860, v325), v743)), v950), v850)
	v998 = v325
	goto L46
L109:
	;
	v911 = v888 + int32(1)
	v915 = *(*int64)(unsafe.Add(mBase, uint32(v871+v884<<(uint(int32(3))%32))))
	v916 = F_inet_hist_value_sel(m, v873, v872, v915, v59)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L8
	} else {
		goto L111
	}
L110:
	;
	v950 = base.F64_div(v918, base.F64_convert_i32_s(v911))
	goto L108
L111:
	;
	v918 = base.F64_add(v904, v916)
	v919 = v884 + (int32(base.Ui32(v868-int32(3))>>(uint(int32(10))%32)) + v874)
	if v919 < v868-v874 {
		v884 = v919
		v888 = v911
		v904 = v918
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v1004 = float64(0.01)
	goto L115
L114:
	;
	v1004 = float64(0.005)
	goto L115
L115:
	;
	v1005 = float64(1)
	v1012 = int32(1)
	if (v282|v288)&v1012&((v985|v989)&v1012) != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v1018 = v993
	goto L118
L117:
	;
	v1018 = base.F64_mul(v1004, base.F64_mul(base.F64_sub(v1005, v297), base.F64_sub(v1005, v998)))
	goto L118
L118:
	;
	F_free_attstatsslot(m, v29+int32(208))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L8
	} else {
		goto L119
	}
L119:
	;
	F_free_attstatsslot(m, v29+int32(168))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L8
	} else {
		goto L120
	}
L120:
	;
	F_free_attstatsslot(m, v29+int32(128))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L8
	} else {
		goto L121
	}
L121:
	;
	F_free_attstatsslot(m, v29+int32(88))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L8
	} else {
		goto L122
	}
L122:
	;
	v1054 = v1018
	goto L13
L123:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	m.T0[v1062].(func(*base.Module, int32))(m, v1061)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L8
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v1065 != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L125
L127:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	m.T0[v1066].(func(*base.Module, int32))(m, v1065)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L8
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v1069 = float64(0)
	if base.F64_lt(v1054, v1069) != 0 {
		v1077 = v1069
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L129
L131:
	;
	m.G0 = v29 + int32(272)
	return base.I64_reinterpret_f64(v1077)
L132:
	;
	if base.F64_gt(v1054, float64(1)) == int32(0) {
		v1077 = v1054
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v1077 = float64(1)
	goto L131
}
func F_newLexeme(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v5 = l4
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 < v7 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v29 = v10
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v31 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v30 + v31
		v36 = v29 + v30<<(uint(int32(3))%32)
		v37 = l2 - l1
		v40 = F_palloc(m, v37+v31)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v36))) = v40
			if v37 != 0 {
				base.MemoryCopy(m, v40, l1, v37)
			} else {
			}
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
			v46 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v44+v37))) = uint8(v46)
			v49 = F_palloc(m, int32(16))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v49
				*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(0)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v54))) = l3
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)) = uint16(v5)
				return
			}
		}
	} else {
		if v7 == int32(0) {
			v13 = int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v13
			v17 = F_palloc_mul(m, int32(8), v13)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v27 = v17
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27
				v29 = v27
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v31 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v30 + v31
				v36 = v29 + v30<<(uint(int32(3))%32)
				v37 = l2 - l1
				v40 = F_palloc(m, v37+v31)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v36))) = v40
					if v37 != 0 {
						base.MemoryCopy(m, v40, l1, v37)
					} else {
					}
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					v46 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v44+v37))) = uint8(v46)
					v49 = F_palloc(m, int32(16))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(0)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v54))) = l3
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
						*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)) = uint16(v5)
						return
					}
				}
			}
		} else {
			v20 = v7 << (uint(int32(1)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v20
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v24 = F_repalloc_mul(m, v22, int32(8), v20)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v27 = v24
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27
				v29 = v27
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v31 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v30 + v31
				v36 = v29 + v30<<(uint(int32(3))%32)
				v37 = l2 - l1
				v40 = F_palloc(m, v37+v31)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v36))) = v40
					if v37 != 0 {
						base.MemoryCopy(m, v40, l1, v37)
					} else {
					}
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					v46 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v44+v37))) = uint8(v46)
					v49 = F_palloc(m, int32(16))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v49
						*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(0)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v54))) = l3
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
						*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)) = uint16(v5)
						return
					}
				}
			}
		}
	}
}
func F_newhicolorrow(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v8 < v9 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
		v35 = v11
		v37 = v8
		v38 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v37 + v38
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v45 = v35 + v41*v8<<(uint(v38)%32)
		v47 = v41 << (uint(v38) % 32)
		if v47 != 0 {
			base.MemoryCopy(m, v45, v35+l1*v41<<(uint(int32(1))%32), v47)
		} else {
		}
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		if v53 <= int32(0) {
			v93 = v8
		} else {
			v59 = int32(0)
			for {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v65 = int32(1)
				v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45+v59<<(uint(v65)%32)))))
				v71 = v64 + v68*int32(24)
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v72 + v65
				v77 = v59 + v65
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v77 < v78 {
					v59 = v77
					continue
				} else {
					break
				}
				break
			}
			v93 = v8
		}
		return v93
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v16 = base.I32_div_s(int32(2147483647), v13<<(uint(int32(1))%32))
		if v16 <= v9 {
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = int32(101)
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
			if v85 != 0 {
				v87 = v85
			} else {
				v87 = int32(12)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = v87
			v93 = int32(0)
			return v93
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
			v23 = F_repalloc_mul_extended(m, v18, int32(2), v9*v13<<(uint(int32(1))%32))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = int32(101)
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
					if v85 != 0 {
						v87 = v85
					} else {
						v87 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = v87
					v93 = int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v23
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v30 << (uint(int32(1)) % 32)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					v35 = v23
					v37 = v34
					v38 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v37 + v38
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v45 = v35 + v41*v8<<(uint(v38)%32)
					v47 = v41 << (uint(v38) % 32)
					if v47 != 0 {
						base.MemoryCopy(m, v45, v35+l1*v41<<(uint(int32(1))%32), v47)
					} else {
					}
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					if v53 <= int32(0) {
						v93 = v8
					} else {
						v59 = int32(0)
						for {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v65 = int32(1)
							v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45+v59<<(uint(v65)%32)))))
							v71 = v64 + v68*int32(24)
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v72 + v65
							v77 = v59 + v65
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
							if v77 < v78 {
								v59 = v77
								continue
							} else {
								break
							}
							break
						}
						v93 = v8
					}
				}
				return v93
			}
		}
	}
}
func F_nextval_internal(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v148 int32
	_ = v148
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v176 int32
	_ = v176
	var v181 int64
	_ = v181
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int64
	_ = v215
	var v226 int64
	_ = v226
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v235 int64
	_ = v235
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v263 int64
	_ = v263
	var v266 int64
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int64
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v371 int64
	_ = v371
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	v11 = int64(0)
	v25 = m.G0
	v27 = v25 - int32(112)
	m.G0 = v27
	F_init_sequence(m, l0, v27+int32(108), v27+int32(104))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L109
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L106
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L102
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[0]))
	v42 = F_pg_class_aclcheck(m, v38, v40, int64(260))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+24)))
	if v45 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v42 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	F_PreventCommandIfReadOnly(m, int32(_a_F_nextval_internal_0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_PreventCommandIfParallelMode(m, int32(_a_F_nextval_internal_0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v54)+16))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v54)+24))
	if v55 != v56 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	m.G0 = v27 + int32(112)
	return v371
L17:
	;
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v54)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v58 + v55
	F_relation_close(m, v44, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v69 = F_SearchSysCache1(m, int32(61), base.I64_extend_i32_u(l0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[1])) = v54
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v54)+16))
	v371 = v66
	goto L16
L21:
	;
	if v69 == int32(0) {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+22)))
	v75 = v73 + v74
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+48)))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v75)+40))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v75)+32))
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v75)+24))
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v75)+16))
	F_ReleaseCatCache(m, v69)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v87 = F_read_seq_tuple(m, v44, v27+int32(100), v27+int32(80))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v27)+100))
	if v89 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+16)))
	v115 = base.I64_extend_i32_u(v110^int32(-1)) & int64(1)
	v116 = v77 - v115
	if base.B2i32(v116 <= v109)&(v110&int32(1)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[2]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93+(v89^int32(-1))<<(uint(int32(2))%32))))
	v107 = v99
	goto L25
L27:
	;
	goto L28
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[3]))
	v107 = v101 + v89<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L29:
	;
	if v139 == int64(0) {
		goto L41
	} else {
		goto L42
	}
L30:
	;
	v125 = v116 + int64(32)
	v136 = int32(1)
	v139 = v125
	v140 = v125
	goto L29
L31:
	;
	goto L32
L32:
	;
	v126 = F_GetRedoRecPtr(m)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v128 = int64(32)
	v129 = v77 + v128
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
	v133 = base.B2i32(base.Ui64(base.I64_rotl(v130, v128)) <= base.Ui64(v126))
	if base.Ui64(base.I64_rotl(v130, v128)) <= base.Ui64(v126) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v134 = v129
	goto L36
L35:
	;
	v134 = v109
	goto L36
L36:
	;
	if base.Ui64(base.I64_rotl(v130, v128)) <= base.Ui64(v126) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v135 = v129
	goto L39
L38:
	;
	v135 = v77
	goto L39
L39:
	;
	v136 = v133
	v139 = v135
	v140 = v134
	goto L29
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v257
	*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v255
	*(*int64)(unsafe.Add(mBase, uint32(v54)+32)) = v80
	v270 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+12)) = uint8(v270)
	*(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[1])) = v54
	if v136 != 0 {
		goto L76
	} else {
		goto L77
	}
L41:
	;
	v253 = v108
	v255 = v108
	v257 = v108
	v263 = v140
	v266 = v11
	goto L40
L42:
	;
	goto L43
L43:
	;
	v148 = v76 & int32(1)
	v160 = v108
	v161 = v108
	v163 = v108
	v164 = v115
	v167 = v139
	v169 = v140
	goto L44
L44:
	;
	if base.B2i32(v80 <= int64(0)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v253 = v226
	v255 = v235
	v257 = v236
	v263 = v238
	v266 = v11
	goto L40
L46:
	;
	if v164 < v77 {
		goto L68
	} else {
		goto L69
	}
L47:
	;
	if v148 == int32(0) {
		goto L3
	} else {
		goto L67
	}
L48:
	;
	v253 = v160
	v255 = v161
	v257 = v163
	v263 = v169
	v266 = v167
	goto L40
L49:
	;
	v176 = base.B2i32(int64(0) <= v79)
	if v176&base.B2i32(v79-v80 < v160) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v210 = base.B2i32(v78 < int64(0))
	if v210&base.B2i32(v160 < v78-v80) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v181 = v160 + v80
	if v176|base.B2i32(v181 <= v79) != 0 {
		v226 = v181
		goto L46
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if int64(0) < v164 {
		goto L48
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	if v148 != 0 {
		v226 = v78
		goto L46
	} else {
		goto L57
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(402653314))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+24)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v194 + int32(4)
	F_errmsg(m, int32(_a_F_nextval_internal_1), v27+int32(16))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_nextval_internal_2), int32(751), int32(_a_F_nextval_internal_3))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v215 = v160 + v80
	if base.B2i32(v78 <= v215)|v210 != 0 {
		v226 = v215
		goto L46
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v164 <= int64(0) {
		goto L47
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	goto L48
L67:
	;
	v226 = v79
	goto L46
L68:
	;
	if v164 == int64(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v235 = v161
	v236 = v163
	v237 = v164
	v238 = v169
	goto L70
L70:
	;
	v240 = v167 - int64(1)
	if v240 != int64(0) {
		v160 = v226
		v161 = v235
		v163 = v236
		v164 = v237
		v167 = v240
		v169 = v238
		goto L44
	} else {
		goto L74
	}
L71:
	;
	v230 = v226
	goto L73
L72:
	;
	v230 = v161
	goto L73
L73:
	;
	v231 = int64(1)
	v235 = v230
	v236 = v226
	v237 = v164 + v231
	v238 = v169 - v231
	goto L70
L74:
	;
	goto L45
L75:
	;
	v344 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+16)) = uint8(v344)
	*(*int64)(unsafe.Add(mBase, uint32(v87))) = v257
	*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = v263 - v266
	v348 = int32(_a_F_nextval_internal_4)
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[4])) = v350 - v344
	F_UnlockReleaseBuffer(m, v89)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L100
	}
L76:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+118)))
	if v276 != int32(112) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	v336 = int32(_a_F_nextval_internal_4)
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[4])) = v338 + int32(1)
	F_MarkBufferDirty(m, v89)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L99
	}
L79:
	;
	v287 = int32(_a_F_nextval_internal_4)
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[4])) = v289 + int32(1)
	F_MarkBufferDirty(m, v89)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L87
	}
L80:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[5]))
	if v280 <= int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v44)+32))
	if v283 != 0 {
		goto L79
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v285 = F_GetTopTransactionId(m)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v284 != 0 {
		goto L79
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	goto L79
L87:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+118)))
	if v296 != int32(112) {
		goto L75
	} else {
		goto L88
	}
L88:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_nextval_internal[5]))
	if v300 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v44)+32))
	if v303 != 0 {
		goto L75
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v304 != 0 {
		goto L75
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	F_XLogRegisterBuffer(m, int32(0), v89, int32(6))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v311 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+16)) = uint8(v311)
	*(*int64)(unsafe.Add(mBase, uint32(v87))) = v253
	*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = int64(0)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v316
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = v318
	F_XLogRegisterData(m, v27-int32(-64), int32(12))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v27)+96))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	F_XLogRegisterData(m, v325, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v331 = F_XLogInsert(m, int32(15), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v107))) = base.I64_rotl(v331, int64(32))
	goto L75
L99:
	;
	goto L75
L100:
	;
	F_relation_close(m, v44, int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v371 = v255
	goto L16
L102:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v395 + int32(4)
	F_errmsg(m, int32(_a_F_nextval_internal_5), v27+int32(48))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_nextval_internal_2), int32(656), int32(_a_F_nextval_internal_3))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
	F_errmsg_internal(m, int32(_a_F_nextval_internal_6), v27)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_nextval_internal_2), int32(681), int32(_a_F_nextval_internal_3))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode(m, int32(402653314))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v429 + int32(4)
	F_errmsg(m, int32(_a_F_nextval_internal_7), v27+int32(32))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_nextval_internal_2), int32(770), int32(_a_F_nextval_internal_3))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_nlevel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v5)+4)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v10 != v5 {
			F_pfree(m, v5)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				return v9
			}
		} else {
			return v9
		}
	}
}
func F_nocachegetattr(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v459 int64
	_ = v459
	var v460 int64
	_ = v460
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v479 int64
	_ = v479
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v19 = v17 + int32(23)
	v20 = int32(1)
	v21 = l1 - v20
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+20)))
	if v22&v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = int32(0)
	v27 = v21 >> (uint(int32(3)) % 32)
	if v27 <= v25 {
		v50 = v25
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v75 = v21
	goto L3
L3:
	;
	v83 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v85 <= v83 {
		v99 = v83
		v100 = v83
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v19))))
	v68 = base.I32_ctz(v62^int32(-1)) + v50<<(uint(int32(3))%32)
	if v68 < v21 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v31 = v25
	goto L6
L6:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v19))))
	if v43 != int32(255) {
		v50 = v31
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v50 = v27
	goto L4
L8:
	;
	v47 = v31 + int32(1)
	if v47 != v27 {
		v31 = v47
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v70 = v68
	goto L12
L11:
	;
	v70 = v21
	goto L12
L12:
	;
	v75 = v70
	goto L3
L13:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
	v102 = v17 + v101
	if v22&int32(2) != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v88 = int32(0)
	if v75 <= v88 {
		v99 = v83
		v100 = v88
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if base.Ui32(v85) < base.Ui32(v75) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v92 = v85
	goto L18
L17:
	;
	v92 = v75
	goto L18
L18:
	;
	v94 = v92 - int32(1)
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2+v94<<(uint(int32(3))%32))+28)))
	v99 = v94
	v100 = v98
	goto L13
L19:
	;
	v434 = l2 + v21<<(uint(int32(3))%32)
	v436 = v434 + int32(28)
	v437 = int32(*(*int16)(unsafe.Add(mBase, uint32(v434)+30)))
	if v437 == int32(-1) {
		goto L94
	} else {
		goto L95
	}
L20:
	;
	if v99 < v75 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	if v99 < v75 {
		goto L74
	} else {
		goto L75
	}
L23:
	;
	v108 = v100
	v109 = v99
	goto L26
L24:
	;
	v187 = v100
	v188 = v99
	goto L25
L25:
	;
	if v21 <= v188 {
		v420 = v187
		goto L19
	} else {
		goto L48
	}
L26:
	;
	v122 = l2 + int32(28) + v109<<(uint(int32(3))%32)
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v122)+2)))
	if v123 == int32(-1) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v187 = v183
	v188 = v75
	goto L25
L28:
	;
	v183 = v180 + v181
	v185 = v109 + int32(1)
	if v185 != v75 {
		v108 = v183
		v109 = v185
		goto L26
	} else {
		goto L47
	}
L29:
	;
	if v143&int32(1) != 0 {
		goto L44
	} else {
		goto L45
	}
L30:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v102))))
	if v127 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
	v164 = int32(0)
	v166 = (v108 + v160 - int32(1)) & (v164 - v160)
	if v164 < v123 {
		v180 = v166
		v181 = v123
		goto L28
	} else {
		goto L43
	}
L33:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
	v136 = (v108 + v130 - int32(1)) & (int32(0) - v130)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v136))))
	v139 = v136
	v140 = v138
	goto L35
L34:
	;
	v139 = v108
	v140 = v127
	goto L35
L35:
	;
	v141 = v139 + v102
	v143 = v140 & int32(255)
	if v143 != int32(1) {
		goto L29
	} else {
		goto L36
	}
L36:
	;
	v147 = int32(18)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)))
	if v149 == v147 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v152 = v147
	goto L39
L38:
	;
	v152 = int32(2)
	goto L39
L39:
	;
	if base.Ui32((v149-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v159 = int32(6)
	goto L42
L41:
	;
	v159 = v152
	goto L42
L42:
	;
	v180 = v139
	v181 = v159
	goto L28
L43:
	;
	v170 = F_strlen(m, v166+v102)
	mBase = m.M
	v180 = v166
	v181 = v170 + int32(1)
	goto L28
L44:
	;
	v180 = v139
	v181 = int32(base.Ui32(v143) >> (uint(int32(1)) % 32))
	goto L28
L45:
	;
	goto L46
L46:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v180 = v139
	v181 = int32(base.Ui32(v177) >> (uint(int32(2)) % 32))
	goto L28
L47:
	;
	goto L27
L48:
	;
	v202 = v187
	v203 = v188
	goto L49
L49:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(base.Ui32(v203)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v217)>>(uint(v203&int32(7))%32))&int32(1) != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v420 = v288
	goto L19
L51:
	;
	v225 = l2 + int32(28) + v203<<(uint(int32(3))%32)
	v226 = int32(*(*int16)(unsafe.Add(mBase, uint32(v225)+2)))
	if v226 == int32(-1) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v288 = v202
	goto L53
L53:
	;
	v292 = v203 + int32(1)
	if v292 != v21 {
		v202 = v288
		v203 = v292
		goto L49
	} else {
		goto L73
	}
L54:
	;
	v288 = v284 + v285
	goto L53
L55:
	;
	if v247&int32(1) != 0 {
		goto L70
	} else {
		goto L71
	}
L56:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202+v102))))
	if v230 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+5)))
	v268 = int32(0)
	v270 = (v202 + v264 - int32(1)) & (v268 - v264)
	if v268 < v226 {
		v284 = v270
		v285 = v226
		goto L54
	} else {
		goto L69
	}
L59:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+5)))
	v239 = (v202 + v233 - int32(1)) & (int32(0) - v233)
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v239))))
	v242 = v239
	v243 = v241
	goto L61
L60:
	;
	v242 = v202
	v243 = v230
	goto L61
L61:
	;
	v245 = v242 + v102
	v247 = v243 & int32(255)
	if v247 != int32(1) {
		goto L55
	} else {
		goto L62
	}
L62:
	;
	v251 = int32(18)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)))
	if v253 == v251 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v256 = v251
	goto L65
L64:
	;
	v256 = int32(2)
	goto L65
L65:
	;
	if base.Ui32((v253-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v263 = int32(6)
	goto L68
L67:
	;
	v263 = v256
	goto L68
L68:
	;
	v284 = v242
	v285 = v263
	goto L54
L69:
	;
	v274 = F_strlen(m, v270+v102)
	mBase = m.M
	v284 = v270
	v285 = v274 + int32(1)
	goto L54
L70:
	;
	v284 = v242
	v285 = int32(base.Ui32(v247) >> (uint(int32(1)) % 32))
	goto L54
L71:
	;
	goto L72
L72:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v284 = v242
	v285 = int32(base.Ui32(v281) >> (uint(int32(2)) % 32))
	goto L54
L73:
	;
	goto L50
L74:
	;
	v295 = int32(1)
	v296 = v99 + v295
	v298 = l2 + int32(28)
	if (v75-v99)&v295 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v366 = v100
	v367 = v99
	goto L76
L76:
	;
	if v21 <= v367 {
		v420 = v366
		goto L19
	} else {
		goto L86
	}
L77:
	;
	v304 = v298 + v99<<(uint(int32(3))%32)
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+5)))
	v312 = int32(*(*int16)(unsafe.Add(mBase, uint32(v304)+2)))
	v314 = (v100+v305-int32(1))&(int32(0)-v305) + v312
	v315 = v296
	goto L79
L78:
	;
	v314 = v100
	v315 = v99
	goto L79
L79:
	;
	if v296 != v75 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v318 = v314
	v319 = v315
	goto L83
L81:
	;
	v354 = v314
	goto L82
L82:
	;
	v366 = v354
	v367 = v75
	goto L76
L83:
	;
	v332 = v298 + v319<<(uint(int32(3))%32)
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+5)))
	v335 = int32(1)
	v337 = int32(0)
	v340 = int32(*(*int16)(unsafe.Add(mBase, uint32(v332)+2)))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+13)))
	v349 = int32(*(*int16)(unsafe.Add(mBase, uint32(v332)+10)))
	v350 = ((v318+v333-v335)&(v337-v333)+v340+v342-v335)&(v337-v342) + v349
	v352 = v319 + int32(2)
	if v352 != v75 {
		v318 = v350
		v319 = v352
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v354 = v350
	goto L82
L85:
	;
	goto L84
L86:
	;
	v381 = v366
	v382 = v367
	goto L87
L87:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(base.Ui32(v382)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v396)>>(uint(v382&int32(7))%32))&int32(1) != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v420 = v414
	goto L19
L89:
	;
	v404 = l2 + int32(28) + v382<<(uint(int32(3))%32)
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+5)))
	v412 = int32(*(*int16)(unsafe.Add(mBase, uint32(v404)+2)))
	v414 = (v381+v405-int32(1))&(int32(0)-v405) + v412
	goto L91
L90:
	;
	v414 = v381
	goto L91
L91:
	;
	v418 = v382 + int32(1)
	if v418 != v21 {
		v381 = v414
		v382 = v418
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	v451 = v449 + v102
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436)+4)))
	if v452 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L94:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420+v102))))
	if v441 != 0 {
		v449 = v420
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436)+5)))
	v449 = (v420 + v442 - int32(1)) & (int32(0) - v442)
	goto L93
L97:
	;
	goto L96
L98:
	;
	m.G0 = v15 + int32(16)
	return v479
L99:
	;
	if base.I32_popcnt(v437) != int32(1) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v479 = base.I64_extend_i32_u(v451)
	goto L98
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L108
	} else {
		goto L109
	}
L103:
	;
	switch base.I32_ctz(v437) {
	case 0:
		goto L107
	case 1:
		goto L106
	case 2:
		goto L105
	case 3:
		goto L104
	default:
		goto L102
	}
L104:
	;
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v451)))
	v479 = v462
	goto L98
L105:
	;
	v461 = int64(*(*int32)(unsafe.Add(mBase, uint32(v451))))
	v479 = v461
	goto L98
L106:
	;
	v460 = int64(*(*int16)(unsafe.Add(mBase, uint32(v451))))
	v479 = v460
	goto L98
L107:
	;
	v459 = int64(*(*int8)(unsafe.Add(mBase, uint32(v451))))
	v479 = v459
	goto L98
L108:
	;
	return int64(0)
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v437
	F_errmsg_internal(m, int32(_a_F_nocachegetattr_0), v15)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_nocachegetattr_1), int32(123), int32(_a_F_nocachegetattr_2))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L108
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_now(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_now[0]))
	return v3
}
func F_numerictypmodout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_palloc(m, int32(64))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		if int32(4) <= v8 {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(base.Ui32(v8-int32(4)) >> (uint(int32(16)) % 32))
			v21 = int32(21)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = (v8<<(uint(v21)%32) - int32(_a_F_numerictypmodout_0)) >> (uint(v21) % 32)
			v30 = F_pg_snprintf(m, v10, int32(64), int32(_a_F_numerictypmodout_1), v6)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int64(0)
			} else {
				m.G0 = v6 + int32(16)
				return base.I64_extend_i32_u(v10)
			}
		} else {
			v32 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v32)
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v10)
		}
	}
}
func F_numericvar_deserialize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v7 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_pfree(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = F_palloc(m, v7<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v16
	v19 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v16))) = uint16(v19)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v23 + int32(2)
	v28 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v28
	v32 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v32
	v36 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v36
	if int32(0) < v7 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = v19
	goto L14
L12:
	;
	goto L13
L13:
	;
	return
L14:
	;
	v47 = F_pq_getmsgint(m, l0, int32(2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v50 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v49+v43<<(uint(v50)%32)))) = uint16(v47)
	v55 = v43 + v50
	if v55 != v7 {
		v43 = v55
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
}
func F_numericvar_to_int64(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int64
	_ = v180
	var v184 int32
	_ = v184
	var v194 int64
	_ = v194
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v204 int64
	_ = v204
	var v207 int64
	_ = v207
	var v210 int64
	_ = v210
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v218 int64
	_ = v218
	var v225 int64
	_ = v225
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v245 int64
	_ = v245
	var v248 int64
	_ = v248
	var v251 int64
	_ = v251
	var v254 int32
	_ = v254
	var v266 int64
	_ = v266
	var v269 int32
	_ = v269
	var v277 int64
	_ = v277
	var v294 int32
	_ = v294
	var v311 int32
	_ = v311
	var v325 int32
	_ = v325
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = F_palloc(m, v17<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v22))) = uint16(v26)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v28 <= v26 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v42 = v40 << (uint(int32(2)) % 32)
	if v42+int32(4) < int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v32 = v28 << (uint(int32(1)) % 32)
	if v32 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v22+int32(2), v37, v32)
	goto L3
L6:
	;
	m.G0 = v15 + int32(16)
	return v325
L7:
	;
	F_pfree(m, v22)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L58
	}
L8:
	;
	v325 = int32(1)
	goto L6
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	F_pfree(m, v22)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L57
	}
L10:
	;
	v48 = v22 + int32(2)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v53 = base.I32_div_s(v42+int32(7), int32(4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v54 <= v53 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if int32(0) < v107 {
		goto L26
	} else {
		goto L27
	}
L12:
	;
	v107 = v54
	v108 = v48
	v110 = v40
	goto L11
L13:
	;
	goto L14
L14:
	;
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48+v53<<(uint(int32(1))%32)))))
	if int32(_a_F_numericvar_to_int64_0) <= v59 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v62 = v53
	goto L18
L16:
	;
	v87 = v53
	goto L17
L17:
	;
	if int32(0) <= v87 {
		v107 = v53
		v108 = v48
		v110 = v40
		goto L11
	} else {
		goto L24
	}
L18:
	;
	v74 = int32(1)
	v76 = v22 + v62<<(uint(v74)%32)
	v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76))))
	v81 = base.B2i32(int32(_a_F_numericvar_to_int64_1) < v79)
	if int32(_a_F_numericvar_to_int64_1) < v79 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v87 = v86
	goto L17
L20:
	;
	v82 = int32(-9999)
	goto L22
L21:
	;
	v82 = v74
	goto L22
L22:
	;
	v83 = v82 + v79
	*(*uint16)(unsafe.Add(mBase, uint32(v76))) = uint16(v83)
	v86 = v62 - int32(1)
	if int32(_a_F_numericvar_to_int64_1) < v79 {
		v62 = v86
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v101 = int32(1)
	v107 = v53 + v101
	v108 = v22
	v110 = v40 + v101
	goto L11
L25:
	;
	v180 = int64(0) - base.I64_extend16_s(base.I64_extend_i32_u(v165))
	if int32(0) < v170 {
		goto L40
	} else {
		goto L41
	}
L26:
	;
	v121 = v107
	v122 = v108
	v124 = v110
	goto L29
L27:
	;
	goto L28
L28:
	;
	if v107 == int32(0) {
		goto L9
	} else {
		goto L39
	}
L29:
	;
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	if v131 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v134 = v121
	goto L34
L32:
	;
	goto L33
L33:
	;
	v154 = int32(1)
	if v154 < v121 {
		v121 = v121 - v154
		v122 = v122 + int32(2)
		v124 = v124 - v154
		goto L29
	} else {
		goto L38
	}
L34:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122+v134<<(uint(int32(1))%32)-int32(2)))))
	if v149 != 0 {
		v165 = v131
		v167 = v134
		v168 = v122
		v170 = v124
		goto L25
	} else {
		goto L36
	}
L36:
	;
	v150 = int32(1)
	if v150 < v134 {
		v134 = v134 - v150
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L9
L38:
	;
	goto L9
L39:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108))))
	v165 = v164
	v167 = v107
	v168 = v108
	v170 = v110
	goto L25
L40:
	;
	v184 = int32(1)
	v194 = v180
	goto L43
L41:
	;
	v266 = v180
	goto L42
L42:
	;
	F_pfree(m, v22)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L52
	}
L43:
	;
	v198 = int64(10000)
	v199 = int64(0)
	v204 = int64(32)
	v207 = int64(base.Ui64(v194) >> (uint(v204) % 64))
	v210 = int64(4294967295)
	v213 = v194 & v210
	v214 = v198 * v213
	v218 = int64(base.Ui64(v214)>>(uint(v204)%64)) + v198*v207
	v225 = v213*v199 + v218&v210
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v194*v199 + v194>>(uint(int64(63))%64)*v198 + v199*v207 + int64(base.Ui64(v218)>>(uint(v204)%64)) + int64(base.Ui64(v225)>>(uint(v204)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v214&v210 | v225<<(uint(v204)%64)
	goto L45
L44:
	;
	v266 = v251
	goto L42
L45:
	;
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	if v236 != v237>>(uint(int64(63))%64) {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	if v184 < v167 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v245 = int64(*(*int16)(unsafe.Add(mBase, uint32(v168+v184<<(uint(int32(1))%32)))))
	v248 = v237 - v245
	if base.B2i32(int64(0) < v245) != base.B2i32(v248 < v237) {
		goto L7
	} else {
		goto L50
	}
L48:
	;
	v251 = v237
	goto L49
L49:
	;
	v254 = v184 + int32(1)
	if v254 <= v170 {
		v184 = v254
		v194 = v251
		goto L43
	} else {
		goto L51
	}
L50:
	;
	v251 = v248
	goto L49
L51:
	;
	goto L44
L52:
	;
	if v49 != int32(_a_F_numericvar_to_int64_2) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v266 == int64(-9223372036854775807-1) {
		v325 = int32(0)
		goto L6
	} else {
		goto L56
	}
L54:
	;
	v277 = v266
	goto L55
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v277
	goto L8
L56:
	;
	v277 = int64(0) - v266
	goto L55
L57:
	;
	goto L8
L58:
	;
	v325 = int32(0)
	goto L6
}
