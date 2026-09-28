package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_bytea(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_uuid_send(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_uuid_decrement(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	v7 = F_palloc(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = base.I32_wrap_i64(l1)
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v12
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v14
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v16 != 0 {
			v100 = v7 + int32(15)
			v101 = v16
			v103 = v101 - int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
			v105 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
			return base.I64_extend_i32_u(v7)
		} else {
			v19 = int32(255)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v19)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
			if v21 != 0 {
				v100 = v7 + int32(14)
				v101 = v21
				v103 = v101 - int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
				v105 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
				return base.I64_extend_i32_u(v7)
			} else {
				v24 = int32(255)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v24)
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)))
				if v26 != 0 {
					v100 = v7 + int32(13)
					v101 = v26
					v103 = v101 - int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
					v105 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
					return base.I64_extend_i32_u(v7)
				} else {
					v29 = int32(255)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v29)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
					if v31 != 0 {
						v100 = v7 + int32(12)
						v101 = v31
						v103 = v101 - int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
						v105 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
						return base.I64_extend_i32_u(v7)
					} else {
						v34 = int32(255)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v34)
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)))
						if v36 != 0 {
							v100 = v7 + int32(11)
							v101 = v36
							v103 = v101 - int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
							v105 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
							return base.I64_extend_i32_u(v7)
						} else {
							v39 = int32(255)
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)) = uint8(v39)
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+10)))
							if v41 != 0 {
								v100 = v7 + int32(10)
								v101 = v41
								v103 = v101 - int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
								v105 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
								return base.I64_extend_i32_u(v7)
							} else {
								v44 = int32(255)
								*(*uint8)(unsafe.Add(mBase, uint32(v7)+10)) = uint8(v44)
								v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)))
								if v46 != 0 {
									v100 = v7 + int32(9)
									v101 = v46
									v103 = v101 - int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
									v105 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
									return base.I64_extend_i32_u(v7)
								} else {
									v49 = int32(255)
									*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)) = uint8(v49)
									v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)))
									if v51 != 0 {
										v100 = v7 + int32(8)
										v101 = v51
										v103 = v101 - int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
										v105 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
										return base.I64_extend_i32_u(v7)
									} else {
										v54 = int32(255)
										*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)) = uint8(v54)
										v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+7)))
										if v56 != 0 {
											v100 = v7 + int32(7)
											v101 = v56
											v103 = v101 - int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
											v105 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
											return base.I64_extend_i32_u(v7)
										} else {
											v59 = int32(255)
											*(*uint8)(unsafe.Add(mBase, uint32(v7)+7)) = uint8(v59)
											v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+6)))
											if v61 != 0 {
												v100 = v7 + int32(6)
												v101 = v61
												v103 = v101 - int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
												v105 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
												return base.I64_extend_i32_u(v7)
											} else {
												v64 = int32(255)
												*(*uint8)(unsafe.Add(mBase, uint32(v7)+6)) = uint8(v64)
												v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)))
												if v66 != 0 {
													v100 = v7 + int32(5)
													v101 = v66
													v103 = v101 - int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
													v105 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
													return base.I64_extend_i32_u(v7)
												} else {
													v69 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)) = uint8(v69)
													v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
													if v71 != 0 {
														v100 = v7 + int32(4)
														v101 = v71
														v103 = v101 - int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
														v105 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
														return base.I64_extend_i32_u(v7)
													} else {
														v74 = int32(255)
														*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v74)
														v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)))
														if v76 != 0 {
															v100 = v7 + int32(3)
															v101 = v76
															v103 = v101 - int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
															v105 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
															return base.I64_extend_i32_u(v7)
														} else {
															v79 = int32(255)
															*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)) = uint8(v79)
															v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)))
															if v81 != 0 {
																v100 = v7 + int32(2)
																v101 = v81
																v103 = v101 - int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
																v105 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
																return base.I64_extend_i32_u(v7)
															} else {
																v84 = int32(255)
																*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)) = uint8(v84)
																v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
																if v86 != 0 {
																	v100 = v7 + int32(1)
																	v101 = v86
																	v103 = v101 - int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
																	v105 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
																	return base.I64_extend_i32_u(v7)
																} else {
																	v89 = int32(255)
																	*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)) = uint8(v89)
																	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
																	if v91 != 0 {
																		v100 = v7
																		v101 = v91
																		v103 = v101 - int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
																		v105 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v105)
																		return base.I64_extend_i32_u(v7)
																	} else {
																		v92 = int32(255)
																		*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v92)
																		F_pfree(m, v7)
																		mBase = m.M
																		v95 = m.ExcPending
																		if v95 != 0 {
																			return int64(0)
																		} else {
																			v96 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v96)
																			return int64(0)
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
func F_uuid_fast_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v152 int64
	_ = v152
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v161 int32
	_ = v161
	v6 = base.I32_wrap_i64(l0)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v8 = int64(56)
	v10 = int64(65280)
	v12 = int64(40)
	v15 = int64(16711680)
	v17 = int64(24)
	v19 = int64(4278190080)
	v21 = int64(8)
	v42 = v7<<(uint(v8)%64) | v7&v10<<(uint(v12)%64) | (v7&v15<<(uint(v17)%64) | v7&v19<<(uint(v21)%64)) | (int64(base.Ui64(v7)>>(uint(v21)%64))&v19 | int64(base.Ui64(v7)>>(uint(v17)%64))&v15 | (int64(base.Ui64(v7)>>(uint(v12)%64))&v10 | int64(base.Ui64(v7)>>(uint(v8)%64))))
	v43 = base.I32_wrap_i64(l1)
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	v79 = v44<<(uint(v8)%64) | v44&v10<<(uint(v12)%64) | (v44&v15<<(uint(v17)%64) | v44&v19<<(uint(v21)%64)) | (int64(base.Ui64(v44)>>(uint(v21)%64))&v19 | int64(base.Ui64(v44)>>(uint(v17)%64))&v15 | (int64(base.Ui64(v44)>>(uint(v12)%64))&v10 | int64(base.Ui64(v44)>>(uint(v8)%64))))
	if v42 != v79 {
		v156 = v79
		v157 = v42
		if base.Ui64(v157) < base.Ui64(v156) {
			v161 = int32(-1)
		} else {
			v161 = int32(1)
		}
		return v161
	} else {
		v81 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
		v82 = int64(56)
		v84 = int64(65280)
		v86 = int64(40)
		v89 = int64(16711680)
		v91 = int64(24)
		v93 = int64(4278190080)
		v95 = int64(8)
		v116 = v81<<(uint(v82)%64) | v81&v84<<(uint(v86)%64) | (v81&v89<<(uint(v91)%64) | v81&v93<<(uint(v95)%64)) | (int64(base.Ui64(v81)>>(uint(v95)%64))&v93 | int64(base.Ui64(v81)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v81)>>(uint(v86)%64))&v84 | int64(base.Ui64(v81)>>(uint(v82)%64))))
		v117 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
		v152 = v117<<(uint(v82)%64) | v117&v84<<(uint(v86)%64) | (v117&v89<<(uint(v91)%64) | v117&v93<<(uint(v95)%64)) | (int64(base.Ui64(v117)>>(uint(v95)%64))&v93 | int64(base.Ui64(v117)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v117)>>(uint(v86)%64))&v84 | int64(base.Ui64(v117)>>(uint(v82)%64))))
		if v116 != v152 {
			v156 = v152
			v157 = v116
			if base.Ui64(v157) < base.Ui64(v156) {
				v161 = int32(-1)
			} else {
				v161 = int32(1)
			}
			return v161
		} else {
			return int32(0)
		}
	}
}
func F_uuid_generate_v1(m *base.Module, l0 int32) int64 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	F_uuid_generate_time(m, v6)
	v10 = v4 + int32(-48)
	F_uuid_unparse(m, v6, v10)
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v18 = F_DirectFunctionCall1Coll(m, int32(3591), int32(0), base.I64_extend_i32_u(v10))
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			m.G0 = v6 - int32(-64)
			return v18
		}
	}
}
func F_uuid_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = int64(56)
	v9 = int64(65280)
	v11 = int64(40)
	v14 = int64(16711680)
	v16 = int64(24)
	v18 = int64(4278190080)
	v20 = int64(8)
	v41 = v6<<(uint(v7)%64) | v6&v9<<(uint(v11)%64) | (v6&v14<<(uint(v16)%64) | v6&v18<<(uint(v20)%64)) | (int64(base.Ui64(v6)>>(uint(v20)%64))&v18 | int64(base.Ui64(v6)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v6)>>(uint(v11)%64))&v9 | int64(base.Ui64(v6)>>(uint(v7)%64))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	v78 = v43<<(uint(v7)%64) | v43&v9<<(uint(v11)%64) | (v43&v14<<(uint(v16)%64) | v43&v18<<(uint(v20)%64)) | (int64(base.Ui64(v43)>>(uint(v20)%64))&v18 | int64(base.Ui64(v43)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v43)>>(uint(v11)%64))&v9 | int64(base.Ui64(v43)>>(uint(v7)%64))))
	if v41 == v78 {
		v81 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v82 = int64(56)
		v84 = int64(65280)
		v86 = int64(40)
		v89 = int64(16711680)
		v91 = int64(24)
		v93 = int64(4278190080)
		v95 = int64(8)
		v116 = v81<<(uint(v82)%64) | v81&v84<<(uint(v86)%64) | (v81&v89<<(uint(v91)%64) | v81&v93<<(uint(v95)%64)) | (int64(base.Ui64(v81)>>(uint(v95)%64))&v93 | int64(base.Ui64(v81)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v81)>>(uint(v86)%64))&v84 | int64(base.Ui64(v81)>>(uint(v82)%64))))
		v117 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
		v152 = v117<<(uint(v82)%64) | v117&v84<<(uint(v86)%64) | (v117&v89<<(uint(v91)%64) | v117&v93<<(uint(v95)%64)) | (int64(base.Ui64(v117)>>(uint(v95)%64))&v93 | int64(base.Ui64(v117)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v117)>>(uint(v86)%64))&v84 | int64(base.Ui64(v117)>>(uint(v82)%64))))
		if v116 == v152 {
			v162 = int32(0)
		} else {
			v154 = v152
			v155 = v116
			if base.Ui64(v155) < base.Ui64(v154) {
				v159 = int32(-1)
			} else {
				v159 = int32(1)
			}
			v162 = v159
		}
	} else {
		v154 = v78
		v155 = v41
		if base.Ui64(v155) < base.Ui64(v154) {
			v159 = int32(-1)
		} else {
			v159 = int32(1)
		}
		v162 = v159
	}
	return base.I64_extend_i32_u(base.B2i32(v162 <= int32(0)))
}
func F_uuid_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v3)+8))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
	return base.I64_extend_i32_u(base.B2i32(v4^v6|(v8^v9) != int64(0)))
}
func F_uuid_nil(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14401(m, l0, int32(_a_F_uuid_nil_0), int32(_a_F_uuid_nil_1), int32(_a_F_uuid_nil_2), int32(_a_F_uuid_nil_3), int32(_a_F_uuid_nil_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
