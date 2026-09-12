package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_decrement(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v10
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v12
		v15 = v6 + int32(15)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
		if v16 != 0 {
			v98 = v16
			v99 = v15
			v101 = v98 - int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
			v103 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
			return v6
		} else {
			v17 = int32(255)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v17)
			v20 = v6 + int32(14)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			if v21 != 0 {
				v98 = v21
				v99 = v20
				v101 = v98 - int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
				v103 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
				return v6
			} else {
				v22 = int32(255)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v22)
				v25 = v6 + int32(13)
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
				if v26 != 0 {
					v98 = v26
					v99 = v25
					v101 = v98 - int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
					v103 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
					return v6
				} else {
					v27 = int32(255)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v27)
					v30 = v6 + int32(12)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
					if v31 != 0 {
						v98 = v31
						v99 = v30
						v101 = v98 - int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
						v103 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
						return v6
					} else {
						v32 = int32(255)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v32)
						v35 = v6 + int32(11)
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
						if v36 != 0 {
							v98 = v36
							v99 = v35
							v101 = v98 - int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
							v103 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
							return v6
						} else {
							v37 = int32(255)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)) = uint8(v37)
							v40 = v6 + int32(10)
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
							if v41 != 0 {
								v98 = v41
								v99 = v40
								v101 = v98 - int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
								v103 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
								return v6
							} else {
								v42 = int32(255)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)) = uint8(v42)
								v45 = v6 + int32(9)
								v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
								if v46 != 0 {
									v98 = v46
									v99 = v45
									v101 = v98 - int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
									v103 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
									return v6
								} else {
									v47 = int32(255)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)) = uint8(v47)
									v50 = v6 + int32(8)
									v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
									if v51 != 0 {
										v98 = v51
										v99 = v50
										v101 = v98 - int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
										v103 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
										return v6
									} else {
										v52 = int32(255)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)) = uint8(v52)
										v55 = v6 + int32(7)
										v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
										if v56 != 0 {
											v98 = v56
											v99 = v55
											v101 = v98 - int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
											v103 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
											return v6
										} else {
											v57 = int32(255)
											*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)) = uint8(v57)
											v60 = v6 + int32(6)
											v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
											if v61 != 0 {
												v98 = v61
												v99 = v60
												v101 = v98 - int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
												return v6
											} else {
												v62 = int32(255)
												*(*uint8)(unsafe.Add(mBase, uint32(v6)+6)) = uint8(v62)
												v65 = v6 + int32(5)
												v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
												if v66 != 0 {
													v98 = v66
													v99 = v65
													v101 = v98 - int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
													return v6
												} else {
													v67 = int32(255)
													*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)) = uint8(v67)
													v70 = v6 + int32(4)
													v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
													if v71 != 0 {
														v98 = v71
														v99 = v70
														v101 = v98 - int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
														v103 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
														return v6
													} else {
														v72 = int32(255)
														*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v72)
														v75 = v6 + int32(3)
														v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
														if v76 != 0 {
															v98 = v76
															v99 = v75
															v101 = v98 - int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
															v103 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
															return v6
														} else {
															v77 = int32(255)
															*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)) = uint8(v77)
															v80 = v6 + int32(2)
															v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
															if v81 != 0 {
																v98 = v81
																v99 = v80
																v101 = v98 - int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
																v103 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
																return v6
															} else {
																v82 = int32(255)
																*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)) = uint8(v82)
																v85 = v6 + int32(1)
																v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
																if v86 != 0 {
																	v98 = v86
																	v99 = v85
																	v101 = v98 - int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
																	v103 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
																	return v6
																} else {
																	v87 = int32(255)
																	*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v87)
																	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
																	if v89 != 0 {
																		v98 = v89
																		v99 = v6
																		v101 = v98 - int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v101)
																		v103 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v103)
																		return v6
																	} else {
																		v90 = int32(255)
																		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v90)
																		F_pfree(m, v6)
																		mBase = m.M
																		v93 = m.ExcPending
																		if v93 != 0 {
																			return int32(0)
																		} else {
																			v94 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v94)
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
func F_uuid_fast_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v4 = int32(16)
	goto L4
