package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsvector_delete_by_indices(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	v4 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(2) <= l2 {
		F_pg_qsort(m, l1, l2, int32(4), int32(1521))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v29 = int32(1)
			v31 = v4
			for {
				v41 = int32(2)
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1+v29<<(uint(v41)%32))))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l1+v31<<(uint(v41)%32))))
				if v44 == v48 {
					v57 = v31
				} else {
					v51 = v31 + int32(1)
					if v29 == v51 {
						v57 = v29
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1+v51<<(uint(int32(2))%32)))) = v44
						v57 = v51
					}
				}
				v60 = v29 + int32(1)
				if v60 != l2 {
					v29 = v60
					v31 = v57
					continue
				} else {
					break
				}
				break
			}
			v66 = v57 + int32(1)
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v82 = F_palloc0(m, int32(base.Ui32(v79)>>(uint(int32(2))%32)))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v85 = v84 - v66
				*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v85
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v87 <= int32(0) {
					v214 = v85
					v216 = int32(32)
				} else {
					v91 = int32(8)
					v92 = l0 + v91
					v93 = int32(2)
					v97 = v82 + v91
					v100 = v97 + v85<<(uint(v93)%32)
					v101 = int32(0)
					v107 = v101
					v108 = v87
					v109 = v101
					v112 = v101
					v113 = v4
					for {
						if v66 <= v113 {
							v129 = v92 + v109<<(uint(int32(2))%32)
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
							v134 = int32(base.Ui32(v130)>>(uint(int32(1))%32)) & int32(2047)
							if v134 != 0 {
								base.MemoryCopy(m, v107+v100, v92+v16<<(uint(v93)%32)+int32(base.Ui32(v130)>>(uint(int32(12))%32)), v134)
							} else {
							}
							v142 = v97 + v112<<(uint(int32(2))%32)
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
							v144 = int32(1)
							v145 = v143 & v144
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
							*(*int32)(unsafe.Add(mBase, uint32(v142))) = v145 | v146&int32(-2)
							v151 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
							*(*int32)(unsafe.Add(mBase, uint32(v142))) = v151&int32(4094) | v107<<(uint(int32(12))%32) | v145
							v159 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
							v163 = int32(base.Ui32(v159)>>(uint(v144)%32)) & int32(2047)
							v164 = v163 + v107
							if v159&v144 != 0 {
								v167 = int32(1)
								v170 = (v164 + v167) & int32(-2)
								v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v172 = int32(2)
								v182 = v92 + v171<<(uint(v172)%32) + (int32(base.Ui32(v159)>>(uint(int32(12))%32))+v163+v167)&int32(_a_F_tsvector_delete_by_indices_0)
								v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182))))
								v187 = v183<<(uint(v167)%32) + v172
								if v187 != 0 {
									base.MemoryCopy(m, v170+v100, v182, v187)
								} else {
								}
								v191 = v170 + v187
							} else {
								v191 = v164
							}
							v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v197 = v191
							v198 = v196
							v200 = v112 + int32(1)
							v201 = v113
						} else {
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l1+v113<<(uint(int32(2))%32))))
							if v109 != v123 {
								v129 = v92 + v109<<(uint(int32(2))%32)
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
								v134 = int32(base.Ui32(v130)>>(uint(int32(1))%32)) & int32(2047)
								if v134 != 0 {
									base.MemoryCopy(m, v107+v100, v92+v16<<(uint(v93)%32)+int32(base.Ui32(v130)>>(uint(int32(12))%32)), v134)
								} else {
								}
								v142 = v97 + v112<<(uint(int32(2))%32)
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
								v144 = int32(1)
								v145 = v143 & v144
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
								*(*int32)(unsafe.Add(mBase, uint32(v142))) = v145 | v146&int32(-2)
								v151 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
								*(*int32)(unsafe.Add(mBase, uint32(v142))) = v151&int32(4094) | v107<<(uint(int32(12))%32) | v145
								v159 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
								v163 = int32(base.Ui32(v159)>>(uint(v144)%32)) & int32(2047)
								v164 = v163 + v107
								if v159&v144 != 0 {
									v167 = int32(1)
									v170 = (v164 + v167) & int32(-2)
									v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v172 = int32(2)
									v182 = v92 + v171<<(uint(v172)%32) + (int32(base.Ui32(v159)>>(uint(int32(12))%32))+v163+v167)&int32(_a_F_tsvector_delete_by_indices_0)
									v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182))))
									v187 = v183<<(uint(v167)%32) + v172
									if v187 != 0 {
										base.MemoryCopy(m, v170+v100, v182, v187)
									} else {
									}
									v191 = v170 + v187
								} else {
									v191 = v164
								}
								v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v197 = v191
								v198 = v196
								v200 = v112 + int32(1)
								v201 = v113
							} else {
								v197 = v107
								v198 = v108
								v200 = v112
								v201 = v113 + int32(1)
							}
						}
						v204 = v109 + int32(1)
						if v204 < v198 {
							v107 = v197
							v108 = v198
							v109 = v204
							v112 = v200
							v113 = v201
							continue
						} else {
							break
						}
						break
					}
					v210 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
					v214 = v210
					v216 = v197<<(uint(int32(2))%32) + int32(32)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v82))) = v214<<(uint(int32(4))%32) + v216
				return v82
			}
		}
	} else {
		v66 = l2
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v82 = F_palloc0(m, int32(base.Ui32(v79)>>(uint(int32(2))%32)))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return int32(0)
		} else {
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v85 = v84 - v66
			*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v85
			v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v87 <= int32(0) {
				v214 = v85
				v216 = int32(32)
			} else {
				v91 = int32(8)
				v92 = l0 + v91
				v93 = int32(2)
				v97 = v82 + v91
				v100 = v97 + v85<<(uint(v93)%32)
				v101 = int32(0)
				v107 = v101
				v108 = v87
				v109 = v101
				v112 = v101
				v113 = v4
				for {
					if v66 <= v113 {
						v129 = v92 + v109<<(uint(int32(2))%32)
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
						v134 = int32(base.Ui32(v130)>>(uint(int32(1))%32)) & int32(2047)
						if v134 != 0 {
							base.MemoryCopy(m, v107+v100, v92+v16<<(uint(v93)%32)+int32(base.Ui32(v130)>>(uint(int32(12))%32)), v134)
						} else {
						}
						v142 = v97 + v112<<(uint(int32(2))%32)
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
						v144 = int32(1)
						v145 = v143 & v144
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
						*(*int32)(unsafe.Add(mBase, uint32(v142))) = v145 | v146&int32(-2)
						v151 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
						*(*int32)(unsafe.Add(mBase, uint32(v142))) = v151&int32(4094) | v107<<(uint(int32(12))%32) | v145
						v159 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
						v163 = int32(base.Ui32(v159)>>(uint(v144)%32)) & int32(2047)
						v164 = v163 + v107
						if v159&v144 != 0 {
							v167 = int32(1)
							v170 = (v164 + v167) & int32(-2)
							v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v172 = int32(2)
							v182 = v92 + v171<<(uint(v172)%32) + (int32(base.Ui32(v159)>>(uint(int32(12))%32))+v163+v167)&int32(_a_F_tsvector_delete_by_indices_0)
							v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182))))
							v187 = v183<<(uint(v167)%32) + v172
							if v187 != 0 {
								base.MemoryCopy(m, v170+v100, v182, v187)
							} else {
							}
							v191 = v170 + v187
						} else {
							v191 = v164
						}
						v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v197 = v191
						v198 = v196
						v200 = v112 + int32(1)
						v201 = v113
					} else {
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l1+v113<<(uint(int32(2))%32))))
						if v109 != v123 {
							v129 = v92 + v109<<(uint(int32(2))%32)
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
							v134 = int32(base.Ui32(v130)>>(uint(int32(1))%32)) & int32(2047)
							if v134 != 0 {
								base.MemoryCopy(m, v107+v100, v92+v16<<(uint(v93)%32)+int32(base.Ui32(v130)>>(uint(int32(12))%32)), v134)
							} else {
							}
							v142 = v97 + v112<<(uint(int32(2))%32)
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
							v144 = int32(1)
							v145 = v143 & v144
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
							*(*int32)(unsafe.Add(mBase, uint32(v142))) = v145 | v146&int32(-2)
							v151 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
							*(*int32)(unsafe.Add(mBase, uint32(v142))) = v151&int32(4094) | v107<<(uint(int32(12))%32) | v145
							v159 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
							v163 = int32(base.Ui32(v159)>>(uint(v144)%32)) & int32(2047)
							v164 = v163 + v107
							if v159&v144 != 0 {
								v167 = int32(1)
								v170 = (v164 + v167) & int32(-2)
								v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v172 = int32(2)
								v182 = v92 + v171<<(uint(v172)%32) + (int32(base.Ui32(v159)>>(uint(int32(12))%32))+v163+v167)&int32(_a_F_tsvector_delete_by_indices_0)
								v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182))))
								v187 = v183<<(uint(v167)%32) + v172
								if v187 != 0 {
									base.MemoryCopy(m, v170+v100, v182, v187)
								} else {
								}
								v191 = v170 + v187
							} else {
								v191 = v164
							}
							v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v197 = v191
							v198 = v196
							v200 = v112 + int32(1)
							v201 = v113
						} else {
							v197 = v107
							v198 = v108
							v200 = v112
							v201 = v113 + int32(1)
						}
					}
					v204 = v109 + int32(1)
					if v204 < v198 {
						v107 = v197
						v108 = v198
						v109 = v204
						v112 = v200
						v113 = v201
						continue
					} else {
						break
					}
					break
				}
				v210 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
				v214 = v210
				v216 = v197<<(uint(int32(2))%32) + int32(32)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v82))) = v214<<(uint(int32(4))%32) + v216
			return v82
		}
	}
}
func F_tsvector_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v26 = int32(2)
	v27 = int32(base.Ui32(v25) >> (uint(v26) % 32))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v30 = int32(base.Ui32(v28) >> (uint(v26) % 32))
	if base.Ui32(v27) < base.Ui32(v30) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v221 != v6 {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v220 = int32(-1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if base.Ui32(v30) < base.Ui32(v27) {
		v195 = v33
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v220 = v195
	goto L4
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v35 < v36 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v220 = int32(-1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v36 < v35 {
		v195 = v33
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v35 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v220 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	v43 = int32(8)
	v44 = v11 + v43
	v45 = int32(2)
	v47 = v44 + v36<<(uint(v45)%32)
	v49 = v6 + v43
	v52 = v49 + v35<<(uint(v45)%32)
	v61 = v44
	v62 = v49
	v66 = int32(0)
	goto L17
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v68 = int32(1)
	v69 = v67 & v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v72 = v70 & v68
	if v69 != v72 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v195 = int32(0)
	goto L8
L19:
	;
	if base.Ui32(v72) < base.Ui32(v69) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v78 = int32(1)
	v80 = int32(2047)
	v81 = int32(base.Ui32(v70)>>(uint(v78)%32)) & v80
	v82 = int32(12)
	v83 = int32(base.Ui32(v70) >> (uint(v82) % 32))
	v85 = int32(base.Ui32(v67) >> (uint(v82) % 32))
	v89 = int32(base.Ui32(v67)>>(uint(v78)%32)) & v80
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v77 = int32(-1)
	goto L24
L23:
	;
	v77 = int32(1)
	goto L24
L24:
	;
	v220 = v77
	goto L4
L25:
	;
	if v69 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	if v81 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v81 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L29:
	;
	v220 = int32(1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v95 = base.B2i32(base.Ui32(v89) < base.Ui32(v81))
	if base.Ui32(v89) < base.Ui32(v81) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v96 = v89
	goto L34
L33:
	;
	v96 = v81
	goto L34
L34:
	;
	v97 = F_memcmp(m, v85+v52, v83+v47, v96)
	mBase = m.M
	if v97 != 0 {
		v195 = v97
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v81 == v89 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v89) < base.Ui32(v81) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v101 = int32(-1)
	goto L39
L38:
	;
	v101 = int32(1)
	goto L39
L39:
	;
	v220 = v101
	goto L4
L40:
	;
	v220 = int32(-1)
	goto L4
L41:
	;
	v184 = int32(4)
	v190 = v66 + int32(1)
	if v190 != v35 {
		v61 = v61 + v184
		v62 = v62 + v184
		v66 = v190
		goto L17
	} else {
		goto L66
	}
L42:
	;
	v110 = int32(1)
	v112 = int32(_a_F_tsvector_le_0)
	v114 = v47 + (v81+v83+v110)&v112
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
	v121 = v52 + (v89+v85+v110)&v112
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121))))
	if v115 == v122 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v129 = v121
	v130 = v114
	v131 = int32(0)
	goto L51
L44:
	;
	if v122 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v115) < base.Ui32(v122) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v128 = int32(-1)
	goto L50
L49:
	;
	v128 = int32(1)
	goto L50
L50:
	;
	v220 = v128
	goto L4
L51:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+2)))
	v144 = int32(_a_F_tsvector_le_1)
	v145 = v143 & v144
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+2)))
	v148 = v146 & v144
	if v145 != v148 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v157) < base.Ui32(v155) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v148) < base.Ui32(v145) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v154 = int32(14)
	v155 = int32(base.Ui32(v143) >> (uint(v154) % 32))
	v157 = int32(base.Ui32(v146) >> (uint(v154) % 32))
	if v155 == v157 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v153 = int32(-1)
	goto L58
