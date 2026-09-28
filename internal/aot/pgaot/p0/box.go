package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_box_contain(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v33 int64
	_ = v33
	v3 = int64(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.F64_le(v5, base.F64_add(v7, float64(1e-06))) == int32(0) {
		v33 = v3
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
		if base.F64_le(v13, base.F64_add(v14, float64(1e-06))) == int32(0) {
			v33 = v3
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			if base.F64_le(v20, base.F64_add(v21, float64(1e-06))) == int32(0) {
				v33 = v3
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
				v33 = base.I64_extend_i32_u(base.F64_le(v27, base.F64_add(v28, float64(1e-06))))
			}
		}
	}
	return v33
}
func F_box_interpt_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v22 int64
	_ = v22
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v40 float64
	_ = v40
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v51 float64
	_ = v51
	var v57 float64
	_ = v57
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v66 int32
	_ = v66
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v74 int32
	_ = v74
	var v75 float64
	_ = v75
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v93 int32
	_ = v93
	var v94 float64
	_ = v94
	var v95 int32
	_ = v95
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v101 float64
	_ = v101
	var v103 int32
	_ = v103
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v114 float64
	_ = v114
	var v117 float64
	_ = v117
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v126 float64
	_ = v126
	var v130 float64
	_ = v130
	var v138 float64
	_ = v138
	var v140 float64
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 float64
	_ = v159
	var v160 float64
	_ = v160
	var v161 float64
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	v9 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v22 = base.I64_reinterpret_f64(v19) & int64(9223372036854775807)
	if base.Ui64(v22) <= base.Ui64(int64(9218868437227405312)) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v18)&int64(9223372036854775807)) {
			v30 = v19
		} else {
			v30 = v18
		}
		if base.F64_gt(v18, v19) != 0 {
			v32 = v19
		} else {
			v32 = v30
		}
		v33 = v32
	} else {
		v33 = v18
	}
	v34 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_le(v33, base.F64_add(v34, float64(1e-06))) == int32(0) {
		v184 = v9
		m.G0 = v16 + int32(48)
		return v184
	} else {
		v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(v22) {
			v43 = v19
		} else {
			v43 = v18
		}
		if base.F64_lt(v18, v19) != 0 {
			v45 = v19
		} else {
			v45 = v43
		}
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v18)&int64(9223372036854775807)) {
			v51 = v18
		} else {
			v51 = v45
		}
		if base.F64_le(v40, base.F64_add(v51, float64(1e-06))) == int32(0) {
			v184 = v9
			m.G0 = v16 + int32(48)
			return v184
		} else {
			v57 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			v60 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
			v61 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			v66 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v60)&int64(9223372036854775807)))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v60)&int64(9223372036854775807)) {
				v67 = v61
			} else {
				v67 = v60
			}
			if base.F64_gt(v60, v61) != 0 {
				v69 = v61
			} else {
				v69 = v67
			}
			v74 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v61)&int64(9223372036854775807)))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v61)&int64(9223372036854775807)) {
				v75 = v60
			} else {
				v75 = v69
			}
			if base.F64_ge(base.F64_add(v57, float64(1e-06)), v75) == int32(0) {
				v184 = v9
				m.G0 = v16 + int32(48)
				return v184
			} else {
				v79 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v61)&int64(9223372036854775807)) {
					v80 = v61
				} else {
					v80 = v60
				}
				if base.F64_lt(v60, v61) != 0 {
					v82 = v61
				} else {
					v82 = v80
				}
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v60)&int64(9223372036854775807)) {
					v83 = v60
				} else {
					v83 = v82
				}
				if base.F64_le(v79, base.F64_add(v83, float64(1e-06))) == int32(0) {
					v184 = v9
					m.G0 = v16 + int32(48)
					return v184
				} else {
					if l0 != 0 {
						F_box_cn(m, v16, l1, int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							v94 = F_lseg_closept_point(m, l0, l2, v16)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								v96 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
								v97 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								v98 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
								v99 = v98
								v100 = v96
								v101 = v97
								v103 = int32(0)
								if base.B2i32(base.F64_le(v99, v101) == v103)|base.B2i32(base.F64_ge(v99, v100) == v103) != 0 {
									v117 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
									if base.F64_ge(v101, v117) == int32(0) {
										v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
										v138 = v121
										*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
										v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
										v144 = int32(1)
										v147 = v16 + int32(16)
										v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
											return int32(0)
										} else {
											if v148 != 0 {
												v184 = v144
												m.G0 = v16 + int32(48)
												return v184
											} else {
												v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
												v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
												v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													if v157 != 0 {
														v184 = v144
														m.G0 = v16 + int32(48)
														return v184
													} else {
														v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
														v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
														v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int32(0)
														} else {
															if v167 != 0 {
																v184 = v144
																m.G0 = v16 + int32(48)
																return v184
															} else {
																v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int32(0)
																} else {
																	v184 = v176
																	m.G0 = v16 + int32(48)
																	return v184
																}
															}
														}
													}
												}
											}
										}
									} else {
										v122 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
										if base.F64_ge(v117, v100) == int32(0) {
											v138 = v122
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
											v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
											v144 = int32(1)
											v147 = v16 + int32(16)
											v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return int32(0)
											} else {
												if v148 != 0 {
													v184 = v144
													m.G0 = v16 + int32(48)
													return v184
												} else {
													v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
													v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
													v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														if v157 != 0 {
															v184 = v144
															m.G0 = v16 + int32(48)
															return v184
														} else {
															v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
															v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
															v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																if v167 != 0 {
																	v184 = v144
																	m.G0 = v16 + int32(48)
																	return v184
																} else {
																	v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																	v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																	v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		v184 = v176
																		m.G0 = v16 + int32(48)
																		return v184
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
											if base.F64_ge(v122, v126) == int32(0) {
												v138 = v122
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
												v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
												v144 = int32(1)
												v147 = v16 + int32(16)
												v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													if v148 != 0 {
														v184 = v144
														m.G0 = v16 + int32(48)
														return v184
													} else {
														v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
														v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
														v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															if v157 != 0 {
																v184 = v144
																m.G0 = v16 + int32(48)
																return v184
															} else {
																v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return int32(0)
																} else {
																	if v167 != 0 {
																		v184 = v144
																		m.G0 = v16 + int32(48)
																		return v184
																	} else {
																		v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																		v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																		v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			v184 = v176
																			m.G0 = v16 + int32(48)
																			return v184
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v130 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												if base.F64_le(v130, v126) == int32(0) {
													v138 = v122
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
													v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
													v144 = int32(1)
													v147 = v16 + int32(16)
													v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														if v148 != 0 {
															v184 = v144
															m.G0 = v16 + int32(48)
															return v184
														} else {
															v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
															v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
															v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																if v157 != 0 {
																	v184 = v144
																	m.G0 = v16 + int32(48)
																	return v184
																} else {
																	v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																	v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																	v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return int32(0)
																	} else {
																		if v167 != 0 {
																			v184 = v144
																			m.G0 = v16 + int32(48)
																			return v184
																		} else {
																			v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																			v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																			v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																			mBase = m.M
																			v177 = m.ExcPending
																			if v177 != 0 {
																				return int32(0)
																			} else {
																				v184 = v176
																				m.G0 = v16 + int32(48)
																				return v184
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v184 = int32(1)
													m.G0 = v16 + int32(48)
													return v184
												}
											}
										}
									}
								} else {
									v109 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
									v110 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
									if base.F64_le(v109, v110) == int32(0) {
										v117 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
										if base.F64_ge(v101, v117) == int32(0) {
											v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
											v138 = v121
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
											v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
											v144 = int32(1)
											v147 = v16 + int32(16)
											v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return int32(0)
											} else {
												if v148 != 0 {
													v184 = v144
													m.G0 = v16 + int32(48)
													return v184
												} else {
													v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
													v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
													v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														if v157 != 0 {
															v184 = v144
															m.G0 = v16 + int32(48)
															return v184
														} else {
															v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
															v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
															v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																if v167 != 0 {
																	v184 = v144
																	m.G0 = v16 + int32(48)
																	return v184
																} else {
																	v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																	v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																	v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		v184 = v176
																		m.G0 = v16 + int32(48)
																		return v184
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v122 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
											if base.F64_ge(v117, v100) == int32(0) {
												v138 = v122
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
												v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
												v144 = int32(1)
												v147 = v16 + int32(16)
												v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													if v148 != 0 {
														v184 = v144
														m.G0 = v16 + int32(48)
														return v184
													} else {
														v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
														v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
														v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															if v157 != 0 {
																v184 = v144
																m.G0 = v16 + int32(48)
																return v184
															} else {
																v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return int32(0)
																} else {
																	if v167 != 0 {
																		v184 = v144
																		m.G0 = v16 + int32(48)
																		return v184
																	} else {
																		v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																		v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																		v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			v184 = v176
																			m.G0 = v16 + int32(48)
																			return v184
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
												if base.F64_ge(v122, v126) == int32(0) {
													v138 = v122
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
													v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
													v144 = int32(1)
													v147 = v16 + int32(16)
													v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														if v148 != 0 {
															v184 = v144
															m.G0 = v16 + int32(48)
															return v184
														} else {
															v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
															v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
															v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																if v157 != 0 {
																	v184 = v144
																	m.G0 = v16 + int32(48)
																	return v184
																} else {
																	v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																	v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																	v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return int32(0)
																	} else {
																		if v167 != 0 {
																			v184 = v144
																			m.G0 = v16 + int32(48)
																			return v184
																		} else {
																			v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																			v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																			v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																			mBase = m.M
																			v177 = m.ExcPending
																			if v177 != 0 {
																				return int32(0)
																			} else {
																				v184 = v176
																				m.G0 = v16 + int32(48)
																				return v184
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v130 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													if base.F64_le(v130, v126) == int32(0) {
														v138 = v122
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
														v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
														v144 = int32(1)
														v147 = v16 + int32(16)
														v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															if v148 != 0 {
																v184 = v144
																m.G0 = v16 + int32(48)
																return v184
															} else {
																v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
																v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
																v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	if v157 != 0 {
																		v184 = v144
																		m.G0 = v16 + int32(48)
																		return v184
																	} else {
																		v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																		v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																		v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																		mBase = m.M
																		v168 = m.ExcPending
																		if v168 != 0 {
																			return int32(0)
																		} else {
																			if v167 != 0 {
																				v184 = v144
																				m.G0 = v16 + int32(48)
																				return v184
																			} else {
																				v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																				v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																				v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																				mBase = m.M
																				v177 = m.ExcPending
																				if v177 != 0 {
																					return int32(0)
																				} else {
																					v184 = v176
																					m.G0 = v16 + int32(48)
																					return v184
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v184 = int32(1)
														m.G0 = v16 + int32(48)
														return v184
													}
												}
											}
										}
									} else {
										v114 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										if base.F64_le(v114, v109) != 0 {
											v184 = int32(1)
											m.G0 = v16 + int32(48)
											return v184
										} else {
											v117 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
											if base.F64_ge(v101, v117) == int32(0) {
												v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												v138 = v121
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
												v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
												v144 = int32(1)
												v147 = v16 + int32(16)
												v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													if v148 != 0 {
														v184 = v144
														m.G0 = v16 + int32(48)
														return v184
													} else {
														v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
														v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
														v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															if v157 != 0 {
																v184 = v144
																m.G0 = v16 + int32(48)
																return v184
															} else {
																v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return int32(0)
																} else {
																	if v167 != 0 {
																		v184 = v144
																		m.G0 = v16 + int32(48)
																		return v184
																	} else {
																		v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																		v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																		v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			v184 = v176
																			m.G0 = v16 + int32(48)
																			return v184
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v122 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												if base.F64_ge(v117, v100) == int32(0) {
													v138 = v122
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
													v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
													v144 = int32(1)
													v147 = v16 + int32(16)
													v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														if v148 != 0 {
															v184 = v144
															m.G0 = v16 + int32(48)
															return v184
														} else {
															v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
															v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
															v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																if v157 != 0 {
																	v184 = v144
																	m.G0 = v16 + int32(48)
																	return v184
																} else {
																	v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																	v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																	v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return int32(0)
																	} else {
																		if v167 != 0 {
																			v184 = v144
																			m.G0 = v16 + int32(48)
																			return v184
																		} else {
																			v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																			v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																			v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																			mBase = m.M
																			v177 = m.ExcPending
																			if v177 != 0 {
																				return int32(0)
																			} else {
																				v184 = v176
																				m.G0 = v16 + int32(48)
																				return v184
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
													if base.F64_ge(v122, v126) == int32(0) {
														v138 = v122
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
														v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
														v144 = int32(1)
														v147 = v16 + int32(16)
														v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															if v148 != 0 {
																v184 = v144
																m.G0 = v16 + int32(48)
																return v184
															} else {
																v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
																v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
																v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int32(0)
																} else {
																	if v157 != 0 {
																		v184 = v144
																		m.G0 = v16 + int32(48)
																		return v184
																	} else {
																		v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																		v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																		v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																		mBase = m.M
																		v168 = m.ExcPending
																		if v168 != 0 {
																			return int32(0)
																		} else {
																			if v167 != 0 {
																				v184 = v144
																				m.G0 = v16 + int32(48)
																				return v184
																			} else {
																				v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																				v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																				v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																				mBase = m.M
																				v177 = m.ExcPending
																				if v177 != 0 {
																					return int32(0)
																				} else {
																					v184 = v176
																					m.G0 = v16 + int32(48)
																					return v184
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v130 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														if base.F64_le(v130, v126) == int32(0) {
															v138 = v122
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
															v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
															v144 = int32(1)
															v147 = v16 + int32(16)
															v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v149 = m.ExcPending
															if v149 != 0 {
																return int32(0)
															} else {
																if v148 != 0 {
																	v184 = v144
																	m.G0 = v16 + int32(48)
																	return v184
																} else {
																	v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
																	v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
																	v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int32(0)
																	} else {
																		if v157 != 0 {
																			v184 = v144
																			m.G0 = v16 + int32(48)
																			return v184
																		} else {
																			v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																			v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																			v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																			v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																			mBase = m.M
																			v168 = m.ExcPending
																			if v168 != 0 {
																				return int32(0)
																			} else {
																				if v167 != 0 {
																					v184 = v144
																					m.G0 = v16 + int32(48)
																					return v184
																				} else {
																					v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																					*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																					v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																					*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																					*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																					*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																					v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																					mBase = m.M
																					v177 = m.ExcPending
																					if v177 != 0 {
																						return int32(0)
																					} else {
																						v184 = v176
																						m.G0 = v16 + int32(48)
																						return v184
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v184 = int32(1)
															m.G0 = v16 + int32(48)
															return v184
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
						v99 = v19
						v100 = v40
						v101 = v34
						v103 = int32(0)
						if base.B2i32(base.F64_le(v99, v101) == v103)|base.B2i32(base.F64_ge(v99, v100) == v103) != 0 {
							v117 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
							if base.F64_ge(v101, v117) == int32(0) {
								v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
								v138 = v121
								*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
								v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
								*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
								*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
								v144 = int32(1)
								v147 = v16 + int32(16)
								v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return int32(0)
								} else {
									if v148 != 0 {
										v184 = v144
										m.G0 = v16 + int32(48)
										return v184
									} else {
										v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
										v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
										v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
										mBase = m.M
										v158 = m.ExcPending
										if v158 != 0 {
											return int32(0)
										} else {
											if v157 != 0 {
												v184 = v144
												m.G0 = v16 + int32(48)
												return v184
											} else {
												v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
												v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
												v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
												v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
												mBase = m.M
												v168 = m.ExcPending
												if v168 != 0 {
													return int32(0)
												} else {
													if v167 != 0 {
														v184 = v144
														m.G0 = v16 + int32(48)
														return v184
													} else {
														v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
														v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
														v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int32(0)
														} else {
															v184 = v176
															m.G0 = v16 + int32(48)
															return v184
														}
													}
												}
											}
										}
									}
								}
							} else {
								v122 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
								if base.F64_ge(v117, v100) == int32(0) {
									v138 = v122
									*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
									v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
									*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
									*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
									*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
									v144 = int32(1)
									v147 = v16 + int32(16)
									v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
										return int32(0)
									} else {
										if v148 != 0 {
											v184 = v144
											m.G0 = v16 + int32(48)
											return v184
										} else {
											v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
											v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
											v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return int32(0)
											} else {
												if v157 != 0 {
													v184 = v144
													m.G0 = v16 + int32(48)
													return v184
												} else {
													v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
													v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
													v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int32(0)
													} else {
														if v167 != 0 {
															v184 = v144
															m.G0 = v16 + int32(48)
															return v184
														} else {
															v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
															v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
															v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int32(0)
															} else {
																v184 = v176
																m.G0 = v16 + int32(48)
																return v184
															}
														}
													}
												}
											}
										}
									}
								} else {
									v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
									if base.F64_ge(v122, v126) == int32(0) {
										v138 = v122
										*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
										v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
										v144 = int32(1)
										v147 = v16 + int32(16)
										v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
											return int32(0)
										} else {
											if v148 != 0 {
												v184 = v144
												m.G0 = v16 + int32(48)
												return v184
											} else {
												v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
												v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
												v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													if v157 != 0 {
														v184 = v144
														m.G0 = v16 + int32(48)
														return v184
													} else {
														v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
														v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
														v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int32(0)
														} else {
															if v167 != 0 {
																v184 = v144
																m.G0 = v16 + int32(48)
																return v184
															} else {
																v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int32(0)
																} else {
																	v184 = v176
																	m.G0 = v16 + int32(48)
																	return v184
																}
															}
														}
													}
												}
											}
										}
									} else {
										v130 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										if base.F64_le(v130, v126) == int32(0) {
											v138 = v122
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
											v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
											v144 = int32(1)
											v147 = v16 + int32(16)
											v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return int32(0)
											} else {
												if v148 != 0 {
													v184 = v144
													m.G0 = v16 + int32(48)
													return v184
												} else {
													v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
													v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
													v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														if v157 != 0 {
															v184 = v144
															m.G0 = v16 + int32(48)
															return v184
														} else {
															v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
															v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
															v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																if v167 != 0 {
																	v184 = v144
																	m.G0 = v16 + int32(48)
																	return v184
																} else {
																	v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																	v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																	v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		v184 = v176
																		m.G0 = v16 + int32(48)
																		return v184
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v184 = int32(1)
											m.G0 = v16 + int32(48)
											return v184
										}
									}
								}
							}
						} else {
							v109 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
							v110 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							if base.F64_le(v109, v110) == int32(0) {
								v117 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
								if base.F64_ge(v101, v117) == int32(0) {
									v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
									v138 = v121
									*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
									v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
									*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
									*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
									*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
									v144 = int32(1)
									v147 = v16 + int32(16)
									v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
										return int32(0)
									} else {
										if v148 != 0 {
											v184 = v144
											m.G0 = v16 + int32(48)
											return v184
										} else {
											v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
											v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
											v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return int32(0)
											} else {
												if v157 != 0 {
													v184 = v144
													m.G0 = v16 + int32(48)
													return v184
												} else {
													v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
													v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
													v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int32(0)
													} else {
														if v167 != 0 {
															v184 = v144
															m.G0 = v16 + int32(48)
															return v184
														} else {
															v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
															v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
															v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int32(0)
															} else {
																v184 = v176
																m.G0 = v16 + int32(48)
																return v184
															}
														}
													}
												}
											}
										}
									}
								} else {
									v122 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
									if base.F64_ge(v117, v100) == int32(0) {
										v138 = v122
										*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
										v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
										v144 = int32(1)
										v147 = v16 + int32(16)
										v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
											return int32(0)
										} else {
											if v148 != 0 {
												v184 = v144
												m.G0 = v16 + int32(48)
												return v184
											} else {
												v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
												v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
												v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													if v157 != 0 {
														v184 = v144
														m.G0 = v16 + int32(48)
														return v184
													} else {
														v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
														v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
														v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int32(0)
														} else {
															if v167 != 0 {
																v184 = v144
																m.G0 = v16 + int32(48)
																return v184
															} else {
																v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int32(0)
																} else {
																	v184 = v176
																	m.G0 = v16 + int32(48)
																	return v184
																}
															}
														}
													}
												}
											}
										}
									} else {
										v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
										if base.F64_ge(v122, v126) == int32(0) {
											v138 = v122
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
											v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
											v144 = int32(1)
											v147 = v16 + int32(16)
											v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return int32(0)
											} else {
												if v148 != 0 {
													v184 = v144
													m.G0 = v16 + int32(48)
													return v184
												} else {
													v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
													v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
													v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														if v157 != 0 {
															v184 = v144
															m.G0 = v16 + int32(48)
															return v184
														} else {
															v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
															v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
															v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																if v167 != 0 {
																	v184 = v144
																	m.G0 = v16 + int32(48)
																	return v184
																} else {
																	v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																	v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																	v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		v184 = v176
																		m.G0 = v16 + int32(48)
																		return v184
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v130 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											if base.F64_le(v130, v126) == int32(0) {
												v138 = v122
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
												v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
												v144 = int32(1)
												v147 = v16 + int32(16)
												v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													if v148 != 0 {
														v184 = v144
														m.G0 = v16 + int32(48)
														return v184
													} else {
														v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
														v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
														v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															if v157 != 0 {
																v184 = v144
																m.G0 = v16 + int32(48)
																return v184
															} else {
																v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return int32(0)
																} else {
																	if v167 != 0 {
																		v184 = v144
																		m.G0 = v16 + int32(48)
																		return v184
																	} else {
																		v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																		v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																		v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			v184 = v176
																			m.G0 = v16 + int32(48)
																			return v184
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v184 = int32(1)
												m.G0 = v16 + int32(48)
												return v184
											}
										}
									}
								}
							} else {
								v114 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								if base.F64_le(v114, v109) != 0 {
									v184 = int32(1)
									m.G0 = v16 + int32(48)
									return v184
								} else {
									v117 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
									if base.F64_ge(v101, v117) == int32(0) {
										v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
										v138 = v121
										*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
										v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
										v144 = int32(1)
										v147 = v16 + int32(16)
										v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
											return int32(0)
										} else {
											if v148 != 0 {
												v184 = v144
												m.G0 = v16 + int32(48)
												return v184
											} else {
												v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
												v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
												v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													if v157 != 0 {
														v184 = v144
														m.G0 = v16 + int32(48)
														return v184
													} else {
														v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
														v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
														v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v168 = m.ExcPending
														if v168 != 0 {
															return int32(0)
														} else {
															if v167 != 0 {
																v184 = v144
																m.G0 = v16 + int32(48)
																return v184
															} else {
																v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int32(0)
																} else {
																	v184 = v176
																	m.G0 = v16 + int32(48)
																	return v184
																}
															}
														}
													}
												}
											}
										}
									} else {
										v122 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
										if base.F64_ge(v117, v100) == int32(0) {
											v138 = v122
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
											v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
											v144 = int32(1)
											v147 = v16 + int32(16)
											v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return int32(0)
											} else {
												if v148 != 0 {
													v184 = v144
													m.G0 = v16 + int32(48)
													return v184
												} else {
													v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
													v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
													v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														if v157 != 0 {
															v184 = v144
															m.G0 = v16 + int32(48)
															return v184
														} else {
															v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
															v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
															v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																if v167 != 0 {
																	v184 = v144
																	m.G0 = v16 + int32(48)
																	return v184
																} else {
																	v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																	v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																	v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int32(0)
																	} else {
																		v184 = v176
																		m.G0 = v16 + int32(48)
																		return v184
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v126 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
											if base.F64_ge(v122, v126) == int32(0) {
												v138 = v122
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
												v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
												v144 = int32(1)
												v147 = v16 + int32(16)
												v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													if v148 != 0 {
														v184 = v144
														m.G0 = v16 + int32(48)
														return v184
													} else {
														v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
														v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
														v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															if v157 != 0 {
																v184 = v144
																m.G0 = v16 + int32(48)
																return v184
															} else {
																v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																mBase = m.M
																v168 = m.ExcPending
																if v168 != 0 {
																	return int32(0)
																} else {
																	if v167 != 0 {
																		v184 = v144
																		m.G0 = v16 + int32(48)
																		return v184
																	} else {
																		v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																		v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																		v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			v184 = v176
																			m.G0 = v16 + int32(48)
																			return v184
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v130 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												if base.F64_le(v130, v126) == int32(0) {
													v138 = v122
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v100
													v140 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v140
													v144 = int32(1)
													v147 = v16 + int32(16)
													v148 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														if v148 != 0 {
															v184 = v144
															m.G0 = v16 + int32(48)
															return v184
														} else {
															v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v150
															v152 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v138
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v100
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v152
															v157 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
																return int32(0)
															} else {
																if v157 != 0 {
																	v184 = v144
																	m.G0 = v16 + int32(48)
																	return v184
																} else {
																	v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																	v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
																	v167 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return int32(0)
																	} else {
																		if v167 != 0 {
																			v184 = v144
																			m.G0 = v16 + int32(48)
																			return v184
																		} else {
																			v169 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v169
																			v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v171
																			v176 = F_lseg_interpt_lseg(m, int32(0), v147, l2)
																			mBase = m.M
																			v177 = m.ExcPending
																			if v177 != 0 {
																				return int32(0)
																			} else {
																				v184 = v176
																				m.G0 = v16 + int32(48)
																				return v184
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v184 = int32(1)
													m.G0 = v16 + int32(48)
													return v184
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
func F_box_lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v11 float64
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_box_ar(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v11 = F_box_ar(m, v3)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(base.F64_lt(base.F64_add(v5, float64(1e-06)), v11))
		}
	}
}
func F_box_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_path_encode(m, int32(0), int32(2), v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v5)
	}
}
func F_box_right(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.I64_extend_i32_u(base.F64_gt(v3, base.F64_add(v5, float64(1e-06))))
}
func F_box_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		F_pq_sendfloat8(m, v5, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			F_pq_sendfloat8(m, v5, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				v18 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
				F_pq_sendfloat8(m, v5, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
					F_pq_sendfloat8(m, v5, v21)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v25))) = v26 << (uint(int32(2)) % 32)
						m.G0 = v5 + int32(16)
						return base.I64_extend_i32_u(v25)
					}
				}
			}
		}
	}
}
func F_box_sub(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v19 float64
	_ = v19
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v79 float64
	_ = v79
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_palloc(m, int32(32))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v17 = base.F64_sub(v15, v16)
		v19 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19)|base.F64_eq(base.F64_abs(v16), v19) == int32(0) {
			v32 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int64(0)
			} else {
				v34 = v32
				v35 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
				v36 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
				v37 = base.F64_sub(v35, v36)
				v39 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v37), v39)|base.F64_eq(base.F64_abs(v35), v39)|base.F64_eq(base.F64_abs(v36), v39) != 0 {
					v52 = v37
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
					v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
					v56 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
					v57 = base.F64_sub(v55, v56)
					v59 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v55), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
						v72 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int64(0)
						} else {
							v74 = v72
							v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
							v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
							v77 = base.F64_sub(v75, v76)
							v79 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
								v92 = v77
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							} else {
								v90 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int64(0)
								} else {
									v92 = v90
									*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
									return base.I64_extend_i32_u(v11)
								}
							}
						}
					} else {
						v74 = v57
						v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
						v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
						v77 = base.F64_sub(v75, v76)
						v79 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
							v92 = v77
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
							return base.I64_extend_i32_u(v11)
						} else {
							v90 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								v92 = v90
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							}
						}
					}
				} else {
					v50 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						v52 = v50
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
						v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
						v56 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
						v57 = base.F64_sub(v55, v56)
						v59 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v55), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
							v72 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								v74 = v72
								v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
								v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
								v77 = base.F64_sub(v75, v76)
								v79 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
									v92 = v77
									*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
									return base.I64_extend_i32_u(v11)
								} else {
									v90 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int64(0)
									} else {
										v92 = v90
										*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
										*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
										return base.I64_extend_i32_u(v11)
									}
								}
							}
						} else {
							v74 = v57
							v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
							v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
							v77 = base.F64_sub(v75, v76)
							v79 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
								v92 = v77
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							} else {
								v90 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int64(0)
								} else {
									v92 = v90
									*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
									return base.I64_extend_i32_u(v11)
								}
							}
						}
					}
				}
			}
		} else {
			v34 = v17
			v35 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v36 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v37 = base.F64_sub(v35, v36)
			v39 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v37), v39)|base.F64_eq(base.F64_abs(v35), v39)|base.F64_eq(base.F64_abs(v36), v39) != 0 {
				v52 = v37
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
				v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v56 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
				v57 = base.F64_sub(v55, v56)
				v59 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v55), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
					v72 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int64(0)
					} else {
						v74 = v72
						v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
						v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
						v77 = base.F64_sub(v75, v76)
						v79 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
							v92 = v77
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
							return base.I64_extend_i32_u(v11)
						} else {
							v90 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								v92 = v90
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							}
						}
					}
				} else {
					v74 = v57
					v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
					v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
					v77 = base.F64_sub(v75, v76)
					v79 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
						v92 = v77
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
						return base.I64_extend_i32_u(v11)
					} else {
						v90 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int64(0)
						} else {
							v92 = v90
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
							return base.I64_extend_i32_u(v11)
						}
					}
				}
			} else {
				v50 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					v52 = v50
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v34
					v55 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
					v56 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
					v57 = base.F64_sub(v55, v56)
					v59 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v55), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
						v72 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int64(0)
						} else {
							v74 = v72
							v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
							v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
							v77 = base.F64_sub(v75, v76)
							v79 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
								v92 = v77
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							} else {
								v90 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int64(0)
								} else {
									v92 = v90
									*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
									return base.I64_extend_i32_u(v11)
								}
							}
						}
					} else {
						v74 = v57
						v75 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
						v76 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
						v77 = base.F64_sub(v75, v76)
						v79 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v77), v79)|base.F64_eq(base.F64_abs(v75), v79)|base.F64_eq(base.F64_abs(v76), v79) != 0 {
							v92 = v77
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
							return base.I64_extend_i32_u(v11)
						} else {
							v90 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								v92 = v90
								*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v92
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v74
								return base.I64_extend_i32_u(v11)
							}
						}
					}
				}
			}
		}
	}
}
func F_box_width(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 float64
	_ = v10
	var v21 float64
	_ = v21
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v5)+16))
	v8 = base.F64_sub(v6, v7)
	v10 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v8), v10)|base.F64_eq(base.F64_abs(v6), v10)|base.F64_eq(base.F64_abs(v7), v10) != 0 {
		v25 = v8
		return base.I64_reinterpret_f64(v25)
	} else {
		v21 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v25 = v21
			return base.I64_reinterpret_f64(v25)
		}
	}
}
