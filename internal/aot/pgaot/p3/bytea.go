package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bytea_bit_count(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int64
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int64
	_ = v172
	var v182 int64
	_ = v182
	v6 = int64(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = int32(1)
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		v17 = v15 & v13
		if v17 != 0 {
			v18 = v13
		} else {
			v18 = int32(4)
		}
		v19 = v9 + v18
		if v15 == int32(1) {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v49 = int32(4)
				v52 = v49 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v49) {
					v58 = v19
					v61 = int32(0)
					v63 = v6
					for {
						v64 = int32(4)
						v65 = v58 + v64
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)))
						v67 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v66)+uint32(_c_F_bytea_bit_count[0]))))
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
						v69 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_bytea_bit_count[0]))))
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
						v71 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_bytea_bit_count[0]))))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
						v73 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_bytea_bit_count[0]))))
						v77 = v67 + (v69 + (v71 + (v63 + v73)))
						v79 = v61 + v64
						if v79 != v49&int32(-4) {
							v58 = v65
							v61 = v79
							v63 = v77
							continue
						} else {
							break
						}
						break
					}
					if v52 == int32(0) {
						v109 = v77
					} else {
						v83 = v65
						v88 = v77
						v90 = v83
						v91 = int32(0)
						v95 = v88
						for {
							v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
							v97 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_bytea_bit_count[0]))))
							v98 = v95 + v97
							v99 = int32(1)
							v102 = v91 + v99
							if v102 != v52 {
								v90 = v90 + v99
								v91 = v102
								v95 = v98
								continue
							} else {
								break
							}
							break
						}
						v109 = v98
					}
				} else {
					v83 = v19
					v88 = v6
					v90 = v83
					v91 = int32(0)
					v95 = v88
					for {
						v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
						v97 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_bytea_bit_count[0]))))
						v98 = v95 + v97
						v99 = int32(1)
						v102 = v91 + v99
						if v102 != v52 {
							v90 = v90 + v99
							v91 = v102
							v95 = v98
							continue
						} else {
							break
						}
						break
					}
					v109 = v98
				}
				return v109
			} else {
				if v22 == int32(18) {
					v33 = int32(16)
				} else {
					v33 = int32(0)
				}
				v44 = v33
				if int32(7) < v44 {
					v111 = int64(0)
					v112 = int32(0)
					if v44 == v112 {
						v182 = int64(0)
					} else {
						v119 = v44 & int32(3)
						if base.Ui32(int32(4)) <= base.Ui32(v44) {
							v124 = v19
							v126 = v111
							v129 = v112
							for {
								v130 = int32(4)
								v131 = v124 + v130
								v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+3)))
								v133 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_bytea_bit_count[0]))))
								v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+2)))
								v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v134)+uint32(_c_F_bytea_bit_count[0]))))
								v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)))
								v137 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_bytea_bit_count[0]))))
								v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
								v139 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_bytea_bit_count[0]))))
								v143 = v133 + (v135 + (v137 + (v126 + v139)))
								v145 = v129 + v130
								if v145 != v44&int32(-4) {
									v124 = v131
									v126 = v143
									v129 = v145
									continue
								} else {
									break
								}
								break
							}
							if v119 == int32(0) {
								v172 = v143
							} else {
								v149 = v131
								v151 = v143
								v156 = v149
								v157 = int32(0)
								v158 = v151
								for {
									v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
									v163 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v162)+uint32(_c_F_bytea_bit_count[0]))))
									v164 = v158 + v163
									v165 = int32(1)
									v168 = v157 + v165
									if v168 != v119 {
										v156 = v156 + v165
										v157 = v168
										v158 = v164
										continue
									} else {
										break
									}
									break
								}
								v172 = v164
							}
						} else {
							v149 = v19
							v151 = v111
							v156 = v149
							v157 = int32(0)
							v158 = v151
							for {
								v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
								v163 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v162)+uint32(_c_F_bytea_bit_count[0]))))
								v164 = v158 + v163
								v165 = int32(1)
								v168 = v157 + v165
								if v168 != v119 {
									v156 = v156 + v165
									v157 = v168
									v158 = v164
									continue
								} else {
									break
								}
								break
							}
							v172 = v164
						}
						v182 = v172
					}
					return v182
				} else {
					if v44 != 0 {
						v49 = v44
						v52 = v49 & int32(3)
						if base.Ui32(int32(4)) <= base.Ui32(v49) {
							v58 = v19
							v61 = int32(0)
							v63 = v6
							for {
								v64 = int32(4)
								v65 = v58 + v64
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)))
								v67 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v66)+uint32(_c_F_bytea_bit_count[0]))))
								v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
								v69 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_bytea_bit_count[0]))))
								v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
								v71 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_bytea_bit_count[0]))))
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
								v73 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_bytea_bit_count[0]))))
								v77 = v67 + (v69 + (v71 + (v63 + v73)))
								v79 = v61 + v64
								if v79 != v49&int32(-4) {
									v58 = v65
									v61 = v79
									v63 = v77
									continue
								} else {
									break
								}
								break
							}
							if v52 == int32(0) {
								v109 = v77
							} else {
								v83 = v65
								v88 = v77
								v90 = v83
								v91 = int32(0)
								v95 = v88
								for {
									v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
									v97 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_bytea_bit_count[0]))))
									v98 = v95 + v97
									v99 = int32(1)
									v102 = v91 + v99
									if v102 != v52 {
										v90 = v90 + v99
										v91 = v102
										v95 = v98
										continue
									} else {
										break
									}
									break
								}
								v109 = v98
							}
						} else {
							v83 = v19
							v88 = v6
							v90 = v83
							v91 = int32(0)
							v95 = v88
							for {
								v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
								v97 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_bytea_bit_count[0]))))
								v98 = v95 + v97
								v99 = int32(1)
								v102 = v91 + v99
								if v102 != v52 {
									v90 = v90 + v99
									v91 = v102
									v95 = v98
									continue
								} else {
									break
								}
								break
							}
							v109 = v98
						}
						return v109
					} else {
						return int64(0)
					}
				}
			}
		} else {
			v34 = int32(1)
			if v17 != 0 {
				v44 = int32(base.Ui32(v15)>>(uint(v34)%32)) - v34
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
			}
			if int32(7) < v44 {
				v111 = int64(0)
				v112 = int32(0)
				if v44 == v112 {
					v182 = int64(0)
				} else {
					v119 = v44 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(v44) {
						v124 = v19
						v126 = v111
						v129 = v112
						for {
							v130 = int32(4)
							v131 = v124 + v130
							v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+3)))
							v133 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v132)+uint32(_c_F_bytea_bit_count[0]))))
							v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+2)))
							v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v134)+uint32(_c_F_bytea_bit_count[0]))))
							v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)))
							v137 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_bytea_bit_count[0]))))
							v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
							v139 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v138)+uint32(_c_F_bytea_bit_count[0]))))
							v143 = v133 + (v135 + (v137 + (v126 + v139)))
							v145 = v129 + v130
							if v145 != v44&int32(-4) {
								v124 = v131
								v126 = v143
								v129 = v145
								continue
							} else {
								break
							}
							break
						}
						if v119 == int32(0) {
							v172 = v143
						} else {
							v149 = v131
							v151 = v143
							v156 = v149
							v157 = int32(0)
							v158 = v151
							for {
								v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
								v163 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v162)+uint32(_c_F_bytea_bit_count[0]))))
								v164 = v158 + v163
								v165 = int32(1)
								v168 = v157 + v165
								if v168 != v119 {
									v156 = v156 + v165
									v157 = v168
									v158 = v164
									continue
								} else {
									break
								}
								break
							}
							v172 = v164
						}
					} else {
						v149 = v19
						v151 = v111
						v156 = v149
						v157 = int32(0)
						v158 = v151
						for {
							v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
							v163 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v162)+uint32(_c_F_bytea_bit_count[0]))))
							v164 = v158 + v163
							v165 = int32(1)
							v168 = v157 + v165
							if v168 != v119 {
								v156 = v156 + v165
								v157 = v168
								v158 = v164
								continue
							} else {
								break
							}
							break
						}
						v172 = v164
					}
					v182 = v172
				}
				return v182
			} else {
				if v44 != 0 {
					v49 = v44
					v52 = v49 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(v49) {
						v58 = v19
						v61 = int32(0)
						v63 = v6
						for {
							v64 = int32(4)
							v65 = v58 + v64
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)))
							v67 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v66)+uint32(_c_F_bytea_bit_count[0]))))
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
							v69 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_bytea_bit_count[0]))))
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
							v71 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_bytea_bit_count[0]))))
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
							v73 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_bytea_bit_count[0]))))
							v77 = v67 + (v69 + (v71 + (v63 + v73)))
							v79 = v61 + v64
							if v79 != v49&int32(-4) {
								v58 = v65
								v61 = v79
								v63 = v77
								continue
							} else {
								break
							}
							break
						}
						if v52 == int32(0) {
							v109 = v77
						} else {
							v83 = v65
							v88 = v77
							v90 = v83
							v91 = int32(0)
							v95 = v88
							for {
								v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
								v97 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_bytea_bit_count[0]))))
								v98 = v95 + v97
								v99 = int32(1)
								v102 = v91 + v99
								if v102 != v52 {
									v90 = v90 + v99
									v91 = v102
									v95 = v98
									continue
								} else {
									break
								}
								break
							}
							v109 = v98
						}
					} else {
						v83 = v19
						v88 = v6
						v90 = v83
						v91 = int32(0)
						v95 = v88
						for {
							v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
							v97 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_bytea_bit_count[0]))))
							v98 = v95 + v97
							v99 = int32(1)
							v102 = v91 + v99
							if v102 != v52 {
								v90 = v90 + v99
								v91 = v102
								v95 = v98
								continue
							} else {
								break
							}
							break
						}
						v109 = v98
					}
					return v109
				} else {
					return int64(0)
				}
			}
		}
	}
}
func F_bytea_substr_no_len(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int32(1)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4 <= v3 {
		v7 = v3
	} else {
		v7 = v4
	}
	v11 = F_detoast_attr_slice(m, v2, v7-int32(1), int32(-1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v11)
	}
}