L57:
	;
	v153 = int32(1)
	goto L58
L58:
	;
	v220 = v153
	goto L4
L59:
	;
	v159 = int32(2)
	v164 = v131 + int32(1)
	if v164 == v122 {
		goto L41
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L52
L62:
	;
	v129 = v129 + v159
	v130 = v130 + v159
	v131 = v164
	goto L51
L63:
	;
	v169 = int32(-1)
	goto L65
L64:
	;
	v169 = int32(1)
	goto L65
L65:
	;
	v195 = v169
	goto L8
L66:
	;
	goto L18
L67:
	;
	F_pfree(m, v6)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v225 != v11 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v11)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return base.B2i32(v220 <= int32(0))
L74:
	;
	goto L73
}
func F_tsvector_to_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v16 = F_palloc(m, v13<<(uint(int32(2))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if int32(0) < v18 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = v9 + int32(8)
	v24 = v18
	v28 = int32(0)
	goto L7
L5:
	;
	v53 = v18
	goto L6
L6:
	;
	v60 = F_construct_array_builtin(m, v16, v53, int32(25))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L11
	}
L7:
	;
	v30 = int32(2)
	v31 = v28 << (uint(v30) % 32)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22+v31)))
	v45 = F_cstring_to_text_with_len(m, v22+v24<<(uint(v30)%32)+int32(base.Ui32(v37)>>(uint(int32(12))%32)), int32(base.Ui32(v37)>>(uint(int32(1))%32))&int32(2047))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v53 = v50
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v31))) = v45
	v49 = v28 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v49 < v50 {
		v24 = v50
		v28 = v49
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	F_pfree(m, v16)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v64 != v9 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_pfree(m, v9)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	return v60
L16:
	;
	goto L15
}
func F_tsvector_unnest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v17 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L38
	}
