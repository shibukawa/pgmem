package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_network_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v11 = Fn13942(m, l0, l1, int32(_a_F_network_abbrev_abort_0), int32(533), int32(_a_F_network_abbrev_abort_1), int32(_a_F_network_abbrev_abort_2), int32(526), int32(_a_F_network_abbrev_abort_3), int32(508), int32(_a_F_network_abbrev_abort_4))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_network_fast_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	v4 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_pg_detoast_datum_packed(m, l1)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v17 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
			if v19&v17 != 0 {
				v22 = v17
			} else {
				v22 = int32(4)
			}
			v23 = v4 + v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			v25 = int32(1)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v27&v25 != 0 {
				v30 = v25
			} else {
				v30 = int32(4)
			}
			v31 = v8 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v24 == v32 {
				v34 = int32(2)
				v35 = v23 + v34
				v37 = v31 + v34
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
				if base.Ui32(v38) < base.Ui32(v39) {
					v41 = v38
				} else {
					v41 = v39
				}
				v43 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
				v44 = F_memcmp(m, v35, v37, v43)
				mBase = m.M
				if v44 != 0 {
					v136 = v44
					v155 = v136
				} else {
					v46 = v41 & int32(7)
					if v46 == int32(0) {
						v127 = v38 - v39
						if v127 != 0 {
							v136 = v127
							v155 = v136
						} else {
							if v24 == int32(2) {
								v132 = int32(4)
							} else {
								v132 = int32(16)
							}
							v133 = F_memcmp(m, v35, v37, v132)
							mBase = m.M
							v155 = v133
						}
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v35))))
						v51 = int32(128)
						v52 = v50 & v51
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v37))))
						if v52 != v54&v51 {
							v143 = v52
							if v143 != 0 {
								v146 = int32(1)
							} else {
								v146 = int32(-1)
							}
							v155 = v146
						} else {
							if v46 == int32(1) {
								v127 = v38 - v39
								if v127 != 0 {
									v136 = v127
									v155 = v136
								} else {
									if v24 == int32(2) {
										v132 = int32(4)
									} else {
										v132 = int32(16)
									}
									v133 = F_memcmp(m, v35, v37, v132)
									mBase = m.M
									v155 = v133
								}
							} else {
								v60 = int32(1)
								v62 = int32(128)
								v63 = v50 << (uint(v60) % 32) & v62
								if v63 != v54<<(uint(v60)%32)&v62 {
									v143 = v63
									if v143 != 0 {
										v146 = int32(1)
									} else {
										v146 = int32(-1)
									}
									v155 = v146
								} else {
									if base.Ui32(v46) < base.Ui32(int32(3)) {
										v127 = v38 - v39
										if v127 != 0 {
											v136 = v127
											v155 = v136
										} else {
											if v24 == int32(2) {
												v132 = int32(4)
											} else {
												v132 = int32(16)
											}
											v133 = F_memcmp(m, v35, v37, v132)
											mBase = m.M
											v155 = v133
										}
									} else {
										v71 = int32(2)
										v73 = int32(128)
										v74 = v50 << (uint(v71) % 32) & v73
										if v74 != v54<<(uint(v71)%32)&v73 {
											v143 = v74
											if v143 != 0 {
												v146 = int32(1)
											} else {
												v146 = int32(-1)
											}
											v155 = v146
										} else {
											if v46 == int32(3) {
												v127 = v38 - v39
												if v127 != 0 {
													v136 = v127
													v155 = v136
												} else {
													if v24 == int32(2) {
														v132 = int32(4)
													} else {
														v132 = int32(16)
													}
													v133 = F_memcmp(m, v35, v37, v132)
													mBase = m.M
													v155 = v133
												}
											} else {
												v82 = int32(3)
												v84 = int32(128)
												v85 = v50 << (uint(v82) % 32) & v84
												if v85 != v54<<(uint(v82)%32)&v84 {
													v143 = v85
													if v143 != 0 {
														v146 = int32(1)
													} else {
														v146 = int32(-1)
													}
													v155 = v146
												} else {
													if base.Ui32(v46) < base.Ui32(int32(5)) {
														v127 = v38 - v39
														if v127 != 0 {
															v136 = v127
															v155 = v136
														} else {
															if v24 == int32(2) {
																v132 = int32(4)
															} else {
																v132 = int32(16)
															}
															v133 = F_memcmp(m, v35, v37, v132)
															mBase = m.M
															v155 = v133
														}
													} else {
														v93 = int32(4)
														v95 = int32(128)
														v96 = v50 << (uint(v93) % 32) & v95
														if v96 != v54<<(uint(v93)%32)&v95 {
															v143 = v96
															if v143 != 0 {
																v146 = int32(1)
															} else {
																v146 = int32(-1)
															}
															v155 = v146
														} else {
															if v46 == int32(5) {
																v127 = v38 - v39
																if v127 != 0 {
																	v136 = v127
																	v155 = v136
																} else {
																	if v24 == int32(2) {
																		v132 = int32(4)
																	} else {
																		v132 = int32(16)
																	}
																	v133 = F_memcmp(m, v35, v37, v132)
																	mBase = m.M
																	v155 = v133
																}
															} else {
																v104 = int32(5)
																v106 = int32(128)
																v107 = v50 << (uint(v104) % 32) & v106
																if v107 != v54<<(uint(v104)%32)&v106 {
																	v143 = v107
																	if v143 != 0 {
																		v146 = int32(1)
																	} else {
																		v146 = int32(-1)
																	}
																	v155 = v146
																} else {
																	if v46 != int32(7) {
																		v127 = v38 - v39
																		if v127 != 0 {
																			v136 = v127
																			v155 = v136
																		} else {
																			if v24 == int32(2) {
																				v132 = int32(4)
																			} else {
																				v132 = int32(16)
																			}
																			v133 = F_memcmp(m, v35, v37, v132)
																			mBase = m.M
																			v155 = v133
																		}
																	} else {
																		v115 = int32(6)
																		v117 = int32(128)
																		v118 = v50 << (uint(v115) % 32) & v117
																		if v118 != v54<<(uint(v115)%32)&v117 {
																			v143 = v118
																			if v143 != 0 {
																				v146 = int32(1)
																			} else {
																				v146 = int32(-1)
																			}
																			v155 = v146
																		} else {
																			v127 = v38 - v39
																			if v127 != 0 {
																				v136 = v127
																				v155 = v136
																			} else {
																				if v24 == int32(2) {
																					v132 = int32(4)
																				} else {
																					v132 = int32(16)
																				}
																				v133 = F_memcmp(m, v35, v37, v132)
																				mBase = m.M
																				v155 = v133
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
			} else {
				v136 = v24 - v32
				v155 = v136
			}
			return v155
		}
	}
}
func F_network_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v17 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
			if v19&v17 != 0 {
				v22 = v17
			} else {
				v22 = int32(4)
			}
			v23 = v3 + v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			v25 = int32(1)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v27&v25 != 0 {
				v30 = v25
			} else {
				v30 = int32(4)
			}
			v31 = v8 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v24 == v32 {
				v34 = int32(2)
				v35 = v23 + v34
				v37 = v31 + v34
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
				if base.Ui32(v38) < base.Ui32(v39) {
					v41 = v38
				} else {
					v41 = v39
				}
				v43 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
				v44 = F_memcmp(m, v35, v37, v43)
				mBase = m.M
				if v44 != 0 {
					v136 = v44
					v155 = v136
				} else {
					v46 = v41 & int32(7)
					if v46 == int32(0) {
						v127 = v38 - v39
						if v127 != 0 {
							v136 = v127
							v155 = v136
						} else {
							if v24 == int32(2) {
								v132 = int32(4)
							} else {
								v132 = int32(16)
							}
							v133 = F_memcmp(m, v35, v37, v132)
							mBase = m.M
							v155 = v133
						}
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v35))))
						v51 = int32(128)
						v52 = v50 & v51
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v37))))
						if v52 != v54&v51 {
							v143 = v52
							if v143 != 0 {
								v146 = int32(1)
							} else {
								v146 = int32(-1)
							}
							v155 = v146
						} else {
							if v46 == int32(1) {
								v127 = v38 - v39
								if v127 != 0 {
									v136 = v127
									v155 = v136
								} else {
									if v24 == int32(2) {
										v132 = int32(4)
									} else {
										v132 = int32(16)
									}
									v133 = F_memcmp(m, v35, v37, v132)
									mBase = m.M
									v155 = v133
								}
							} else {
								v60 = int32(1)
								v62 = int32(128)
								v63 = v50 << (uint(v60) % 32) & v62
								if v63 != v54<<(uint(v60)%32)&v62 {
									v143 = v63
									if v143 != 0 {
										v146 = int32(1)
									} else {
										v146 = int32(-1)
									}
									v155 = v146
								} else {
									if base.Ui32(v46) < base.Ui32(int32(3)) {
										v127 = v38 - v39
										if v127 != 0 {
											v136 = v127
											v155 = v136
										} else {
											if v24 == int32(2) {
												v132 = int32(4)
											} else {
												v132 = int32(16)
											}
											v133 = F_memcmp(m, v35, v37, v132)
											mBase = m.M
											v155 = v133
										}
									} else {
										v71 = int32(2)
										v73 = int32(128)
										v74 = v50 << (uint(v71) % 32) & v73
										if v74 != v54<<(uint(v71)%32)&v73 {
											v143 = v74
											if v143 != 0 {
												v146 = int32(1)
											} else {
												v146 = int32(-1)
											}
											v155 = v146
										} else {
											if v46 == int32(3) {
												v127 = v38 - v39
												if v127 != 0 {
													v136 = v127
													v155 = v136
												} else {
													if v24 == int32(2) {
														v132 = int32(4)
													} else {
														v132 = int32(16)
													}
													v133 = F_memcmp(m, v35, v37, v132)
													mBase = m.M
													v155 = v133
												}
											} else {
												v82 = int32(3)
												v84 = int32(128)
												v85 = v50 << (uint(v82) % 32) & v84
												if v85 != v54<<(uint(v82)%32)&v84 {
													v143 = v85
													if v143 != 0 {
														v146 = int32(1)
													} else {
														v146 = int32(-1)
													}
													v155 = v146
												} else {
													if base.Ui32(v46) < base.Ui32(int32(5)) {
														v127 = v38 - v39
														if v127 != 0 {
															v136 = v127
															v155 = v136
														} else {
															if v24 == int32(2) {
																v132 = int32(4)
															} else {
																v132 = int32(16)
															}
															v133 = F_memcmp(m, v35, v37, v132)
															mBase = m.M
															v155 = v133
														}
													} else {
														v93 = int32(4)
														v95 = int32(128)
														v96 = v50 << (uint(v93) % 32) & v95
														if v96 != v54<<(uint(v93)%32)&v95 {
															v143 = v96
															if v143 != 0 {
																v146 = int32(1)
															} else {
																v146 = int32(-1)
															}
															v155 = v146
														} else {
															if v46 == int32(5) {
																v127 = v38 - v39
																if v127 != 0 {
																	v136 = v127
																	v155 = v136
																} else {
																	if v24 == int32(2) {
																		v132 = int32(4)
																	} else {
																		v132 = int32(16)
																	}
																	v133 = F_memcmp(m, v35, v37, v132)
																	mBase = m.M
																	v155 = v133
																}
															} else {
																v104 = int32(5)
																v106 = int32(128)
																v107 = v50 << (uint(v104) % 32) & v106
																if v107 != v54<<(uint(v104)%32)&v106 {
																	v143 = v107
																	if v143 != 0 {
																		v146 = int32(1)
																	} else {
																		v146 = int32(-1)
																	}
																	v155 = v146
																} else {
																	if v46 != int32(7) {
																		v127 = v38 - v39
																		if v127 != 0 {
																			v136 = v127
																			v155 = v136
																		} else {
																			if v24 == int32(2) {
																				v132 = int32(4)
																			} else {
																				v132 = int32(16)
																			}
																			v133 = F_memcmp(m, v35, v37, v132)
																			mBase = m.M
																			v155 = v133
																		}
																	} else {
																		v115 = int32(6)
																		v117 = int32(128)
																		v118 = v50 << (uint(v115) % 32) & v117
																		if v118 != v54<<(uint(v115)%32)&v117 {
																			v143 = v118
																			if v143 != 0 {
																				v146 = int32(1)
																			} else {
																				v146 = int32(-1)
																			}
																			v155 = v146
																		} else {
																			v127 = v38 - v39
																			if v127 != 0 {
																				v136 = v127
																				v155 = v136
																			} else {
																				if v24 == int32(2) {
																					v132 = int32(4)
																				} else {
																					v132 = int32(16)
																				}
																				v133 = F_memcmp(m, v35, v37, v132)
																				mBase = m.M
																				v155 = v133
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
			} else {
				v136 = v24 - v32
				v155 = v136
			}
			return base.B2i32(int32(0) < v155)
		}
	}
}
func F_network_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v18 = int32(1)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
			if v20&v18 != 0 {
				v23 = v18
			} else {
				v23 = int32(4)
			}
			v24 = v4 + v23
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
			v26 = int32(1)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v28&v26 != 0 {
				v31 = v26
			} else {
				v31 = int32(4)
			}
			v32 = v9 + v31
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
			if v25 == v33 {
				v35 = int32(2)
				v36 = v24 + v35
				v38 = v32 + v35
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
				if base.Ui32(v39) < base.Ui32(v40) {
					v42 = v39
				} else {
					v42 = v40
				}
				v44 = int32(base.Ui32(v42) >> (uint(int32(3)) % 32))
				v45 = F_memcmp(m, v36, v38, v44)
				mBase = m.M
				if v45 != 0 {
					v137 = v45
					v156 = v137
				} else {
					v47 = v42 & int32(7)
					if v47 == int32(0) {
						v128 = v39 - v40
						if v128 != 0 {
							v137 = v128
							v156 = v137
						} else {
							if v25 == int32(2) {
								v133 = int32(4)
							} else {
								v133 = int32(16)
							}
							v134 = F_memcmp(m, v36, v38, v133)
							mBase = m.M
							v156 = v134
						}
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v36))))
						v52 = int32(128)
						v53 = v51 & v52
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v38))))
						if v53 != v55&v52 {
							v144 = v53
							if v144 != 0 {
								v147 = int32(1)
							} else {
								v147 = int32(-1)
							}
							v156 = v147
						} else {
							if v47 == int32(1) {
								v128 = v39 - v40
								if v128 != 0 {
									v137 = v128
									v156 = v137
								} else {
									if v25 == int32(2) {
										v133 = int32(4)
									} else {
										v133 = int32(16)
									}
									v134 = F_memcmp(m, v36, v38, v133)
									mBase = m.M
									v156 = v134
								}
							} else {
								v61 = int32(1)
								v63 = int32(128)
								v64 = v51 << (uint(v61) % 32) & v63
								if v64 != v55<<(uint(v61)%32)&v63 {
									v144 = v64
									if v144 != 0 {
										v147 = int32(1)
									} else {
										v147 = int32(-1)
									}
									v156 = v147
								} else {
									if base.Ui32(v47) < base.Ui32(int32(3)) {
										v128 = v39 - v40
										if v128 != 0 {
											v137 = v128
											v156 = v137
										} else {
											if v25 == int32(2) {
												v133 = int32(4)
											} else {
												v133 = int32(16)
											}
											v134 = F_memcmp(m, v36, v38, v133)
											mBase = m.M
											v156 = v134
										}
									} else {
										v72 = int32(2)
										v74 = int32(128)
										v75 = v51 << (uint(v72) % 32) & v74
										if v75 != v55<<(uint(v72)%32)&v74 {
											v144 = v75
											if v144 != 0 {
												v147 = int32(1)
											} else {
												v147 = int32(-1)
											}
											v156 = v147
										} else {
											if v47 == int32(3) {
												v128 = v39 - v40
												if v128 != 0 {
													v137 = v128
													v156 = v137
												} else {
													if v25 == int32(2) {
														v133 = int32(4)
													} else {
														v133 = int32(16)
													}
													v134 = F_memcmp(m, v36, v38, v133)
													mBase = m.M
													v156 = v134
												}
											} else {
												v83 = int32(3)
												v85 = int32(128)
												v86 = v51 << (uint(v83) % 32) & v85
												if v86 != v55<<(uint(v83)%32)&v85 {
													v144 = v86
													if v144 != 0 {
														v147 = int32(1)
													} else {
														v147 = int32(-1)
													}
													v156 = v147
												} else {
													if base.Ui32(v47) < base.Ui32(int32(5)) {
														v128 = v39 - v40
														if v128 != 0 {
															v137 = v128
															v156 = v137
														} else {
															if v25 == int32(2) {
																v133 = int32(4)
															} else {
																v133 = int32(16)
															}
															v134 = F_memcmp(m, v36, v38, v133)
															mBase = m.M
															v156 = v134
														}
													} else {
														v94 = int32(4)
														v96 = int32(128)
														v97 = v51 << (uint(v94) % 32) & v96
														if v97 != v55<<(uint(v94)%32)&v96 {
															v144 = v97
															if v144 != 0 {
																v147 = int32(1)
															} else {
																v147 = int32(-1)
															}
															v156 = v147
														} else {
															if v47 == int32(5) {
																v128 = v39 - v40
																if v128 != 0 {
																	v137 = v128
																	v156 = v137
																} else {
																	if v25 == int32(2) {
																		v133 = int32(4)
																	} else {
																		v133 = int32(16)
																	}
																	v134 = F_memcmp(m, v36, v38, v133)
																	mBase = m.M
																	v156 = v134
																}
															} else {
																v105 = int32(5)
																v107 = int32(128)
																v108 = v51 << (uint(v105) % 32) & v107
																if v108 != v55<<(uint(v105)%32)&v107 {
																	v144 = v108
																	if v144 != 0 {
																		v147 = int32(1)
																	} else {
																		v147 = int32(-1)
																	}
																	v156 = v147
																} else {
																	if v47 != int32(7) {
																		v128 = v39 - v40
																		if v128 != 0 {
																			v137 = v128
																			v156 = v137
																		} else {
																			if v25 == int32(2) {
																				v133 = int32(4)
																			} else {
																				v133 = int32(16)
																			}
																			v134 = F_memcmp(m, v36, v38, v133)
																			mBase = m.M
																			v156 = v134
																		}
																	} else {
																		v116 = int32(6)
																		v118 = int32(128)
																		v119 = v51 << (uint(v116) % 32) & v118
																		if v119 != v55<<(uint(v116)%32)&v118 {
																			v144 = v119
																			if v144 != 0 {
																				v147 = int32(1)
																			} else {
																				v147 = int32(-1)
																			}
																			v156 = v147
																		} else {
																			v128 = v39 - v40
																			if v128 != 0 {
																				v137 = v128
																				v156 = v137
																			} else {
																				if v25 == int32(2) {
																					v133 = int32(4)
																				} else {
																					v133 = int32(16)
																				}
																				v134 = F_memcmp(m, v36, v38, v133)
																				mBase = m.M
																				v156 = v134
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
			} else {
				v137 = v25 - v33
				v156 = v137
			}
			if int32(0) < v156 {
				v159 = v4
			} else {
				v159 = v9
			}
			return v159
		}
	}
}
func F_network_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v17 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
			if v19&v17 != 0 {
				v22 = v17
			} else {
				v22 = int32(4)
			}
			v23 = v3 + v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			v25 = int32(1)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v27&v25 != 0 {
				v30 = v25
			} else {
				v30 = int32(4)
			}
			v31 = v8 + v30
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v24 == v32 {
				v34 = int32(2)
				v35 = v23 + v34
				v37 = v31 + v34
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
				if base.Ui32(v38) < base.Ui32(v39) {
					v41 = v38
				} else {
					v41 = v39
				}
				v43 = int32(base.Ui32(v41) >> (uint(int32(3)) % 32))
				v44 = F_memcmp(m, v35, v37, v43)
				mBase = m.M
				if v44 != 0 {
					v136 = v44
					v155 = v136
				} else {
					v46 = v41 & int32(7)
					if v46 == int32(0) {
						v127 = v38 - v39
						if v127 != 0 {
							v136 = v127
							v155 = v136
						} else {
							if v24 == int32(2) {
								v132 = int32(4)
							} else {
								v132 = int32(16)
							}
							v133 = F_memcmp(m, v35, v37, v132)
							mBase = m.M
							v155 = v133
						}
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v35))))
						v51 = int32(128)
						v52 = v50 & v51
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v37))))
						if v52 != v54&v51 {
							v143 = v52
							if v143 != 0 {
								v146 = int32(1)
							} else {
								v146 = int32(-1)
							}
							v155 = v146
						} else {
							if v46 == int32(1) {
								v127 = v38 - v39
								if v127 != 0 {
									v136 = v127
									v155 = v136
								} else {
									if v24 == int32(2) {
										v132 = int32(4)
									} else {
										v132 = int32(16)
									}
									v133 = F_memcmp(m, v35, v37, v132)
									mBase = m.M
									v155 = v133
								}
							} else {
								v60 = int32(1)
								v62 = int32(128)
								v63 = v50 << (uint(v60) % 32) & v62
								if v63 != v54<<(uint(v60)%32)&v62 {
									v143 = v63
									if v143 != 0 {
										v146 = int32(1)
									} else {
										v146 = int32(-1)
									}
									v155 = v146
								} else {
									if base.Ui32(v46) < base.Ui32(int32(3)) {
										v127 = v38 - v39
										if v127 != 0 {
											v136 = v127
											v155 = v136
										} else {
											if v24 == int32(2) {
												v132 = int32(4)
											} else {
												v132 = int32(16)
											}
											v133 = F_memcmp(m, v35, v37, v132)
											mBase = m.M
											v155 = v133
										}
									} else {
										v71 = int32(2)
										v73 = int32(128)
										v74 = v50 << (uint(v71) % 32) & v73
										if v74 != v54<<(uint(v71)%32)&v73 {
											v143 = v74
											if v143 != 0 {
												v146 = int32(1)
											} else {
												v146 = int32(-1)
											}
											v155 = v146
										} else {
											if v46 == int32(3) {
												v127 = v38 - v39
												if v127 != 0 {
													v136 = v127
													v155 = v136
												} else {
													if v24 == int32(2) {
														v132 = int32(4)
													} else {
														v132 = int32(16)
													}
													v133 = F_memcmp(m, v35, v37, v132)
													mBase = m.M
													v155 = v133
												}
											} else {
												v82 = int32(3)
												v84 = int32(128)
												v85 = v50 << (uint(v82) % 32) & v84
												if v85 != v54<<(uint(v82)%32)&v84 {
													v143 = v85
													if v143 != 0 {
														v146 = int32(1)
													} else {
														v146 = int32(-1)
													}
													v155 = v146
												} else {
													if base.Ui32(v46) < base.Ui32(int32(5)) {
														v127 = v38 - v39
														if v127 != 0 {
															v136 = v127
															v155 = v136
														} else {
															if v24 == int32(2) {
																v132 = int32(4)
															} else {
																v132 = int32(16)
															}
															v133 = F_memcmp(m, v35, v37, v132)
															mBase = m.M
															v155 = v133
														}
													} else {
														v93 = int32(4)
														v95 = int32(128)
														v96 = v50 << (uint(v93) % 32) & v95
														if v96 != v54<<(uint(v93)%32)&v95 {
															v143 = v96
															if v143 != 0 {
																v146 = int32(1)
															} else {
																v146 = int32(-1)
															}
															v155 = v146
														} else {
															if v46 == int32(5) {
																v127 = v38 - v39
																if v127 != 0 {
																	v136 = v127
																	v155 = v136
																} else {
																	if v24 == int32(2) {
																		v132 = int32(4)
																	} else {
																		v132 = int32(16)
																	}
																	v133 = F_memcmp(m, v35, v37, v132)
																	mBase = m.M
																	v155 = v133
																}
															} else {
																v104 = int32(5)
																v106 = int32(128)
																v107 = v50 << (uint(v104) % 32) & v106
																if v107 != v54<<(uint(v104)%32)&v106 {
																	v143 = v107
																	if v143 != 0 {
																		v146 = int32(1)
																	} else {
																		v146 = int32(-1)
																	}
																	v155 = v146
																} else {
																	if v46 != int32(7) {
																		v127 = v38 - v39
																		if v127 != 0 {
																			v136 = v127
																			v155 = v136
																		} else {
																			if v24 == int32(2) {
																				v132 = int32(4)
																			} else {
																				v132 = int32(16)
																			}
																			v133 = F_memcmp(m, v35, v37, v132)
																			mBase = m.M
																			v155 = v133
																		}
																	} else {
																		v115 = int32(6)
																		v117 = int32(128)
																		v118 = v50 << (uint(v115) % 32) & v117
																		if v118 != v54<<(uint(v115)%32)&v117 {
																			v143 = v118
																			if v143 != 0 {
																				v146 = int32(1)
																			} else {
																				v146 = int32(-1)
																			}
																			v155 = v146
																		} else {
																			v127 = v38 - v39
																			if v127 != 0 {
																				v136 = v127
																				v155 = v136
																			} else {
																				if v24 == int32(2) {
																					v132 = int32(4)
																				} else {
																					v132 = int32(16)
																				}
																				v133 = F_memcmp(m, v35, v37, v132)
																				mBase = m.M
																				v155 = v133
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
			} else {
				v136 = v24 - v32
				v155 = v136
			}
			return base.B2i32(v155 <= int32(0))
		}
	}
}