L1:
	;
	return v66
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	v40 = v35
	v41 = v36
	v42 = v37
	goto L13
L4:
	;
	if (l0|l1)&int32(3) != 0 {
		v35 = l0
		v36 = l1
		v37 = v4
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v12 = l0
	v13 = l1
	v14 = v4
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != v18 {
		v35 = v12
		v36 = v13
		v37 = v14
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v20 = int32(4)
	v21 = v13 + v20
	v23 = v12 + v20
	v25 = v14 - v20
	if base.Ui32(int32(3)) < base.Ui32(v25) {
		v12 = v23
		v13 = v21
		v14 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = v23
	v36 = v21
	v37 = v25
	goto L3
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v45 == v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = v45 - v46
	goto L1
L15:
	;
	v48 = int32(1)
	v53 = v42 - v48
	if v53 != 0 {
		v40 = v40 + v48
		v41 = v41 + v48
		v42 = v53
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L2
}
func F_uuid_generate_v1(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 + int32(-64)
	m.G0 = v5
	F_uuid_generate_time(m, v5)
	F_uuid_unparse(m, v5, v3+int32(-48))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v18 = F_DirectFunctionCall1Coll(m, int32(3395), int32(0), v3+int32(-48))
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 - int32(-64)
			return v18
		}
	}
}
func F_uuid_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(16)
	goto L4
L1:
	;
	return base.B2i32(v66 <= int32(0))
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	v40 = v35
	v41 = v36
	v42 = v37
	goto L13
L4:
	;
	if (v2|v3)&int32(3) != 0 {
		v35 = v2
		v36 = v3
		v37 = v4
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v12 = v2
	v13 = v3
	v14 = v4
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != v18 {
		v35 = v12
		v36 = v13
		v37 = v14
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v20 = int32(4)
	v21 = v13 + v20
	v23 = v12 + v20
	v25 = v14 - v20
	if base.Ui32(int32(3)) < base.Ui32(v25) {
		v12 = v23
		v13 = v21
		v14 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = v23
	v36 = v21
	v37 = v25
	goto L3
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v45 == v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = v45 - v46
	goto L1
L15:
	;
	v48 = int32(1)
	v53 = v42 - v48
	if v53 != 0 {
		v40 = v40 + v48
		v41 = v41 + v48
		v42 = v53
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L2
}
func F_uuid_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(16)
	goto L4
L1:
	;
	return base.B2i32(v66 != int32(0))
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	v40 = v35
	v41 = v36
	v42 = v37
	goto L13
L4:
	;
	if (v2|v3)&int32(3) != 0 {
		v35 = v2
		v36 = v3
		v37 = v4
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v12 = v2
	v13 = v3
	v14 = v4
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != v18 {
		v35 = v12
		v36 = v13
		v37 = v14
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v20 = int32(4)
	v21 = v13 + v20
	v23 = v12 + v20
	v25 = v14 - v20
	if base.Ui32(int32(3)) < base.Ui32(v25) {
		v12 = v23
		v13 = v21
		v14 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = v23
	v36 = v21
	v37 = v25
	goto L3
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v45 == v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = v45 - v46
	goto L1
L15:
	;
	v48 = int32(1)
	v53 = v42 - v48
	if v53 != 0 {
		v40 = v40 + v48
		v41 = v41 + v48
		v42 = v53
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L2
}
func F_uuid_nil(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v7 = int32(_a_F_uuid_nil_0)
	v8 = *(*int64)(unsafe.Add(mBase, _c_F_uuid_nil[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+29)) = v8
	v10 = *(*int64)(unsafe.Add(mBase, _c_F_uuid_nil[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = v10
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_uuid_nil[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+16)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, _c_F_uuid_nil[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v14
	v16 = *(*int64)(unsafe.Add(mBase, _c_F_uuid_nil[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v16
	v20 = F_DirectFunctionCall1Coll(m, int32(3395), int32(0), v5)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(48)
		return v20
	}
}