L2:
	;
	v20 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	goto L15
L5:
	;
	return int32(0)
L6:
	;
	v24 = int32(_a_F_tsvector_unnest_0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_tsvector_unnest[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_tsvector_unnest[0])) = v27
	v30 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v30
	F_TupleDescInitEntry(m, v30, int32(1), int32(_a_F_tsvector_unnest_1), int32(25), int32(-1), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v30, int32(2), int32(_a_F_tsvector_unnest_2), int32(1005), int32(-1), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v30, int32(3), int32(_a_F_tsvector_unnest_3), int32(1009), int32(-1), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v57 = F_get_call_result_type(m, l0, int32(0), v14+int32(16))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if v57 != int32(1) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v64 = F_pg_detoast_datum_copy(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v64
	*(*int32)(unsafe.Add(mBase, _c_F_tsvector_unnest[0])) = v25
	goto L4
L14:
	;
	m.G0 = v14 + int32(32)
	return v243
L15:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v73)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if base.Ui64(v74) < base.Ui64(base.I64_extend_i32_s(v76)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+30)) = uint8(v79)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+28)) = uint16(v79)
	v84 = v75 + int32(8)
	v85 = int32(2)
	v91 = v84 + base.I32_wrap_i64(v74)<<(uint(v85)%32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v100 = F_cstring_to_text_with_len(m, v84+v76<<(uint(v85)%32)+int32(base.Ui32(v92)>>(uint(int32(12))%32)), int32(base.Ui32(v92)>>(uint(int32(1))%32))&int32(2047))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L37
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v100
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v103&int32(1) != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
	v212 = F_heap_form_tuple(m, v207, v14+int32(16), v14+int32(28))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L35
	}
