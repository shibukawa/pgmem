package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_box_contain(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v32 int32
	_ = v32
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.F64_le(v5, base.F64_add(v7, float64(1e-06))) == v2 {
		v32 = v2
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
		if base.F64_le(v13, base.F64_add(v14, float64(1e-06))) == int32(0) {
			v32 = v2
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			if base.F64_le(v20, base.F64_add(v21, float64(1e-06))) == int32(0) {
				v32 = v2
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
				v32 = base.F64_le(v27, base.F64_add(v28, float64(1e-06)))
			}
		}
	}
	return v32
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
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v95 float64
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
	var v102 int32
	_ = v102
	var v108 float64
	_ = v108
	var v109 float64
	_ = v109
	var v113 float64
	_ = v113
	var v116 float64
	_ = v116
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v125 float64
	_ = v125
	var v129 float64
	_ = v129
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 float64
	_ = v149
	var v151 float64
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 float64
	_ = v158
	var v159 float64
	_ = v159
	var v160 float64
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 float64
	_ = v168
	var v170 float64
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
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
		v183 = v9
		m.G0 = v16 + int32(48)
		return v183
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
			v183 = v9
			m.G0 = v16 + int32(48)
			return v183
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
				v183 = v9
				m.G0 = v16 + int32(48)
				return v183
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
					v183 = v9
					m.G0 = v16 + int32(48)
					return v183
				} else {
					if l0 != 0 {
						F_box_cn(m, v16, l1)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							v93 = F_lseg_closept_point(m, l0, l2, v16)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								v95 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
								v96 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
								v97 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
								v98 = v97
								v99 = v95
								v100 = v96
								v102 = int32(0)
								if base.B2i32(base.F64_le(v98, v100) == v102)|base.B2i32(base.F64_ge(v98, v99) == v102) != 0 {
									v116 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
									if base.F64_ge(v100, v116) == int32(0) {
										v120 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
										v137 = v120
										*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
										v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
										v143 = int32(1)
										v146 = v16 + int32(16)
										v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
											return int32(0)
										} else {
											if v147 != 0 {
												v183 = v143
												m.G0 = v16 + int32(48)
												return v183
											} else {
												v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
												v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
												v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return int32(0)
												} else {
													if v156 != 0 {
														v183 = v143
														m.G0 = v16 + int32(48)
														return v183
													} else {
														v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
														v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
														v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return int32(0)
														} else {
															if v166 != 0 {
																v183 = v143
																m.G0 = v16 + int32(48)
																return v183
															} else {
																v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return int32(0)
																} else {
																	v183 = v175
																	m.G0 = v16 + int32(48)
																	return v183
																}
															}
														}
													}
												}
											}
										}
									} else {
										v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
										if base.F64_ge(v116, v99) == int32(0) {
											v137 = v121
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
											v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
											v143 = int32(1)
											v146 = v16 + int32(16)
											v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return int32(0)
											} else {
												if v147 != 0 {
													v183 = v143
													m.G0 = v16 + int32(48)
													return v183
												} else {
													v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
													v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
													v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														if v156 != 0 {
															v183 = v143
															m.G0 = v16 + int32(48)
															return v183
														} else {
															v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
															v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
															v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v167 = m.ExcPending
															if v167 != 0 {
																return int32(0)
															} else {
																if v166 != 0 {
																	v183 = v143
																	m.G0 = v16 + int32(48)
																	return v183
																} else {
																	v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																	v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																	v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return int32(0)
																	} else {
																		v183 = v175
																		m.G0 = v16 + int32(48)
																		return v183
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v125 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
											if base.F64_ge(v121, v125) == int32(0) {
												v137 = v121
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
												v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
												v143 = int32(1)
												v146 = v16 + int32(16)
												v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													if v147 != 0 {
														v183 = v143
														m.G0 = v16 + int32(48)
														return v183
													} else {
														v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
														v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
														v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															if v156 != 0 {
																v183 = v143
																m.G0 = v16 + int32(48)
																return v183
															} else {
																v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																mBase = m.M
																v167 = m.ExcPending
																if v167 != 0 {
																	return int32(0)
																} else {
																	if v166 != 0 {
																		v183 = v143
																		m.G0 = v16 + int32(48)
																		return v183
																	} else {
																		v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																		v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																		v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																		mBase = m.M
																		v176 = m.ExcPending
																		if v176 != 0 {
																			return int32(0)
																		} else {
																			v183 = v175
																			m.G0 = v16 + int32(48)
																			return v183
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v129 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												if base.F64_le(v129, v125) == int32(0) {
													v137 = v121
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
													v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
													v143 = int32(1)
													v146 = v16 + int32(16)
													v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														if v147 != 0 {
															v183 = v143
															m.G0 = v16 + int32(48)
															return v183
														} else {
															v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
															v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
															v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																if v156 != 0 {
																	v183 = v143
																	m.G0 = v16 + int32(48)
																	return v183
																} else {
																	v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																	v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																	v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																	mBase = m.M
																	v167 = m.ExcPending
																	if v167 != 0 {
																		return int32(0)
																	} else {
																		if v166 != 0 {
																			v183 = v143
																			m.G0 = v16 + int32(48)
																			return v183
																		} else {
																			v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																			v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																			v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																			mBase = m.M
																			v176 = m.ExcPending
																			if v176 != 0 {
																				return int32(0)
																			} else {
																				v183 = v175
																				m.G0 = v16 + int32(48)
																				return v183
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v183 = int32(1)
													m.G0 = v16 + int32(48)
													return v183
												}
											}
										}
									}
								} else {
									v108 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
									v109 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
									if base.F64_le(v108, v109) == int32(0) {
										v116 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
										if base.F64_ge(v100, v116) == int32(0) {
											v120 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
											v137 = v120
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
											v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
											v143 = int32(1)
											v146 = v16 + int32(16)
											v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return int32(0)
											} else {
												if v147 != 0 {
													v183 = v143
													m.G0 = v16 + int32(48)
													return v183
												} else {
													v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
													v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
													v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														if v156 != 0 {
															v183 = v143
															m.G0 = v16 + int32(48)
															return v183
														} else {
															v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
															v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
															v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v167 = m.ExcPending
															if v167 != 0 {
																return int32(0)
															} else {
																if v166 != 0 {
																	v183 = v143
																	m.G0 = v16 + int32(48)
																	return v183
																} else {
																	v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																	v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																	v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return int32(0)
																	} else {
																		v183 = v175
																		m.G0 = v16 + int32(48)
																		return v183
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
											if base.F64_ge(v116, v99) == int32(0) {
												v137 = v121
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
												v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
												v143 = int32(1)
												v146 = v16 + int32(16)
												v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													if v147 != 0 {
														v183 = v143
														m.G0 = v16 + int32(48)
														return v183
													} else {
														v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
														v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
														v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															if v156 != 0 {
																v183 = v143
																m.G0 = v16 + int32(48)
																return v183
															} else {
																v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																mBase = m.M
																v167 = m.ExcPending
																if v167 != 0 {
																	return int32(0)
																} else {
																	if v166 != 0 {
																		v183 = v143
																		m.G0 = v16 + int32(48)
																		return v183
																	} else {
																		v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																		v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																		v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																		mBase = m.M
																		v176 = m.ExcPending
																		if v176 != 0 {
																			return int32(0)
																		} else {
																			v183 = v175
																			m.G0 = v16 + int32(48)
																			return v183
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v125 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
												if base.F64_ge(v121, v125) == int32(0) {
													v137 = v121
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
													v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
													v143 = int32(1)
													v146 = v16 + int32(16)
													v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														if v147 != 0 {
															v183 = v143
															m.G0 = v16 + int32(48)
															return v183
														} else {
															v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
															v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
															v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																if v156 != 0 {
																	v183 = v143
																	m.G0 = v16 + int32(48)
																	return v183
																} else {
																	v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																	v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																	v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																	mBase = m.M
																	v167 = m.ExcPending
																	if v167 != 0 {
																		return int32(0)
																	} else {
																		if v166 != 0 {
																			v183 = v143
																			m.G0 = v16 + int32(48)
																			return v183
																		} else {
																			v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																			v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																			v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																			mBase = m.M
																			v176 = m.ExcPending
																			if v176 != 0 {
																				return int32(0)
																			} else {
																				v183 = v175
																				m.G0 = v16 + int32(48)
																				return v183
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v129 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													if base.F64_le(v129, v125) == int32(0) {
														v137 = v121
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
														v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
														v143 = int32(1)
														v146 = v16 + int32(16)
														v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return int32(0)
														} else {
															if v147 != 0 {
																v183 = v143
																m.G0 = v16 + int32(48)
																return v183
															} else {
																v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
																v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
																v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	if v156 != 0 {
																		v183 = v143
																		m.G0 = v16 + int32(48)
																		return v183
																	} else {
																		v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																		v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																		v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																		mBase = m.M
																		v167 = m.ExcPending
																		if v167 != 0 {
																			return int32(0)
																		} else {
																			if v166 != 0 {
																				v183 = v143
																				m.G0 = v16 + int32(48)
																				return v183
																			} else {
																				v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																				v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																				v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																				mBase = m.M
																				v176 = m.ExcPending
																				if v176 != 0 {
																					return int32(0)
																				} else {
																					v183 = v175
																					m.G0 = v16 + int32(48)
																					return v183
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v183 = int32(1)
														m.G0 = v16 + int32(48)
														return v183
													}
												}
											}
										}
									} else {
										v113 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										if base.F64_le(v113, v108) != 0 {
											v183 = int32(1)
											m.G0 = v16 + int32(48)
											return v183
										} else {
											v116 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
											if base.F64_ge(v100, v116) == int32(0) {
												v120 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												v137 = v120
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
												v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
												v143 = int32(1)
												v146 = v16 + int32(16)
												v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													if v147 != 0 {
														v183 = v143
														m.G0 = v16 + int32(48)
														return v183
													} else {
														v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
														v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
														v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															if v156 != 0 {
																v183 = v143
																m.G0 = v16 + int32(48)
																return v183
															} else {
																v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																mBase = m.M
																v167 = m.ExcPending
																if v167 != 0 {
																	return int32(0)
																} else {
																	if v166 != 0 {
																		v183 = v143
																		m.G0 = v16 + int32(48)
																		return v183
																	} else {
																		v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																		v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																		v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																		mBase = m.M
																		v176 = m.ExcPending
																		if v176 != 0 {
																			return int32(0)
																		} else {
																			v183 = v175
																			m.G0 = v16 + int32(48)
																			return v183
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												if base.F64_ge(v116, v99) == int32(0) {
													v137 = v121
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
													v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
													v143 = int32(1)
													v146 = v16 + int32(16)
													v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														if v147 != 0 {
															v183 = v143
															m.G0 = v16 + int32(48)
															return v183
														} else {
															v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
															v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
															v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																if v156 != 0 {
																	v183 = v143
																	m.G0 = v16 + int32(48)
																	return v183
																} else {
																	v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																	v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																	v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																	mBase = m.M
																	v167 = m.ExcPending
																	if v167 != 0 {
																		return int32(0)
																	} else {
																		if v166 != 0 {
																			v183 = v143
																			m.G0 = v16 + int32(48)
																			return v183
																		} else {
																			v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																			v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																			v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																			mBase = m.M
																			v176 = m.ExcPending
																			if v176 != 0 {
																				return int32(0)
																			} else {
																				v183 = v175
																				m.G0 = v16 + int32(48)
																				return v183
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v125 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
													if base.F64_ge(v121, v125) == int32(0) {
														v137 = v121
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
														v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
														v143 = int32(1)
														v146 = v16 + int32(16)
														v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return int32(0)
														} else {
															if v147 != 0 {
																v183 = v143
																m.G0 = v16 + int32(48)
																return v183
															} else {
																v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
																v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
																v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int32(0)
																} else {
																	if v156 != 0 {
																		v183 = v143
																		m.G0 = v16 + int32(48)
																		return v183
																	} else {
																		v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																		v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																		v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																		mBase = m.M
																		v167 = m.ExcPending
																		if v167 != 0 {
																			return int32(0)
																		} else {
																			if v166 != 0 {
																				v183 = v143
																				m.G0 = v16 + int32(48)
																				return v183
																			} else {
																				v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																				v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																				*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																				v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																				mBase = m.M
																				v176 = m.ExcPending
																				if v176 != 0 {
																					return int32(0)
																				} else {
																					v183 = v175
																					m.G0 = v16 + int32(48)
																					return v183
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v129 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														if base.F64_le(v129, v125) == int32(0) {
															v137 = v121
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
															v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
															v143 = int32(1)
															v146 = v16 + int32(16)
															v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
																return int32(0)
															} else {
																if v147 != 0 {
																	v183 = v143
																	m.G0 = v16 + int32(48)
																	return v183
																} else {
																	v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
																	v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
																	v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																	mBase = m.M
																	v157 = m.ExcPending
																	if v157 != 0 {
																		return int32(0)
																	} else {
																		if v156 != 0 {
																			v183 = v143
																			m.G0 = v16 + int32(48)
																			return v183
																		} else {
																			v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																			v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																			v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																			v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																			mBase = m.M
																			v167 = m.ExcPending
																			if v167 != 0 {
																				return int32(0)
																			} else {
																				if v166 != 0 {
																					v183 = v143
																					m.G0 = v16 + int32(48)
																					return v183
																				} else {
																					v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																					*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																					v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																					*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																					*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																					*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																					v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																					mBase = m.M
																					v176 = m.ExcPending
																					if v176 != 0 {
																						return int32(0)
																					} else {
																						v183 = v175
																						m.G0 = v16 + int32(48)
																						return v183
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v183 = int32(1)
															m.G0 = v16 + int32(48)
															return v183
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
						v98 = v19
						v99 = v40
						v100 = v34
						v102 = int32(0)
						if base.B2i32(base.F64_le(v98, v100) == v102)|base.B2i32(base.F64_ge(v98, v99) == v102) != 0 {
							v116 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
							if base.F64_ge(v100, v116) == int32(0) {
								v120 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
								v137 = v120
								*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
								v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
								*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
								*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
								v143 = int32(1)
								v146 = v16 + int32(16)
								v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
									return int32(0)
								} else {
									if v147 != 0 {
										v183 = v143
										m.G0 = v16 + int32(48)
										return v183
									} else {
										v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
										v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
										v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return int32(0)
										} else {
											if v156 != 0 {
												v183 = v143
												m.G0 = v16 + int32(48)
												return v183
											} else {
												v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
												v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
												v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
												v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													if v166 != 0 {
														v183 = v143
														m.G0 = v16 + int32(48)
														return v183
													} else {
														v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
														v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
														v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
															return int32(0)
														} else {
															v183 = v175
															m.G0 = v16 + int32(48)
															return v183
														}
													}
												}
											}
										}
									}
								}
							} else {
								v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
								if base.F64_ge(v116, v99) == int32(0) {
									v137 = v121
									*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
									v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
									*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
									*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
									*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
									v143 = int32(1)
									v146 = v16 + int32(16)
									v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
									mBase = m.M
									v148 = m.ExcPending
									if v148 != 0 {
										return int32(0)
									} else {
										if v147 != 0 {
											v183 = v143
											m.G0 = v16 + int32(48)
											return v183
										} else {
											v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
											v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
											v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return int32(0)
											} else {
												if v156 != 0 {
													v183 = v143
													m.G0 = v16 + int32(48)
													return v183
												} else {
													v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
													v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
													v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														if v166 != 0 {
															v183 = v143
															m.G0 = v16 + int32(48)
															return v183
														} else {
															v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
															v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
															v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return int32(0)
															} else {
																v183 = v175
																m.G0 = v16 + int32(48)
																return v183
															}
														}
													}
												}
											}
										}
									}
								} else {
									v125 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
									if base.F64_ge(v121, v125) == int32(0) {
										v137 = v121
										*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
										v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
										v143 = int32(1)
										v146 = v16 + int32(16)
										v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
											return int32(0)
										} else {
											if v147 != 0 {
												v183 = v143
												m.G0 = v16 + int32(48)
												return v183
											} else {
												v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
												v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
												v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return int32(0)
												} else {
													if v156 != 0 {
														v183 = v143
														m.G0 = v16 + int32(48)
														return v183
													} else {
														v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
														v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
														v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return int32(0)
														} else {
															if v166 != 0 {
																v183 = v143
																m.G0 = v16 + int32(48)
																return v183
															} else {
																v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return int32(0)
																} else {
																	v183 = v175
																	m.G0 = v16 + int32(48)
																	return v183
																}
															}
														}
													}
												}
											}
										}
									} else {
										v129 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										if base.F64_le(v129, v125) == int32(0) {
											v137 = v121
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
											v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
											v143 = int32(1)
											v146 = v16 + int32(16)
											v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return int32(0)
											} else {
												if v147 != 0 {
													v183 = v143
													m.G0 = v16 + int32(48)
													return v183
												} else {
													v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
													v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
													v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														if v156 != 0 {
															v183 = v143
															m.G0 = v16 + int32(48)
															return v183
														} else {
															v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
															v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
															v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v167 = m.ExcPending
															if v167 != 0 {
																return int32(0)
															} else {
																if v166 != 0 {
																	v183 = v143
																	m.G0 = v16 + int32(48)
																	return v183
																} else {
																	v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																	v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																	v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return int32(0)
																	} else {
																		v183 = v175
																		m.G0 = v16 + int32(48)
																		return v183
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v183 = int32(1)
											m.G0 = v16 + int32(48)
											return v183
										}
									}
								}
							}
						} else {
							v108 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
							v109 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							if base.F64_le(v108, v109) == int32(0) {
								v116 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
								if base.F64_ge(v100, v116) == int32(0) {
									v120 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
									v137 = v120
									*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
									v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
									*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
									*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
									*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
									v143 = int32(1)
									v146 = v16 + int32(16)
									v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
									mBase = m.M
									v148 = m.ExcPending
									if v148 != 0 {
										return int32(0)
									} else {
										if v147 != 0 {
											v183 = v143
											m.G0 = v16 + int32(48)
											return v183
										} else {
											v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
											v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
											v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return int32(0)
											} else {
												if v156 != 0 {
													v183 = v143
													m.G0 = v16 + int32(48)
													return v183
												} else {
													v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
													v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
													v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														if v166 != 0 {
															v183 = v143
															m.G0 = v16 + int32(48)
															return v183
														} else {
															v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
															v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
															v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return int32(0)
															} else {
																v183 = v175
																m.G0 = v16 + int32(48)
																return v183
															}
														}
													}
												}
											}
										}
									}
								} else {
									v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
									if base.F64_ge(v116, v99) == int32(0) {
										v137 = v121
										*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
										v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
										v143 = int32(1)
										v146 = v16 + int32(16)
										v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
											return int32(0)
										} else {
											if v147 != 0 {
												v183 = v143
												m.G0 = v16 + int32(48)
												return v183
											} else {
												v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
												v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
												v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return int32(0)
												} else {
													if v156 != 0 {
														v183 = v143
														m.G0 = v16 + int32(48)
														return v183
													} else {
														v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
														v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
														v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return int32(0)
														} else {
															if v166 != 0 {
																v183 = v143
																m.G0 = v16 + int32(48)
																return v183
															} else {
																v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return int32(0)
																} else {
																	v183 = v175
																	m.G0 = v16 + int32(48)
																	return v183
																}
															}
														}
													}
												}
											}
										}
									} else {
										v125 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
										if base.F64_ge(v121, v125) == int32(0) {
											v137 = v121
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
											v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
											v143 = int32(1)
											v146 = v16 + int32(16)
											v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return int32(0)
											} else {
												if v147 != 0 {
													v183 = v143
													m.G0 = v16 + int32(48)
													return v183
												} else {
													v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
													v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
													v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														if v156 != 0 {
															v183 = v143
															m.G0 = v16 + int32(48)
															return v183
														} else {
															v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
															v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
															v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v167 = m.ExcPending
															if v167 != 0 {
																return int32(0)
															} else {
																if v166 != 0 {
																	v183 = v143
																	m.G0 = v16 + int32(48)
																	return v183
																} else {
																	v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																	v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																	v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return int32(0)
																	} else {
																		v183 = v175
																		m.G0 = v16 + int32(48)
																		return v183
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v129 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											if base.F64_le(v129, v125) == int32(0) {
												v137 = v121
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
												v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
												v143 = int32(1)
												v146 = v16 + int32(16)
												v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													if v147 != 0 {
														v183 = v143
														m.G0 = v16 + int32(48)
														return v183
													} else {
														v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
														v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
														v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															if v156 != 0 {
																v183 = v143
																m.G0 = v16 + int32(48)
																return v183
															} else {
																v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																mBase = m.M
																v167 = m.ExcPending
																if v167 != 0 {
																	return int32(0)
																} else {
																	if v166 != 0 {
																		v183 = v143
																		m.G0 = v16 + int32(48)
																		return v183
																	} else {
																		v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																		v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																		v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																		mBase = m.M
																		v176 = m.ExcPending
																		if v176 != 0 {
																			return int32(0)
																		} else {
																			v183 = v175
																			m.G0 = v16 + int32(48)
																			return v183
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v183 = int32(1)
												m.G0 = v16 + int32(48)
												return v183
											}
										}
									}
								}
							} else {
								v113 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
								if base.F64_le(v113, v108) != 0 {
									v183 = int32(1)
									m.G0 = v16 + int32(48)
									return v183
								} else {
									v116 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
									if base.F64_ge(v100, v116) == int32(0) {
										v120 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
										v137 = v120
										*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
										v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
										v143 = int32(1)
										v146 = v16 + int32(16)
										v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
											return int32(0)
										} else {
											if v147 != 0 {
												v183 = v143
												m.G0 = v16 + int32(48)
												return v183
											} else {
												v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
												v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
												v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
												mBase = m.M
												v157 = m.ExcPending
												if v157 != 0 {
													return int32(0)
												} else {
													if v156 != 0 {
														v183 = v143
														m.G0 = v16 + int32(48)
														return v183
													} else {
														v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
														v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
														v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return int32(0)
														} else {
															if v166 != 0 {
																v183 = v143
																m.G0 = v16 + int32(48)
																return v183
															} else {
																v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return int32(0)
																} else {
																	v183 = v175
																	m.G0 = v16 + int32(48)
																	return v183
																}
															}
														}
													}
												}
											}
										}
									} else {
										v121 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
										if base.F64_ge(v116, v99) == int32(0) {
											v137 = v121
											*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
											v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
											v143 = int32(1)
											v146 = v16 + int32(16)
											v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
											mBase = m.M
											v148 = m.ExcPending
											if v148 != 0 {
												return int32(0)
											} else {
												if v147 != 0 {
													v183 = v143
													m.G0 = v16 + int32(48)
													return v183
												} else {
													v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
													v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
													v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
													mBase = m.M
													v157 = m.ExcPending
													if v157 != 0 {
														return int32(0)
													} else {
														if v156 != 0 {
															v183 = v143
															m.G0 = v16 + int32(48)
															return v183
														} else {
															v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
															v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
															v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v167 = m.ExcPending
															if v167 != 0 {
																return int32(0)
															} else {
																if v166 != 0 {
																	v183 = v143
																	m.G0 = v16 + int32(48)
																	return v183
																} else {
																	v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																	v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																	v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return int32(0)
																	} else {
																		v183 = v175
																		m.G0 = v16 + int32(48)
																		return v183
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v125 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
											if base.F64_ge(v121, v125) == int32(0) {
												v137 = v121
												*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
												v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
												*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
												*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
												v143 = int32(1)
												v146 = v16 + int32(16)
												v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													if v147 != 0 {
														v183 = v143
														m.G0 = v16 + int32(48)
														return v183
													} else {
														v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
														v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
														*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
														*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
														*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
														v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int32(0)
														} else {
															if v156 != 0 {
																v183 = v143
																m.G0 = v16 + int32(48)
																return v183
															} else {
																v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																mBase = m.M
																v167 = m.ExcPending
																if v167 != 0 {
																	return int32(0)
																} else {
																	if v166 != 0 {
																		v183 = v143
																		m.G0 = v16 + int32(48)
																		return v183
																	} else {
																		v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																		v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																		*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																		v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																		mBase = m.M
																		v176 = m.ExcPending
																		if v176 != 0 {
																			return int32(0)
																		} else {
																			v183 = v175
																			m.G0 = v16 + int32(48)
																			return v183
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v129 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
												if base.F64_le(v129, v125) == int32(0) {
													v137 = v121
													*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
													v139 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
													*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
													*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
													*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v139
													v143 = int32(1)
													v146 = v16 + int32(16)
													v147 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														if v147 != 0 {
															v183 = v143
															m.G0 = v16 + int32(48)
															return v183
														} else {
															v149 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v149
															v151 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
															*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v137
															*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
															*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v151
															v156 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int32(0)
															} else {
																if v156 != 0 {
																	v183 = v143
																	m.G0 = v16 + int32(48)
																	return v183
																} else {
																	v158 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
																	v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																	v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v160
																	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v158
																	v166 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																	mBase = m.M
																	v167 = m.ExcPending
																	if v167 != 0 {
																		return int32(0)
																	} else {
																		if v166 != 0 {
																			v183 = v143
																			m.G0 = v16 + int32(48)
																			return v183
																		} else {
																			v168 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v168
																			v170 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v160
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v159
																			*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v170
																			v175 = F_lseg_interpt_lseg(m, int32(0), v146, l2)
																			mBase = m.M
																			v176 = m.ExcPending
																			if v176 != 0 {
																				return int32(0)
																			} else {
																				v183 = v175
																				m.G0 = v16 + int32(48)
																				return v183
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v183 = int32(1)
													m.G0 = v16 + int32(48)
													return v183
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
func F_box_lt(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_box_ar(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v11 = F_box_ar(m, v3)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return base.F64_lt(base.F64_add(v5, float64(1e-06)), v11)
		}
	}
}
func F_box_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_path_encode(m, int32(0), int32(2), v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_box_right(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.F64_gt(v3, base.F64_add(v5, float64(1e-06)))
}
func F_box_send(m *base.Module, l0 int32) int32 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		F_pq_sendfloat8(m, v5, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			F_pq_sendfloat8(m, v5, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
				F_pq_sendfloat8(m, v5, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
					F_pq_sendfloat8(m, v5, v21)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v25))) = v26 << (uint(int32(2)) % 32)
						m.G0 = v5 + int32(16)
						return v25
					}
				}
			}
		}
	}
}
func F_box_sub(m *base.Module, l0 int32) int32 {
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
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v53 float64
	_ = v53
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v89 int32
	_ = v89
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_palloc(m, int32(32))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v17 = base.F64_sub(v15, v16)
		v19 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.B2i32(base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19) == int32(0))&base.F64_ne(base.F64_abs(v16), v19) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v33 = base.F64_sub(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.B2i32(base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35) == int32(0))&base.F64_ne(base.F64_abs(v32), v35) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v33
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
				v49 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				v50 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
				v51 = base.F64_sub(v49, v50)
				v53 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.B2i32(base.F64_ne(base.F64_abs(v51), v53)|base.F64_eq(base.F64_abs(v49), v53) == int32(0))&base.F64_ne(base.F64_abs(v50), v53) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v65 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
					v66 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
					v67 = base.F64_sub(v65, v66)
					v69 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.B2i32(base.F64_ne(base.F64_abs(v67), v69)|base.F64_eq(base.F64_abs(v65), v69) == int32(0))&base.F64_ne(base.F64_abs(v66), v69) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v67
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v51
						return v11
					}
				}
			}
		}
	}
}
func F_box_width(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v5)+16))
	v8 = base.F64_sub(v6, v7)
	v10 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v8), v10)|base.F64_eq(base.F64_abs(v6), v10)|base.F64_eq(base.F64_abs(v7), v10) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v26 = F_Float8GetDatum(m, v8)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			return v26
		}
	}
}