L21:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v107 = int32(2)
	v110 = int32(1)
	v121 = v84 + v106<<(uint(v107)%32) + (int32(base.Ui32(v103)>>(uint(v110)%32))&int32(2047)+int32(base.Ui32(v103)>>(uint(int32(12))%32))+v110)&int32(_a_F_tsvector_unnest_4)
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121))))
	v125 = F_palloc(m, v122<<(uint(v107)%32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v194 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+29)) = uint16(v194)
	goto L20
L24:
	;
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121))))
	v130 = F_palloc(m, v127<<(uint(int32(2))%32))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121))))
	if v132 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v139 = int32(0)
	goto L29
L27:
	;
	v180 = v2
	goto L28
L28:
	;
	v186 = F_construct_array_builtin(m, v125, v180, int32(21))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L33
	}
L29:
	;
	v148 = v139 << (uint(int32(2)) % 32)
	v150 = int32(1)
	v152 = v121 + int32(2) + v139<<(uint(v150)%32)
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	*(*int32)(unsafe.Add(mBase, uint32(v125+v148))) = v153 & int32(_a_F_tsvector_unnest_5)
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v161 = int32(68) - int32(base.Ui32(v158)>>(uint(int32(14))%32))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)) = uint8(v161)
	v167 = F_cstring_to_text_with_len(m, v14+int32(15), v150)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L31
	}
L30:
	;
	v180 = v172
	goto L28
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148+v130))) = v167
	v171 = v139 + int32(1)
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121))))
	if base.Ui32(v171) < base.Ui32(v172) {
		v139 = v171
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v186
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121))))
	v191 = F_construct_array_builtin(m, v130, v189, int32(25))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v191
	goto L20
L35:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v73)))
	*(*int64)(unsafe.Add(mBase, uint32(v73))) = v214 + int64(1)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v218)+20)) = int32(1)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v212)+16))
	v222 = F_HeapTupleHeaderGetDatum(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v243 = v222
	goto L14
L37:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+20)) = int32(2)
	v229 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v229)
	v243 = int32(0)
	goto L14
L38:
	;
	F_errmsg_internal(m, int32(_a_F_tsvector_unnest_6), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_tsvector_unnest_7), int32(653), int32(_a_F_tsvector_unnest_8))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
